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

type ReputationInitParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	CheckIP *bool `json:"checkIp,omitempty" tf:"check_ip,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	CheckUsername *bool `json:"checkUsername,omitempty" tf:"check_username,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Defaults to 10.
	// Defaults to `10`.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ReputationObservation struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	CheckIP *bool `json:"checkIp,omitempty" tf:"check_ip,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	CheckUsername *bool `json:"checkUsername,omitempty" tf:"check_username,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Defaults to 10.
	// Defaults to `10`.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

type ReputationParameters struct {

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	CheckIP *bool `json:"checkIp,omitempty" tf:"check_ip,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	CheckUsername *bool `json:"checkUsername,omitempty" tf:"check_username,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	ExecutionLogging *bool `json:"executionLogging,omitempty" tf:"execution_logging,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Defaults to 10.
	// Defaults to `10`.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`
}

// ReputationSpec defines the desired state of Reputation
type ReputationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReputationParameters `json:"forProvider"`
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
	InitProvider ReputationInitParameters `json:"initProvider,omitempty"`
}

// ReputationStatus defines the observed state of Reputation.
type ReputationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReputationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reputation is the Schema for the Reputations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Reputation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ReputationSpec   `json:"spec"`
	Status ReputationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReputationList contains a list of Reputations
type ReputationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reputation `json:"items"`
}

// Repository type metadata.
var (
	Reputation_Kind             = "Reputation"
	Reputation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reputation_Kind}.String()
	Reputation_KindAPIVersion   = Reputation_Kind + "." + CRDGroupVersion.String()
	Reputation_GroupVersionKind = CRDGroupVersion.WithKind(Reputation_Kind)
)

func init() {
	SchemeBuilder.Register(&Reputation{}, &ReputationList{})
}