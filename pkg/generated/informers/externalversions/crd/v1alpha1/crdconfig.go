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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	crdv1alpha1 "github.com/kerberushq/kerberus/pkg/apis/crd/v1alpha1"
	versioned "github.com/kerberushq/kerberus/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/kerberushq/kerberus/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kerberushq/kerberus/pkg/generated/listers/crd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CRDConfigInformer provides access to a shared informer and lister for
// CRDConfigs.
type CRDConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CRDConfigLister
}

type cRDConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCRDConfigInformer constructs a new informer for CRDConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCRDConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCRDConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCRDConfigInformer constructs a new informer for CRDConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCRDConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().CRDConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().CRDConfigs(namespace).Watch(options)
			},
		},
		&crdv1alpha1.CRDConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *cRDConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCRDConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cRDConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crdv1alpha1.CRDConfig{}, f.defaultInformer)
}

func (f *cRDConfigInformer) Lister() v1alpha1.CRDConfigLister {
	return v1alpha1.NewCRDConfigLister(f.Informer().GetIndexer())
}
