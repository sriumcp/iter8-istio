/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package canary

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.ibm.com/istio-research/iter8-controller/pkg/analytics/checkandincrement"
	iter8v1alpha1 "github.ibm.com/istio-research/iter8-controller/pkg/apis/iter8/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1alpha3 "github.com/knative/pkg/apis/istio/v1alpha3"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const (
	Baseline  = "baseline"
	Candidate = "candidate"
	Stable    = "stable"
)

var ServiceSelector map[string]string

func (r *ReconcileCanary) syncIstio(context context.Context, canary *iter8v1alpha1.Canary) (reconcile.Result, error) {
	serviceName := canary.Spec.TargetService.Name
	serviceNamespace := canary.Spec.TargetService.Namespace
	if serviceNamespace == "" {
		serviceNamespace = canary.Namespace
	}

	// Get k8s service
	service := &corev1.Service{}
	err := r.Get(context, types.NamespacedName{Name: serviceName, Namespace: serviceNamespace}, service)
	if err != nil {
		canary.Status.MarkHasNotService("NotFound", "")
		err = r.Status().Update(context, canary)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	ServiceSelector = make(map[string]string)
	log.Info("istio sync", "service name", service.GetName())
	for key, val := range service.Spec.Selector {
		ServiceSelector[key] = val
	}
	log.Info("istio sync", "Initilize selector", ServiceSelector)
	canary.Status.MarkHasService()

	// Get deployment list
	deployments := &appsv1.DeploymentList{}
	if err = r.List(context, client.MatchingLabels(service.Spec.Selector), deployments); err != nil || len(deployments.Items) == 0 {
		// TODO: add new type of status to set unavailable deployments
		canary.Status.MarkHasNotService("NotFound", "")
		err = r.Status().Update(context, canary)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	// Get current deployment and candidate deployment
	var baseline, candidate *appsv1.Deployment
	for _, d := range deployments.Items {
		if val, ok := d.ObjectMeta.Labels[canaryLabel]; ok {
			if val == Candidate {
				candidate = d.DeepCopy()
			} else if val == Baseline {
				baseline = d.DeepCopy()
			}
		}
	}

	//	log.Info("istio sync", "Checking baseline and canary...")
	// Case 1: Baseline is not existed, mark the latest deployment as baseline
	// Waits for canary
	if baseline == nil {
		d, err := getLatestDeployment(deployments)
		if err != nil {
			return reconcile.Result{}, err
		}
		d.ObjectMeta.SetLabels(map[string]string{canaryLabel: Baseline})
		err = r.Update(context, d)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Info("istio sync", "Label baseline", d.GetName())
		return reconcile.Result{}, nil
	} else if candidate == nil {
		// Promote the latest deployment as candidate
		d, err := getCanaryDeployment(deployments, baseline)
		if err != nil {
			return reconcile.Result{}, err
		}
		d.ObjectMeta.SetLabels(map[string]string{canaryLabel: Candidate})
		err = r.Update(context, d)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Info("istio sync", "Label candidate", d.GetName())
		return reconcile.Result{Requeue: true}, err
	}

	log.Info("istio-sync", "baseline", baseline.GetName(), Candidate, candidate.GetName())

	// Remove stable rules if there is any
	stableName := getStableName(canary)
	dr := &v1alpha3.DestinationRule{}
	if err = r.Get(context, types.NamespacedName{Name: stableName, Namespace: canary.GetNamespace()}, dr); err == nil {
		r.Delete(context, dr)
		log.Info("istio sync", "delete stable dr", stableName)
	}
	vs := &v1alpha3.VirtualService{}
	if err = r.Get(context, types.NamespacedName{Name: stableName, Namespace: canary.GetNamespace()}, vs); err == nil {
		r.Delete(context, vs)
		log.Info("istio sync", "delete stable vs", stableName)
	}

	// Get info on Canary
	traffic := canary.Spec.TrafficControl
	now := time.Now()
	// TODO: check err in getting the time value
	interval, _ := traffic.GetIntervalDuration()

	// Start Canary Process
	// Setup Istio Routing Rules
	// TODO: should include deployment info here
	drName := getDestinationRuleName(canary)
	dr = &v1alpha3.DestinationRule{}
	if err = r.Get(context, types.NamespacedName{Name: drName, Namespace: canary.Namespace}, dr); err != nil {
		dr = newDestinationRule(canary, baseline, candidate)
		err := r.Create(context, dr)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Info("istio-sync", "create destinationRule", drName)
	}

	vsName := getVirtualServiceName(canary)
	vs = &v1alpha3.VirtualService{}
	if err = r.Get(context, types.NamespacedName{Name: vsName, Namespace: canary.Namespace}, vs); err != nil {
		vs = makeVirtualService(0, canary)
		err := r.Create(context, vs)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Info("istio-sync", "create virtualservice", vsName)
	}

	// Check canary rollout status
	rolloutPercent := float64(getWeight(Candidate, vs))
	log.Info("istio-sync", "prev rollout percent", rolloutPercent, "max traffic percent", traffic.GetMaxTrafficPercent())
	if rolloutPercent < traffic.GetMaxTrafficPercent() &&
		now.After(canary.Status.LastIncrementTime.Add(interval)) {

		switch canary.Spec.TrafficControl.Strategy {
		case "manual":
			rolloutPercent += traffic.GetStepSize()
		case "check_and_increment":
			// Get latest analysis
			payload := MakeRequest(canary, baseline, candidate)
			response, err := checkandincrement.Invoke(log, canary.Spec.Analysis.AnalyticsService, payload)
			if err != nil {
				// TODO: Need new condition
				canary.Status.MarkHasNotService("ErrorAnalytics", "%v", err)
				err = r.Status().Update(context, canary)
				deleteRules(context, r, canary)
				return reconcile.Result{}, err
			}

			baselineTraffic := response.Baseline.TrafficPercentage
			canaryTraffic := response.Canary.TrafficPercentage
			log.Info("NewTraffic", "baseline", baselineTraffic, "canary", canaryTraffic)
			rolloutPercent = canaryTraffic

			lastState, err := json.Marshal(response.LastState)
			if err != nil {
				// TODO: Need new condition
				canary.Status.MarkHasNotService("ErrorAnalyticsResponse", "%v", err)
				err = r.Status().Update(context, canary)
				return reconcile.Result{}, err
			}
			canary.Status.AnalysisState = runtime.RawExtension{Raw: lastState}
		}

		log.Info("istio-sync", "new rollout perccent", rolloutPercent)
		rv := vs.ObjectMeta.ResourceVersion
		vs = makeVirtualService(int(rolloutPercent), canary)
		setResourceVersion(rv, vs)
		log.Info("istio-sync", "updated vs", *vs)
		err := r.Update(context, vs)
		if err != nil {
			return reconcile.Result{}, err
		}

		canary.Status.LastIncrementTime = metav1.NewTime(now)
	}

	result := reconcile.Result{}
	if getWeight(Candidate, vs) == int(traffic.GetMaxTrafficPercent()) {
		// Rollout done.
		canary.Status.MarkRolloutCompleted()
		canary.Status.Progressing = false
		// remove labels
		removeCanaryLabel(context, r, baseline)
		removeCanaryLabel(context, r, candidate)
		// delete rules
		deleteRules(context, r, canary)
		// generate new rules to shift all traffic to candidate
		stableDr, stableVs := newStableRules(candidate, canary)
		r.Create(context, stableDr)
		r.Create(context, stableVs)
	} else {
		canary.Status.MarkRolloutNotCompleted("Progressing", "")
		canary.Status.Progressing = true
		result.RequeueAfter = interval
	}

	err = r.Status().Update(context, canary)

	return result, err
}

func removeCanaryLabel(context context.Context, r *ReconcileCanary, d *appsv1.Deployment) {
	labels := d.GetLabels()
	delete(labels, canaryLabel)
	d.SetLabels(labels)
	r.Update(context, d)
	log.Info("istio sync", "remove labels", d.GetName())
}

func deleteRules(context context.Context, r *ReconcileCanary, canary *iter8v1alpha1.Canary) (err error) {
	drName := getDestinationRuleName(canary)
	vsName := getVirtualServiceName(canary)

	dr := &v1alpha3.DestinationRule{}
	if err = r.Get(context, types.NamespacedName{Name: drName, Namespace: canary.Namespace}, dr); err == nil {
		r.Delete(context, dr)
		log.Info("istio sync", "delete dr", drName)
	}

	vs := &v1alpha3.VirtualService{}
	if err = r.Get(context, types.NamespacedName{Name: vsName, Namespace: canary.Namespace}, vs); err == nil {
		r.Delete(context, vs)
		log.Info("istio sync", "delete vs", vsName)
	}

	return err
}

// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: reviews-canary
// spec:
//   host: reviews
//   subsets:
//   - name: base
// 	   labels:
// 	     iter8.ibm.com/canary: base
//   - name: candidate
// 	   labels:
// 	     iter8.ibm.com/canary: candidate

func newDestinationRule(canary *iter8v1alpha1.Canary, baseline, candidate *appsv1.Deployment) *v1alpha3.DestinationRule {
	bLabels := baseline.GetLabels()
	cLabels := candidate.GetLabels()
	delete(bLabels, canaryLabel)
	delete(cLabels, canaryLabel)
	dr := &v1alpha3.DestinationRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getDestinationRuleName(canary),
			Namespace: canary.Namespace,
			// TODO: add owner references
		},
		Spec: v1alpha3.DestinationRuleSpec{
			Host: canary.Spec.TargetService.Name,
			Subsets: []v1alpha3.Subset{
				v1alpha3.Subset{
					Name:   Baseline,
					Labels: bLabels,
				},
				v1alpha3.Subset{
					Name:   Candidate,
					Labels: cLabels,
				},
			},
		},
	}

	return dr
}

