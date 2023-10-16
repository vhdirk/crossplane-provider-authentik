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

type ApplicationInitParameters struct {

	// (List of Number)
	BackchannelProviders []*float64 `json:"backchannelProviders,omitempty" tf:"backchannel_providers,omitempty"`

	// (String)
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String)
	MetaDescription *string `json:"metaDescription,omitempty" tf:"meta_description,omitempty"`

	// (String)
	MetaIcon *string `json:"metaIcon,omitempty" tf:"meta_icon,omitempty"`

	// (String)
	MetaLaunchURL *string `json:"metaLaunchUrl,omitempty" tf:"meta_launch_url,omitempty"`

	// (String)
	MetaPublisher *string `json:"metaPublisher,omitempty" tf:"meta_publisher,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	OpenInNewTab *bool `json:"openInNewTab,omitempty" tf:"open_in_new_tab,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Number)
	ProtocolProvider *float64 `json:"protocolProvider,omitempty" tf:"protocol_provider,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ApplicationObservation struct {

	// (List of Number)
	BackchannelProviders []*float64 `json:"backchannelProviders,omitempty" tf:"backchannel_providers,omitempty"`

	// (String)
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	MetaDescription *string `json:"metaDescription,omitempty" tf:"meta_description,omitempty"`

	// (String)
	MetaIcon *string `json:"metaIcon,omitempty" tf:"meta_icon,omitempty"`

	// (String)
	MetaLaunchURL *string `json:"metaLaunchUrl,omitempty" tf:"meta_launch_url,omitempty"`

	// (String)
	MetaPublisher *string `json:"metaPublisher,omitempty" tf:"meta_publisher,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	OpenInNewTab *bool `json:"openInNewTab,omitempty" tf:"open_in_new_tab,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Number)
	ProtocolProvider *float64 `json:"protocolProvider,omitempty" tf:"protocol_provider,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type ApplicationParameters struct {

	// (List of Number)
	// +kubebuilder:validation:Optional
	BackchannelProviders []*float64 `json:"backchannelProviders,omitempty" tf:"backchannel_providers,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MetaDescription *string `json:"metaDescription,omitempty" tf:"meta_description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MetaIcon *string `json:"metaIcon,omitempty" tf:"meta_icon,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MetaLaunchURL *string `json:"metaLaunchUrl,omitempty" tf:"meta_launch_url,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MetaPublisher *string `json:"metaPublisher,omitempty" tf:"meta_publisher,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	OpenInNewTab *bool `json:"openInNewTab,omitempty" tf:"open_in_new_tab,omitempty"`

	// (String) Defaults to any.
	// Defaults to `any`.
	// +kubebuilder:validation:Optional
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	ProtocolProvider *float64 `json:"protocolProvider,omitempty" tf:"protocol_provider,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.slug) || (has(self.initProvider) && has(self.initProvider.slug))",message="spec.forProvider.slug is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
