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

type UserLoginInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	RememberMeOffset *string `json:"rememberMeOffset,omitempty" tf:"remember_me_offset,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	TerminateOtherSessions *bool `json:"terminateOtherSessions,omitempty" tf:"terminate_other_sessions,omitempty"`
}

type UserLoginObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	RememberMeOffset *string `json:"rememberMeOffset,omitempty" tf:"remember_me_offset,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	TerminateOtherSessions *bool `json:"terminateOtherSessions,omitempty" tf:"terminate_other_sessions,omitempty"`
}

type UserLoginParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	// +kubebuilder:validation:Optional
	RememberMeOffset *string `json:"rememberMeOffset,omitempty" tf:"remember_me_offset,omitempty"`

	// (String) Defaults to seconds=0.
	// Defaults to `seconds=0`.
	// +kubebuilder:validation:Optional
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	TerminateOtherSessions *bool `json:"terminateOtherSessions,omitempty" tf:"terminate_other_sessions,omitempty"`
}

// UserLoginSpec defines the desired state of UserLogin
type UserLoginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserLoginParameters `json:"forProvider"`
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
	InitProvider UserLoginInitParameters `json:"initProvider,omitempty"`
}

// UserLoginStatus defines the observed state of UserLogin.
type UserLoginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserLoginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserLogin is the Schema for the UserLogins API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type UserLogin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   UserLoginSpec   `json:"spec"`
	Status UserLoginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserLoginList contains a list of UserLogins
type UserLoginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserLogin `json:"items"`
}

// Repository type metadata.
var (
	UserLogin_Kind             = "UserLogin"
	UserLogin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserLogin_Kind}.String()
	UserLogin_KindAPIVersion   = UserLogin_Kind + "." + CRDGroupVersion.String()
	UserLogin_GroupVersionKind = CRDGroupVersion.WithKind(UserLogin_Kind)
)

func init() {
	SchemeBuilder.Register(&UserLogin{}, &UserLoginList{})
}