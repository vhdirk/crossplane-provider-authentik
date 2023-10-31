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

type StageBindingInitParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) Defaults to retry.
	// Defaults to `retry`.
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type StageBindingObservation struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Defaults to retry.
	// Defaults to `retry`.
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type StageBindingParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) Defaults to retry.
	// Defaults to `retry`.
	// +kubebuilder:validation:Optional
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	// +kubebuilder:validation:Optional
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

// StageBindingSpec defines the desired state of StageBinding
type StageBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StageBindingParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider StageBindingInitParameters `json:"initProvider,omitempty"`
}

// StageBindingStatus defines the observed state of StageBinding.
type StageBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StageBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StageBinding is the Schema for the StageBindings API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type StageBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.order) || (has(self.initProvider) && has(self.initProvider.order))",message="spec.forProvider.order is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stage) || (has(self.initProvider) && has(self.initProvider.stage))",message="spec.forProvider.stage is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.target) || (has(self.initProvider) && has(self.initProvider.target))",message="spec.forProvider.target is a required parameter"
	Spec   StageBindingSpec   `json:"spec"`
	Status StageBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageBindingList contains a list of StageBindings
type StageBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StageBinding `json:"items"`
}

// Repository type metadata.
var (
	StageBinding_Kind             = "StageBinding"
	StageBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StageBinding_Kind}.String()
	StageBinding_KindAPIVersion   = StageBinding_Kind + "." + CRDGroupVersion.String()
	StageBinding_GroupVersionKind = CRDGroupVersion.WithKind(StageBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&StageBinding{}, &StageBindingList{})
}