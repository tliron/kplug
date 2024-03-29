// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// DatabaseTableColumnApplyConfiguration represents an declarative configuration of the DatabaseTableColumn type for use
// with apply.
type DatabaseTableColumnApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// DatabaseTableColumnApplyConfiguration constructs an declarative configuration of the DatabaseTableColumn type for use with
// apply.
func DatabaseTableColumn() *DatabaseTableColumnApplyConfiguration {
	return &DatabaseTableColumnApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *DatabaseTableColumnApplyConfiguration) WithName(value string) *DatabaseTableColumnApplyConfiguration {
	b.Name = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DatabaseTableColumnApplyConfiguration) WithType(value string) *DatabaseTableColumnApplyConfiguration {
	b.Type = &value
	return b
}
