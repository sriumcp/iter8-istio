// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RuleLister helps list Rules.
type RuleLister interface {
	// List lists all Rules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.Rule, err error)
	// Rules returns an object that can list and get Rules.
	Rules(namespace string) RuleNamespaceLister
	RuleListerExpansion
}

// ruleLister implements the RuleLister interface.
type ruleLister struct {
	indexer cache.Indexer
}

// NewRuleLister returns a new RuleLister.
func NewRuleLister(indexer cache.Indexer) RuleLister {
	return &ruleLister{indexer: indexer}
}

// List lists all Rules in the indexer.
func (s *ruleLister) List(selector labels.Selector) (ret []*v1alpha2.Rule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Rule))
	})
	return ret, err
}

// Rules returns an object that can list and get Rules.
func (s *ruleLister) Rules(namespace string) RuleNamespaceLister {
	return ruleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuleNamespaceLister helps list and get Rules.
type RuleNamespaceLister interface {
	// List lists all Rules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.Rule, err error)
	// Get retrieves the Rule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.Rule, error)
	RuleNamespaceListerExpansion
}

// ruleNamespaceLister implements the RuleNamespaceLister
// interface.
type ruleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Rules in the indexer for a given namespace.
func (s ruleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.Rule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Rule))
	})
	return ret, err
}

// Get retrieves the Rule from the indexer for a given namespace and name.
func (s ruleNamespaceLister) Get(name string) (*v1alpha2.Rule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("rule"), name)
	}
	return obj.(*v1alpha2.Rule), nil
}