func getDestinationRuleName(canary *iter8v1alpha1.Canary) string {
	return canary.Spec.TargetService.Name + "-iter8.canary"
}

// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: reviews-canary
// spec:
//   hosts:
// 	- reviews
//   http:
//   - route:
// 		- destination:
// 			host: reviews
// 			subset: base
// 	  	  weight: 50
// 		- destination:
// 			host: reviews
// 			subset: candidate
// 	  	  weight: 50

func makeVirtualService(rolloutPercent int, canary *iter8v1alpha1.Canary) *v1alpha3.VirtualService {
	vs := &v1alpha3.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getVirtualServiceName(canary),
			Namespace: canary.Namespace,
			// TODO: add owner references
		},
		Spec: v1alpha3.VirtualServiceSpec{
			Hosts: []string{canary.Spec.TargetService.Name},
			HTTP: []v1alpha3.HTTPRoute{
				{
					Route: []v1alpha3.HTTPRouteDestination{
						{
							Destination: v1alpha3.Destination{
								Host:   canary.Spec.TargetService.Name,
								Subset: Baseline,
							},
							Weight: 100 - rolloutPercent,
						},
						{
							Destination: v1alpha3.Destination{
								Host:   canary.Spec.TargetService.Name,
								Subset: Candidate,
							},
							Weight: rolloutPercent,
						},
					},
				},
			},
		},
	}

	return vs
}

