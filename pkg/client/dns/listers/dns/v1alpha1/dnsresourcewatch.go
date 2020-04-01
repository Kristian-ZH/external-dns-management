/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DNSResourceWatchLister helps list DNSResourceWatches.
type DNSResourceWatchLister interface {
	// List lists all DNSResourceWatches in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DNSAnnotation, err error)
	// DNSResourceWatches returns an object that can list and get DNSResourceWatches.
	DNSResourceWatches(namespace string) DNSResourceWatchNamespaceLister
	DNSResourceWatchListerExpansion
}

// dNSResourceWatchLister implements the DNSResourceWatchLister interface.
type dNSResourceWatchLister struct {
	indexer cache.Indexer
}

// NewDNSResourceWatchLister returns a new DNSResourceWatchLister.
func NewDNSResourceWatchLister(indexer cache.Indexer) DNSResourceWatchLister {
	return &dNSResourceWatchLister{indexer: indexer}
}

// List lists all DNSResourceWatches in the indexer.
func (s *dNSResourceWatchLister) List(selector labels.Selector) (ret []*v1alpha1.DNSAnnotation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSAnnotation))
	})
	return ret, err
}

// DNSResourceWatches returns an object that can list and get DNSResourceWatches.
func (s *dNSResourceWatchLister) DNSResourceWatches(namespace string) DNSResourceWatchNamespaceLister {
	return dNSResourceWatchNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DNSResourceWatchNamespaceLister helps list and get DNSResourceWatches.
type DNSResourceWatchNamespaceLister interface {
	// List lists all DNSResourceWatches in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DNSAnnotation, err error)
	// Get retrieves the DNSAnnotation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DNSAnnotation, error)
	DNSResourceWatchNamespaceListerExpansion
}

// dNSResourceWatchNamespaceLister implements the DNSResourceWatchNamespaceLister
// interface.
type dNSResourceWatchNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DNSResourceWatches in the indexer for a given namespace.
func (s dNSResourceWatchNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DNSAnnotation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DNSAnnotation))
	})
	return ret, err
}

// Get retrieves the DNSAnnotation from the indexer for a given namespace and name.
func (s dNSResourceWatchNamespaceLister) Get(name string) (*v1alpha1.DNSAnnotation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsresourcewatch"), name)
	}
	return obj.(*v1alpha1.DNSAnnotation), nil
}
