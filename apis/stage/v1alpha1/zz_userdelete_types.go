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

type UserDeleteInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UserDeleteObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UserDeleteParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// UserDeleteSpec defines the desired state of UserDelete
type UserDeleteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserDeleteParameters `json:"forProvider"`
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
	InitProvider UserDeleteInitParameters `json:"initProvider,omitempty"`
}

// UserDeleteStatus defines the observed state of UserDelete.
type UserDeleteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserDeleteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserDelete is the Schema for the UserDeletes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type UserDelete struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   UserDeleteSpec   `json:"spec"`
	Status UserDeleteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserDeleteList contains a list of UserDeletes
type UserDeleteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserDelete `json:"items"`
}

// Repository type metadata.
var (
	UserDelete_Kind             = "UserDelete"
	UserDelete_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserDelete_Kind}.String()
	UserDelete_KindAPIVersion   = UserDelete_Kind + "." + CRDGroupVersion.String()
	UserDelete_GroupVersionKind = CRDGroupVersion.WithKind(UserDelete_Kind)
)

func init() {
	SchemeBuilder.Register(&UserDelete{}, &UserDeleteList{})
}
