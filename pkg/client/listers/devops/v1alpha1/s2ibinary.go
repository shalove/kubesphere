/*
Copyright 2019 The KubeSphere authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha1"
)

// S2iBinaryLister helps list S2iBinaries.
type S2iBinaryLister interface {
	// List lists all S2iBinaries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S2iBinary, err error)
	// S2iBinaries returns an object that can list and get S2iBinaries.
	S2iBinaries(namespace string) S2iBinaryNamespaceLister
	S2iBinaryListerExpansion
}

// s2iBinaryLister implements the S2iBinaryLister interface.
type s2iBinaryLister struct {
	indexer cache.Indexer
}

// NewS2iBinaryLister returns a new S2iBinaryLister.
func NewS2iBinaryLister(indexer cache.Indexer) S2iBinaryLister {
	return &s2iBinaryLister{indexer: indexer}
}

// List lists all S2iBinaries in the indexer.
func (s *s2iBinaryLister) List(selector labels.Selector) (ret []*v1alpha1.S2iBinary, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S2iBinary))
	})
	return ret, err
}

// S2iBinaries returns an object that can list and get S2iBinaries.
func (s *s2iBinaryLister) S2iBinaries(namespace string) S2iBinaryNamespaceLister {
	return s2iBinaryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// S2iBinaryNamespaceLister helps list and get S2iBinaries.
type S2iBinaryNamespaceLister interface {
	// List lists all S2iBinaries in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.S2iBinary, err error)
	// Get retrieves the S2iBinary from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.S2iBinary, error)
	S2iBinaryNamespaceListerExpansion
}

// s2iBinaryNamespaceLister implements the S2iBinaryNamespaceLister
// interface.
type s2iBinaryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all S2iBinaries in the indexer for a given namespace.
func (s s2iBinaryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.S2iBinary, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S2iBinary))
	})
	return ret, err
}

// Get retrieves the S2iBinary from the indexer for a given namespace and name.
func (s s2iBinaryNamespaceLister) Get(name string) (*v1alpha1.S2iBinary, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s2ibinary"), name)
	}
	return obj.(*v1alpha1.S2iBinary), nil
}
