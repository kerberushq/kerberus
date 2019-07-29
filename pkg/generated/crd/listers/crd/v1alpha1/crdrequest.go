/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kerberushq/kerberus/pkg/apis/crd/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CRDRequestLister helps list CRDRequests.
type CRDRequestLister interface {
	// List lists all CRDRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CRDRequest, err error)
	// CRDRequests returns an object that can list and get CRDRequests.
	CRDRequests(namespace string) CRDRequestNamespaceLister
	CRDRequestListerExpansion
}

// cRDRequestLister implements the CRDRequestLister interface.
type cRDRequestLister struct {
	indexer cache.Indexer
}

// NewCRDRequestLister returns a new CRDRequestLister.
func NewCRDRequestLister(indexer cache.Indexer) CRDRequestLister {
	return &cRDRequestLister{indexer: indexer}
}

// List lists all CRDRequests in the indexer.
func (s *cRDRequestLister) List(selector labels.Selector) (ret []*v1alpha1.CRDRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CRDRequest))
	})
	return ret, err
}

// CRDRequests returns an object that can list and get CRDRequests.
func (s *cRDRequestLister) CRDRequests(namespace string) CRDRequestNamespaceLister {
	return cRDRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CRDRequestNamespaceLister helps list and get CRDRequests.
type CRDRequestNamespaceLister interface {
	// List lists all CRDRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CRDRequest, err error)
	// Get retrieves the CRDRequest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CRDRequest, error)
	CRDRequestNamespaceListerExpansion
}

// cRDRequestNamespaceLister implements the CRDRequestNamespaceLister
// interface.
type cRDRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CRDRequests in the indexer for a given namespace.
func (s cRDRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CRDRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CRDRequest))
	})
	return ret, err
}

// Get retrieves the CRDRequest from the indexer for a given namespace and name.
func (s cRDRequestNamespaceLister) Get(name string) (*v1alpha1.CRDRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("crdrequest"), name)
	}
	return obj.(*v1alpha1.CRDRequest), nil
}