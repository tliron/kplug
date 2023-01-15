// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/tliron/kplug/examples/database-table/operator/apis/clientset/versioned"
	internalinterfaces "github.com/tliron/kplug/examples/database-table/operator/apis/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tliron/kplug/examples/database-table/operator/apis/listers/myorg.org/v1alpha1"
	myorgorgv1alpha1 "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatabaseTableInformer provides access to a shared informer and lister for
// DatabaseTables.
type DatabaseTableInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatabaseTableLister
}

type databaseTableInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatabaseTableInformer constructs a new informer for DatabaseTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatabaseTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatabaseTableInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatabaseTableInformer constructs a new informer for DatabaseTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatabaseTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MyorgV1alpha1().DatabaseTables(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MyorgV1alpha1().DatabaseTables(namespace).Watch(context.TODO(), options)
			},
		},
		&myorgorgv1alpha1.DatabaseTable{},
		resyncPeriod,
		indexers,
	)
}

func (f *databaseTableInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatabaseTableInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *databaseTableInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&myorgorgv1alpha1.DatabaseTable{}, f.defaultInformer)
}

func (f *databaseTableInformer) Lister() v1alpha1.DatabaseTableLister {
	return v1alpha1.NewDatabaseTableLister(f.Informer().GetIndexer())
}