// Should add deployment names
func getVirtualServiceName(canary *iter8v1alpha1.Canary) string {
	return canary.Spec.TargetService.Name + "-iter8.canary"
}

func getWeight(subset string, vs *v1alpha3.VirtualService) int {
	for _, route := range vs.Spec.HTTP[0].Route {
		if route.Destination.Subset == subset {
			return route.Weight
		}
	}
	return 0
}

func setResourceVersion(rv string, vs *v1alpha3.VirtualService) {
	vs.ObjectMeta.ResourceVersion = rv
}

<<<<<<< HEAD
func getLatestDeployment(ds *appsv1.DeploymentList) (*appsv1.Deployment, error) {
	latestTs := *new(time.Time)
	index := -1
	for i, d := range ds.Items {
		if val, ok := d.ObjectMeta.Labels[canaryLabel]; !ok || val != Baseline {
			ct := d.ObjectMeta.CreationTimestamp
			if ct.After(latestTs) {
				latestTs = ct.Time
				index = i
			}
		}
	}

	if index == -1 {
		return nil, errors.New("Latest deployment not found")
	}
	return &ds.Items[index], nil
}

func getCanaryDeployment(ds *appsv1.DeploymentList, baseline *appsv1.Deployment) (*appsv1.Deployment, error) {
	baselineTs := baseline.ObjectMeta.CreationTimestamp.Time
	index := -1
	for i, d := range ds.Items {
		if val, ok := d.ObjectMeta.Labels[canaryLabel]; !ok || val != Baseline {
			ct := d.ObjectMeta.CreationTimestamp.Time
			if ct.After(baselineTs) {
				index = i
			}
		}
	}

	if index == -1 {
		return nil, errors.New("Latest deployment not found")
	}
	return &ds.Items[index], nil
=======
func newStableRules(d *appsv1.Deployment, canary *iter8v1alpha1.Canary) (*v1alpha3.DestinationRule, *v1alpha3.VirtualService) {
	dr := &v1alpha3.DestinationRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getStableName(canary),
			Namespace: canary.Namespace,
			// TODO: add owner references
		},
		Spec: v1alpha3.DestinationRuleSpec{
			Host: canary.Spec.TargetService.Name,
			Subsets: []v1alpha3.Subset{
				v1alpha3.Subset{
					Name:   Stable,
					Labels: d.GetLabels(),
				},
			},
		},
	}

	vs := &v1alpha3.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      getStableName(canary),
			Namespace: canary.Namespace,
			// TODO: add owner references
		},
		Spec: v1alpha3.VirtualServiceSpec{
			Hosts: []string{canary.Spec.TargetService.Name},
			HTTP: []v1alpha3.HTTPRoute{
				{
					Route: []v1alpha3.HTTPRouteDestination{
						{
							Destination: v1alpha3.Destination{
								Host:   canary.Spec.TargetService.Name,
								Subset: Stable,
							},
							Weight: 100,
						},
					},
				},
			},
		},
	}

	return dr, vs
}

func getStableName(canary *iter8v1alpha1.Canary) string {
	return canary.GetName() + "iter-stable"
>>>>>>> update istio logic
}
