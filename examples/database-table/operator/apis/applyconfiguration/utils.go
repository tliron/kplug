// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	myorgorgv1alpha1 "github.com/tliron/kplug/examples/database-table/operator/apis/applyconfiguration/myorg.org/v1alpha1"
	v1alpha1 "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=myorg.org, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("DatabaseTable"):
		return &myorgorgv1alpha1.DatabaseTableApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DatabaseTableColumn"):
		return &myorgorgv1alpha1.DatabaseTableColumnApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DatabaseTableSpec"):
		return &myorgorgv1alpha1.DatabaseTableSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DatabaseTableStatus"):
		return &myorgorgv1alpha1.DatabaseTableStatusApplyConfiguration{}

	}
	return nil
}
