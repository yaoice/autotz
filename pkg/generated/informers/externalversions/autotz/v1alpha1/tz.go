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

	autotzv1alpha1 "github.com/yaoice/autotz/pkg/apis/autotz/v1alpha1"
	versioned "github.com/yaoice/autotz/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/yaoice/autotz/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/yaoice/autotz/pkg/generated/listers/autotz/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TZInformer provides access to a shared informer and lister for
// TZs.
type TZInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TZLister
}

type tZInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTZInformer constructs a new informer for TZ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTZInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTZInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTZInformer constructs a new informer for TZ type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTZInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutotzV1alpha1().TZs().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutotzV1alpha1().TZs().Watch(options)
			},
		},
		&autotzv1alpha1.TZ{},
		resyncPeriod,
		indexers,
	)
}

func (f *tZInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTZInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tZInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autotzv1alpha1.TZ{}, f.defaultInformer)
}

func (f *tZInformer) Lister() v1alpha1.TZLister {
	return v1alpha1.NewTZLister(f.Informer().GetIndexer())
}
