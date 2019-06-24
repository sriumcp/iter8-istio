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

package experiment

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	servingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.ibm.com/istio-research/iter8-controller/pkg/analytics/checkandincrement"
	iter8v1alpha1 "github.ibm.com/istio-research/iter8-controller/pkg/apis/iter8/v1alpha1"
)

func (r *ReconcileExperiment) syncKnative(context context.Context, instance *iter8v1alpha1.Experiment) (reconcile.Result, error) {
	log := Logger(context)

	// Get Knative service
	serviceName := instance.Spec.TargetService.Name
	serviceNamespace := instance.Spec.TargetService.Namespace
	if serviceNamespace == "" {
		serviceNamespace = instance.Namespace
	}

	kservice := &servingv1alpha1.Service{}
	err := r.Get(context, types.NamespacedName{Name: serviceName, Namespace: serviceNamespace}, kservice)
	if err != nil {
		instance.Status.MarkHasNotService("NotFound", "")
		err = r.Status().Update(context, instance)
		if err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	if kservice.Spec.Template == nil {
		instance.Status.MarkHasNotService("MissingTemplate", "")
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	// link service to this experiment. Only one experiment can control a service
	labels := kservice.GetLabels()
	if experiment, found := labels[experimentLabel]; found && experiment != instance.GetName() {
		instance.Status.MarkHasNotService("ExistingExperiment", "service is already controlled by %v", experiment)
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	if labels == nil {
		labels = make(map[string]string)
	}

	if _, ok := labels[experimentLabel]; !ok {
		labels[experimentLabel] = instance.GetName()
		kservice.SetLabels(labels)
		if err = r.Update(context, kservice); err != nil {
			return reconcile.Result{}, err
		}
	}

	// Check the experiment targets existing traffic targets
	ksvctraffic := kservice.Spec.Traffic
	if ksvctraffic == nil {
		instance.Status.MarkHasNotService("MissingTraffic", "")
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	baseline := instance.Spec.TargetService.Baseline
	baselineTraffic := getTrafficByName(kservice, baseline)
	if baselineTraffic == nil {
		instance.Status.MarkHasNotService("MissingBaselineRevision", "%s", baseline)
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	candidate := instance.Spec.TargetService.Candidate
	candidateTraffic := getTrafficByName(kservice, candidate)
	if candidateTraffic == nil {
		instance.Status.MarkHasNotService("MissingCandidateRevision", "%s", candidate)
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	instance.Status.MarkHasService()

	traffic := instance.Spec.TrafficControl
	now := time.Now()
	interval, _ := traffic.GetIntervalDuration() // TODO: admissioncontrollervalidation

	if instance.Status.StartTimestamp == "" {
		instance.Status.StartTimestamp = strconv.FormatInt(metav1.NewTime(now).UTC().Unix(), 10)
		updateGrafanaURL(instance, serviceNamespace)
	}

	// check experiment is finished
	if traffic.GetMaxIterations() <= instance.Status.CurrentIteration {

		update := false
		if instance.Status.AssessmentSummary.AllSuccessCriteriaMet {
			log.Info("Experiment completed with success")
			// experiment is successful
			switch traffic.GetOnSuccess() {
			case "baseline":
				// Rollback
				if candidateTraffic.Percent != 0 {
					candidateTraffic.Percent = 0
					update = true
				}
				if baselineTraffic.Percent != 100 {
					baselineTraffic.Percent = 100
					update = true
				}
				instance.Status.MarkNotRollForward("OnSuccessBaseline", "")
				instance.Status.TrafficSplit.Baseline = 100
				instance.Status.TrafficSplit.Candidate = 0
			case "candidate":
				// Rollforward
				if candidateTraffic.Percent != 100 {
					candidateTraffic.Percent = 100
					update = true
				}
				if baselineTraffic.Percent != 0 {
					baselineTraffic.Percent = 0
					update = true
				}
				instance.Status.MarkRollForward()
				instance.Status.TrafficSplit.Baseline = 100
				instance.Status.TrafficSplit.Candidate = 0
			case "both":
				instance.Status.MarkNotRollForward("OnSuccessBoth", "")
			}
		} else {
			log.Info("Experiment completed with failure")

			// Switch traffic back to baseline
			if candidateTraffic.Percent != 0 {
				candidateTraffic.Percent = 0
				update = true
			}
			if baselineTraffic.Percent != 100 {
				baselineTraffic.Percent = 100
				update = true
			}
		}

		labels := kservice.GetLabels()
		_, has := labels[experimentLabel]
		if has {
			delete(labels, experimentLabel)
		}

		if has || update {
			err := r.Update(context, kservice)
			if err != nil {
				return reconcile.Result{}, err // retry
			}
		}

		// Clear analysis state
		instance.Status.AnalysisState.Raw = []byte("{}")

		// End experiment
		instance.Status.MarkExperimentCompleted()
		err = r.Status().Update(context, instance)
		return reconcile.Result{}, err
	}

	// Check if traffic should be updated.

	if now.After(instance.Status.LastIncrementTime.Add(interval)) {
		log.Info("process iteration.")

		newRolloutPercent := float64(candidateTraffic.Percent)
		switch instance.Spec.TrafficControl.GetStrategy() {
		case "increment_without_check":
			newRolloutPercent += traffic.GetStepSize()
		case "check_and_increment":
			// Get underlying k8s services
			// TODO: should just get the service name. See issue #83
			baselineService, err := r.getServiceForRevision(context, kservice, baselineTraffic.RevisionName)
			if err != nil {
				// TODO: maybe we want another condition
				instance.Status.MarkHasNotService("MissingCoreService", "%v", err)
				err = r.Status().Update(context, instance)
				return reconcile.Result{}, err
			}
			candidateService, err := r.getServiceForRevision(context, kservice, candidateTraffic.RevisionName)
			if err != nil {
				// TODO: maybe we want another condition
				instance.Status.MarkHasNotService("MissingCoreService", "%v", err)
				err = r.Status().Update(context, instance)
				return reconcile.Result{}, err
			}

			// Get latest analysis
			payload := MakeRequest(instance, baselineService, candidateService)
			response, err := checkandincrement.Invoke(log, instance.Spec.Analysis.GetServiceEndpoint(), payload)
			if err != nil {
				instance.Status.MarkExperimentNotCompleted("ErrorAnalytics", "%v", err)
				err = r.Status().Update(context, instance)
				return reconcile.Result{RequeueAfter: 5 * time.Second}, err
			}

			// Abort?
			if response.Assessment.Summary.AbortExperiment {
				log.Info("abort experiment.")
				if candidateTraffic.Percent != 0 || baselineTraffic.Percent != 100 {
					baselineTraffic.Percent = 100
					candidateTraffic.Percent = 0
					err := r.Update(context, kservice)
					if err != nil {
						return reconcile.Result{}, err // retry
					}
				}
			}

			baselineTraffic := response.Baseline.TrafficPercentage
			candidateTraffic := response.Canary.TrafficPercentage
			log.Info("NewTraffic", "baseline", baselineTraffic, "candidate", candidateTraffic)
			newRolloutPercent = candidateTraffic

			if response.LastState == nil {
				instance.Status.AnalysisState.Raw = []byte("{}")
			} else {
				lastState, err := json.Marshal(response.LastState)
				if err != nil {
					instance.Status.MarkExperimentNotCompleted("ErrorAnalyticsResponse", "%v", err)
					err = r.Status().Update(context, instance)
					return reconcile.Result{RequeueAfter: 5 * time.Second}, err
				}
				instance.Status.AnalysisState = runtime.RawExtension{Raw: lastState}
			}
			instance.Status.AssessmentSummary = response.Assessment.Summary
			instance.Status.CurrentIteration++
		}

		// Set traffic percentable on all routes
		needUpdate := false
		for _, target := range ksvctraffic {
			if target.RevisionName == baseline {
				if target.Percent != 100-int(newRolloutPercent) {
					target.Percent = 100 - int(newRolloutPercent)
					needUpdate = true
				}
			} else if target.RevisionName == candidate {
				if target.Percent != int(newRolloutPercent) {
					target.Percent = int(newRolloutPercent)
					needUpdate = true
				}
			} else {
				if target.Percent != 0 {
					target.Percent = 0
					needUpdate = true
				}
			}
		}
		if needUpdate {
			log.Info("update traffic", "rolloutPercent", newRolloutPercent)

			err = r.Update(context, kservice) // TODO: patch?
			if err != nil {
				// TODO: the analysis service will be called again upon retry. Maybe we do want that.
				return reconcile.Result{}, err
			}
		}

		instance.Status.LastIncrementTime = metav1.NewTime(now)
	}

	instance.Status.MarkExperimentNotCompleted("Progressing", "")
	err = r.Status().Update(context, instance)
	return reconcile.Result{RequeueAfter: interval}, err
}

func getTrafficByName(service *servingv1alpha1.Service, name string) *servingv1alpha1.TrafficTarget {
	for _, traffic := range service.Spec.Traffic {
		if traffic.RevisionName == name {
			return &traffic
		}
	}
	return nil
}

func (r *ReconcileExperiment) getServiceForRevision(context context.Context, ksvc *servingv1alpha1.Service, revisionName string) (*corev1.Service, error) {
	revision := &servingv1alpha1.Revision{}
	err := r.Get(context, types.NamespacedName{Name: revisionName, Namespace: ksvc.GetNamespace()}, revision)
	if err != nil {
		return nil, err
	}
	service := &corev1.Service{}
	err = r.Get(context, types.NamespacedName{Name: revision.Status.ServiceName, Namespace: ksvc.GetNamespace()}, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}
