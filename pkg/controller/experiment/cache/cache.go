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

package cache

import (
	"context"
	"sync"

	"github.com/go-logr/logr"

	iter8v1alpha1 "github.com/iter8-tools/iter8-controller/pkg/apis/iter8/v1alpha1"
	"github.com/iter8-tools/iter8-controller/pkg/controller/experiment/cache/abstract"
	"github.com/iter8-tools/iter8-controller/pkg/controller/experiment/util"
)

// Interface defines the interface for iter8cache
type Interface interface {
	// Given name and namespace of the target deployment, return the experiment key
	DeploymentToExperiment(name, namespace string) (experiment, experimentNamespace string, exist bool)
	// Given name and namespace of the target service, return the experiment key
	ServiceToExperiment(name, namespace string) (experiment, experimentNamespace string, exist bool)
	RegisterExperiment(context context.Context, instance *iter8v1alpha1.Experiment) context.Context
	RemoveExperiment(instance *iter8v1alpha1.Experiment)

	MarkTargetDeploymentFound(name, namespace string) bool
	MarkTargetServiceFound(name, namespace string) bool

	MarkTargetDeploymentMissing(name, namespace string) bool
	MarkTargetServiceMissing(name, namespace string) bool
}

var _ Interface = &Impl{}

// Impl is the implementation of Iter8Cache
type Impl struct {
	logger logr.Logger

	// the mutext to protect the maps
	m sync.RWMutex
	// an ExperimentAbstract store with experimentName.experimentNamespace as key for access
	experimentAbstractStore map[string]*abstract.Experiment

	// a lookup map from target to experiment
	// targetName.targetNamespace -> experimentName.experimentNamespace
	deployment2Experiment map[string]string

	// a lookup map from target service to experiment
	service2Experiment map[string]string
}

// New returns a new iter8cache implementation
func New(logger logr.Logger) Interface {
	return &Impl{
		experimentAbstractStore: make(map[string]*abstract.Experiment),
		deployment2Experiment:   make(map[string]string),
		service2Experiment:      make(map[string]string),
		logger:                  logger,
	}
}

// RegisterExperiment creates new abstracts into the cache and snapshot the abstract into context
func (c *Impl) RegisterExperiment(ctx context.Context, instance *iter8v1alpha1.Experiment) context.Context {
	c.m.Lock()
	defer c.m.Unlock()

	eakey := experimentKey(instance)
	if _, ok := c.experimentAbstractStore[eakey]; !ok {
		targetNamespace := util.GetServiceNamespace(instance)
		ea := abstract.NewExperiment(instance, targetNamespace)
		c.experimentAbstractStore[eakey] = ea
		service := instance.Spec.TargetService.Name
		baseline := instance.Spec.TargetService.Baseline
		candidate := instance.Spec.TargetService.Candidate

		c.service2Experiment[targetKey(service, targetNamespace)] = eakey
		c.deployment2Experiment[targetKey(baseline, targetNamespace)] = eakey
		c.deployment2Experiment[targetKey(candidate, targetNamespace)] = eakey
	}

	ctx = context.WithValue(ctx, util.AbstractKey, c.experimentAbstractStore[eakey])

	return ctx
}

// DeploymentToExperiment returns the experiment key given name and namespace of target deployment
func (c *Impl) DeploymentToExperiment(targetName, targetNamespace string) (string, string, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	if _, ok := c.deployment2Experiment[tKey]; !ok {
		return "", "", false
	}
	name, namespace := resolveExperimentKey(c.deployment2Experiment[tKey])

	return name, namespace, true
}

func (c *Impl) MarkTargetDeploymentFound(targetName, targetNamespace string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	eaKey, ok := c.deployment2Experiment[tKey]
	if !ok {
		return false
	}

	c.experimentAbstractStore[eaKey].TargetsAbstract.MarkTargetFound(targetName, true)

	return true
}

func (c *Impl) MarkTargetDeploymentMissing(targetName, targetNamespace string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	eaKey, ok := c.deployment2Experiment[tKey]
	if !ok {
		return false
	}

	c.experimentAbstractStore[eaKey].TargetsAbstract.MarkTargetFound(targetName, false)

	return true
}

// ServiceToExperiment returns the experiment key given name and namespace of target service
func (c *Impl) ServiceToExperiment(targetName, targetNamespace string) (string, string, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	if _, ok := c.service2Experiment[tKey]; !ok {
		return "", "", false
	}

	name, namespace := resolveExperimentKey(c.service2Experiment[tKey])

	return name, namespace, true
}

func (c *Impl) MarkTargetServiceFound(targetName, targetNamespace string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	eaKey, ok := c.service2Experiment[tKey]
	if !ok {
		return false
	}

	c.experimentAbstractStore[eaKey].TargetsAbstract.MarkServiceFound(true)

	return true
}

func (c *Impl) MarkTargetServiceMissing(targetName, targetNamespace string) bool {
	c.m.Lock()
	defer c.m.Unlock()

	tKey := targetKey(targetName, targetNamespace)
	eaKey, ok := c.service2Experiment[tKey]
	if !ok {
		return false
	}

	c.experimentAbstractStore[eaKey].TargetsAbstract.MarkServiceFound(false)

	return true
}

// RemoveExperiment removes the experiment abstract from the cache
func (c *Impl) RemoveExperiment(instance *iter8v1alpha1.Experiment) {
	c.m.Lock()
	defer c.m.Unlock()

	eakey := experimentKey(instance)
	ea, ok := c.experimentAbstractStore[eakey]
	if !ok {
		return
	}

	ta := ea.TargetsAbstract
	targetNamespace := ta.Namespace
	delete(c.service2Experiment, targetKey(ta.ServiceName, targetNamespace))
	for name := range ta.Status {
		delete(c.deployment2Experiment, targetKey(name, targetNamespace))
	}
	delete(c.experimentAbstractStore, eakey)
}