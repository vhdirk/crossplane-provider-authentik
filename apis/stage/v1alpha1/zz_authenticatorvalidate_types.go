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

type AuthenticatorValidateInitParameters struct {

	// (List of String)
	ConfigurationStages []*string `json:"configurationStages,omitempty" tf:"configuration_stages,omitempty"`

	// (List of String)
	DeviceClasses []*string `json:"deviceClasses,omitempty" tf:"device_classes,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	LastAuthThreshold *string `json:"lastAuthThreshold,omitempty" tf:"last_auth_threshold,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `skip`
	// - `deny`
	// - `configure`
	NotConfiguredAction *string `json:"notConfiguredAction,omitempty" tf:"not_configured_action,omitempty"`

	// (List of String)
	WebauthnAllowedDeviceTypes []*string `json:"webauthnAllowedDeviceTypes,omitempty" tf:"webauthn_allowed_device_types,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `required`
	// - `preferred`
	// - `discouraged`
	// Defaults to `preferred`.
	WebauthnUserVerification *string `json:"webauthnUserVerification,omitempty" tf:"webauthn_user_verification,omitempty"`
}

type AuthenticatorValidateObservation struct {

	// (List of String)
	ConfigurationStages []*string `json:"configurationStages,omitempty" tf:"configuration_stages,omitempty"`

	// (List of String)
	DeviceClasses []*string `json:"deviceClasses,omitempty" tf:"device_classes,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	LastAuthThreshold *string `json:"lastAuthThreshold,omitempty" tf:"last_auth_threshold,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `skip`
	// - `deny`
	// - `configure`
	NotConfiguredAction *string `json:"notConfiguredAction,omitempty" tf:"not_configured_action,omitempty"`

	// (List of String)
	WebauthnAllowedDeviceTypes []*string `json:"webauthnAllowedDeviceTypes,omitempty" tf:"webauthn_allowed_device_types,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `required`
	// - `preferred`
	// - `discouraged`
	// Defaults to `preferred`.
	WebauthnUserVerification *string `json:"webauthnUserVerification,omitempty" tf:"webauthn_user_verification,omitempty"`
}

type AuthenticatorValidateParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	ConfigurationStages []*string `json:"configurationStages,omitempty" tf:"configuration_stages,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	DeviceClasses []*string `json:"deviceClasses,omitempty" tf:"device_classes,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	// +kubebuilder:validation:Optional
	LastAuthThreshold *string `json:"lastAuthThreshold,omitempty" tf:"last_auth_threshold,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `skip`
	// - `deny`
	// - `configure`
	// +kubebuilder:validation:Optional
	NotConfiguredAction *string `json:"notConfiguredAction,omitempty" tf:"not_configured_action,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	WebauthnAllowedDeviceTypes []*string `json:"webauthnAllowedDeviceTypes,omitempty" tf:"webauthn_allowed_device_types,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `required`
	// - `preferred`
	// - `discouraged`
	// Defaults to `preferred`.
	// +kubebuilder:validation:Optional
	WebauthnUserVerification *string `json:"webauthnUserVerification,omitempty" tf:"webauthn_user_verification,omitempty"`
}

// AuthenticatorValidateSpec defines the desired state of AuthenticatorValidate
type AuthenticatorValidateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuthenticatorValidateParameters `json:"forProvider"`
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
	InitProvider AuthenticatorValidateInitParameters `json:"initProvider,omitempty"`
}

// AuthenticatorValidateStatus defines the observed state of AuthenticatorValidate.
type AuthenticatorValidateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuthenticatorValidateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuthenticatorValidate is the Schema for the AuthenticatorValidates API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type AuthenticatorValidate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notConfiguredAction) || (has(self.initProvider) && has(self.initProvider.notConfiguredAction))",message="spec.forProvider.notConfiguredAction is a required parameter"
	Spec   AuthenticatorValidateSpec   `json:"spec"`
	Status AuthenticatorValidateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthenticatorValidateList contains a list of AuthenticatorValidates
type AuthenticatorValidateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthenticatorValidate `json:"items"`
}

// Repository type metadata.
var (
	AuthenticatorValidate_Kind             = "AuthenticatorValidate"
	AuthenticatorValidate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthenticatorValidate_Kind}.String()
	AuthenticatorValidate_KindAPIVersion   = AuthenticatorValidate_Kind + "." + CRDGroupVersion.String()
	AuthenticatorValidate_GroupVersionKind = CRDGroupVersion.WithKind(AuthenticatorValidate_Kind)
)

func init() {
	SchemeBuilder.Register(&AuthenticatorValidate{}, &AuthenticatorValidateList{})
}
