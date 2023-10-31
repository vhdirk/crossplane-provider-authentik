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

type UserInitParameters struct {

	// (String) JSON format expected. Use jsonencode() to pass objects. Defaults to {}.
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes *string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (String)
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (List of String) Generated.
	// Generated.
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to users.
	// Defaults to `users`.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Defaults to internal.
	// Defaults to `internal`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {

	// (String) JSON format expected. Use jsonencode() to pass objects. Defaults to {}.
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	Attributes *string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (String)
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (List of String) Generated.
	// Generated.
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to users.
	// Defaults to `users`.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Defaults to internal.
	// Defaults to `internal`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// (String) JSON format expected. Use jsonencode() to pass objects. Defaults to {}.
	// JSON format expected. Use jsonencode() to pass objects. Defaults to `{}`.
	// +kubebuilder:validation:Optional
	Attributes *string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (List of String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to users.
	// Defaults to `users`.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Defaults to internal.
	// Defaults to `internal`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}