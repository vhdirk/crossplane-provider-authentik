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

type FlowStageBindingInitParameters struct {

	// (Boolean) Evaluate policies during the Flow planning process. Defaults to true.
	// Evaluate policies during the Flow planning process. Defaults to `true`.
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `retry`
	// - `restart`
	// - `restart_with_context`
	// Defaults to `retry`.
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Evaluate policies when the Stage is present to the user. Defaults to false.
	// Evaluate policies when the Stage is present to the user. Defaults to `false`.
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type FlowStageBindingObservation struct {

	// (Boolean) Evaluate policies during the Flow planning process. Defaults to true.
	// Evaluate policies during the Flow planning process. Defaults to `true`.
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `retry`
	// - `restart`
	// - `restart_with_context`
	// Defaults to `retry`.
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Evaluate policies when the Stage is present to the user. Defaults to false.
	// Evaluate policies when the Stage is present to the user. Defaults to `false`.
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

type FlowStageBindingParameters struct {

	// (Boolean) Evaluate policies during the Flow planning process. Defaults to true.
	// Evaluate policies during the Flow planning process. Defaults to `true`.
	// +kubebuilder:validation:Optional
	EvaluateOnPlan *bool `json:"evaluateOnPlan,omitempty" tf:"evaluate_on_plan,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `retry`
	// - `restart`
	// - `restart_with_context`
	// Defaults to `retry`.
	// +kubebuilder:validation:Optional
	InvalidResponseAction *string `json:"invalidResponseAction,omitempty" tf:"invalid_response_action,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	// +kubebuilder:validation:Optional
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Boolean) Evaluate policies when the Stage is present to the user. Defaults to false.
	// Evaluate policies when the Stage is present to the user. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ReEvaluatePolicies *bool `json:"reEvaluatePolicies,omitempty" tf:"re_evaluate_policies,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Stage *string `json:"stage,omitempty" tf:"stage,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`
}

// FlowStageBindingSpec defines the desired state of FlowStageBinding
type FlowStageBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlowStageBindingParameters `json:"forProvider"`
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
	InitProvider FlowStageBindingInitParameters `json:"initProvider,omitempty"`
}

// FlowStageBindingStatus defines the observed state of FlowStageBinding.
type FlowStageBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlowStageBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FlowStageBinding is the Schema for the FlowStageBindings API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type FlowStageBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.order) || (has(self.initProvider) && has(self.initProvider.order))",message="spec.forProvider.order is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stage) || (has(self.initProvider) && has(self.initProvider.stage))",message="spec.forProvider.stage is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.target) || (has(self.initProvider) && has(self.initProvider.target))",message="spec.forProvider.target is a required parameter"
	Spec   FlowStageBindingSpec   `json:"spec"`
	Status FlowStageBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowStageBindingList contains a list of FlowStageBindings
type FlowStageBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowStageBinding `json:"items"`
}

// Repository type metadata.
var (
	FlowStageBinding_Kind             = "FlowStageBinding"
	FlowStageBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlowStageBinding_Kind}.String()
	FlowStageBinding_KindAPIVersion   = FlowStageBinding_Kind + "." + CRDGroupVersion.String()
	FlowStageBinding_GroupVersionKind = CRDGroupVersion.WithKind(FlowStageBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FlowStageBinding{}, &FlowStageBindingList{})
}
