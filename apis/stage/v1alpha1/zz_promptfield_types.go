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

type PromptFieldInitParameters struct {

	// (String)
	FieldKey *string `json:"fieldKey,omitempty" tf:"field_key,omitempty"`

	// (String)
	InitialValue *string `json:"initialValue,omitempty" tf:"initial_value,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	InitialValueExpression *bool `json:"initialValueExpression,omitempty" tf:"initial_value_expression,omitempty"`

	// (String)
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String)
	Placeholder *string `json:"placeholder,omitempty" tf:"placeholder,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	PlaceholderExpression *bool `json:"placeholderExpression,omitempty" tf:"placeholder_expression,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	SubText *string `json:"subText,omitempty" tf:"sub_text,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `text`
	// - `text_area`
	// - `text_read_only`
	// - `text_area_read_only`
	// - `username`
	// - `email`
	// - `password`
	// - `number`
	// - `checkbox`
	// - `radio-button-group`
	// - `dropdown`
	// - `date`
	// - `date-time`
	// - `file`
	// - `separator`
	// - `hidden`
	// - `static`
	// - `ak-locale`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PromptFieldObservation struct {

	// (String)
	FieldKey *string `json:"fieldKey,omitempty" tf:"field_key,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	InitialValue *string `json:"initialValue,omitempty" tf:"initial_value,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	InitialValueExpression *bool `json:"initialValueExpression,omitempty" tf:"initial_value_expression,omitempty"`

	// (String)
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String)
	Placeholder *string `json:"placeholder,omitempty" tf:"placeholder,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	PlaceholderExpression *bool `json:"placeholderExpression,omitempty" tf:"placeholder_expression,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	SubText *string `json:"subText,omitempty" tf:"sub_text,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `text`
	// - `text_area`
	// - `text_read_only`
	// - `text_area_read_only`
	// - `username`
	// - `email`
	// - `password`
	// - `number`
	// - `checkbox`
	// - `radio-button-group`
	// - `dropdown`
	// - `date`
	// - `date-time`
	// - `file`
	// - `separator`
	// - `hidden`
	// - `static`
	// - `ak-locale`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PromptFieldParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	FieldKey *string `json:"fieldKey,omitempty" tf:"field_key,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	InitialValue *string `json:"initialValue,omitempty" tf:"initial_value,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	InitialValueExpression *bool `json:"initialValueExpression,omitempty" tf:"initial_value_expression,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Placeholder *string `json:"placeholder,omitempty" tf:"placeholder,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	PlaceholderExpression *bool `json:"placeholderExpression,omitempty" tf:"placeholder_expression,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	SubText *string `json:"subText,omitempty" tf:"sub_text,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `text`
	// - `text_area`
	// - `text_read_only`
	// - `text_area_read_only`
	// - `username`
	// - `email`
	// - `password`
	// - `number`
	// - `checkbox`
	// - `radio-button-group`
	// - `dropdown`
	// - `date`
	// - `date-time`
	// - `file`
	// - `separator`
	// - `hidden`
	// - `static`
	// - `ak-locale`
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// PromptFieldSpec defines the desired state of PromptField
type PromptFieldSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PromptFieldParameters `json:"forProvider"`
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
	InitProvider PromptFieldInitParameters `json:"initProvider,omitempty"`
}

// PromptFieldStatus defines the observed state of PromptField.
type PromptFieldStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PromptFieldObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PromptField is the Schema for the PromptFields API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type PromptField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fieldKey) || (has(self.initProvider) && has(self.initProvider.fieldKey))",message="spec.forProvider.fieldKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   PromptFieldSpec   `json:"spec"`
	Status PromptFieldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PromptFieldList contains a list of PromptFields
type PromptFieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PromptField `json:"items"`
}

// Repository type metadata.
var (
	PromptField_Kind             = "PromptField"
	PromptField_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PromptField_Kind}.String()
	PromptField_KindAPIVersion   = PromptField_Kind + "." + CRDGroupVersion.String()
	PromptField_GroupVersionKind = CRDGroupVersion.WithKind(PromptField_Kind)
)

func init() {
	SchemeBuilder.Register(&PromptField{}, &PromptFieldList{})
}
