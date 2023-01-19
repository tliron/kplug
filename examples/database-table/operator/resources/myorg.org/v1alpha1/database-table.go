package v1alpha1

import (
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatabaseTableGVK = SchemeGroupVersion.WithKind(DatabaseTableKind)

const (
	DatabaseTableKind     = "DatabaseTable"
	DatabaseTableListKind = "DatabaseTableList"

	DatabaseTableSingular  = "databasetable"
	DatabaseTablePlural    = "databasetables"
	DatabaseTableShortName = "dt" // = DatabaseTable
)

//
// DatabaseTable
//

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DatabaseTable struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseTableSpec   `json:"spec"`
	Status DatabaseTableStatus `json:"status"`
}

type DatabaseTableSpec struct {
	ServerName              string                 `json:"serverName"`
	PreferredImplementation string                 `json:"preferredImplementation"`
	Columns                 []DatabaseTableColumn  `json:"columns,omitempty"`
	References              []core.ObjectReference `json:"references,omitempty"`
}

type DatabaseTableColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DatabaseTableStatus struct {
	Implementation string `json:"implementation"`
}

//
// DatabaseTableList
//

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DatabaseTableList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata"`

	Items []DatabaseTable `json:"items"`
}
