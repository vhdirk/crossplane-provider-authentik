/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ScopeInitParameters struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

type ScopeObservation struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

type ScopeParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ScopeName *string `json:"scopeName,omitempty" tf:"scope_name,omitempty"`
}

// ScopeSpec defines the desired state of Scope
type ScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScopeParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ScopeInitParameters `json:"initProvider,omitempty"`
}

// ScopeStatus defines the observed state of Scope.
type ScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Scope is the Schema for the Scopes API. Manage Scope Provider Property mappings
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expression) || (has(self.initProvider) && has(self.initProvider.expression))",message="spec.forProvider.expression is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeName) || (has(self.initProvider) && has(self.initProvider.scopeName))",message="spec.forProvider.scopeName is a required parameter"
	Spec   ScopeSpec   `json:"spec"`
	Status ScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScopeList contains a list of Scopes
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scope `json:"items"`
}

// Repository type metadata.
var (
	Scope_Kind             = "Scope"
	Scope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Scope_Kind}.String()
	Scope_KindAPIVersion   = Scope_Kind + "." + CRDGroupVersion.String()
	Scope_GroupVersionKind = CRDGroupVersion.WithKind(Scope_Kind)
)

func init() {
	SchemeBuilder.Register(&Scope{}, &ScopeList{})
}
