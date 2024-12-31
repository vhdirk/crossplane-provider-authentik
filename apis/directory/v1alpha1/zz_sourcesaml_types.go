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

type SourceSAMLInitParameters struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	AllowIdpInitiated *bool `json:"allowIdpInitiated,omitempty" tf:"allow_idp_initiated,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowRef *v1.Reference `json:"authenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowSelector *v1.Selector `json:"authenticationFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `REDIRECT`
	// - `POST`
	// - `POST_AUTO`
	// Defaults to `REDIRECT`.
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// Reference to a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowRef *v1.Reference `json:"enrollmentFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowSelector *v1.Selector `json:"enrollmentFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String)
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
	// Defaults to `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`.
	NameIDPolicy *string `json:"nameIdPolicy,omitempty" tf:"name_id_policy,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	PreAuthenticationFlow *string `json:"preAuthenticationFlow,omitempty" tf:"pre_authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate preAuthenticationFlow.
	// +kubebuilder:validation:Optional
	PreAuthenticationFlowRef *v1.Reference `json:"preAuthenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate preAuthenticationFlow.
	// +kubebuilder:validation:Optional
	PreAuthenticationFlowSelector *v1.Selector `json:"preAuthenticationFlowSelector,omitempty" tf:"-"`

	// (String)
	SLOURL *string `json:"sloUrl,omitempty" tf:"slo_url,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
	// - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
	// Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`

	// (String)
	SigningKp *string `json:"signingKp,omitempty" tf:"signing_kp,omitempty"`

	// (String)
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// (String) Defaults to days=1.
	// Defaults to `days=1`.
	TemporaryUserDeleteAfter *string `json:"temporaryUserDeleteAfter,omitempty" tf:"temporary_user_delete_after,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourceSAMLObservation struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	AllowIdpInitiated *bool `json:"allowIdpInitiated,omitempty" tf:"allow_idp_initiated,omitempty"`

	// (String)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `REDIRECT`
	// - `POST`
	// - `POST_AUTO`
	// Defaults to `REDIRECT`.
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String)
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String) SAML Metadata Generated.
	// SAML Metadata Generated.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
	// Defaults to `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`.
	NameIDPolicy *string `json:"nameIdPolicy,omitempty" tf:"name_id_policy,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String)
	PreAuthenticationFlow *string `json:"preAuthenticationFlow,omitempty" tf:"pre_authentication_flow,omitempty"`

	// (String)
	SLOURL *string `json:"sloUrl,omitempty" tf:"slo_url,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
	// - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
	// Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`

	// (String)
	SigningKp *string `json:"signingKp,omitempty" tf:"signing_kp,omitempty"`

	// (String)
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// (String) Defaults to days=1.
	// Defaults to `days=1`.
	TemporaryUserDeleteAfter *string `json:"temporaryUserDeleteAfter,omitempty" tf:"temporary_user_delete_after,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourceSAMLParameters struct {

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	AllowIdpInitiated *bool `json:"allowIdpInitiated,omitempty" tf:"allow_idp_initiated,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowRef *v1.Reference `json:"authenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authenticationFlow.
	// +kubebuilder:validation:Optional
	AuthenticationFlowSelector *v1.Selector `json:"authenticationFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `REDIRECT`
	// - `POST`
	// - `POST_AUTO`
	// Defaults to `REDIRECT`.
	// +kubebuilder:validation:Optional
	BindingType *string `json:"bindingType,omitempty" tf:"binding_type,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	// +kubebuilder:validation:Optional
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	EnrollmentFlow *string `json:"enrollmentFlow,omitempty" tf:"enrollment_flow,omitempty"`

	// Reference to a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowRef *v1.Reference `json:"enrollmentFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate enrollmentFlow.
	// +kubebuilder:validation:Optional
	EnrollmentFlowSelector *v1.Selector `json:"enrollmentFlowSelector,omitempty" tf:"-"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `name_link`
	// - `name_deny`
	// Defaults to `identifier`.
	// +kubebuilder:validation:Optional
	GroupMatchingMode *string `json:"groupMatchingMode,omitempty" tf:"group_matching_mode,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`
	// - `urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:WindowsDomainQualifiedName`
	// - `urn:oasis:names:tc:SAML:2.0:nameid-format:transient`
	// Defaults to `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`.
	// +kubebuilder:validation:Optional
	NameIDPolicy *string `json:"nameIdPolicy,omitempty" tf:"name_id_policy,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `all`
	// - `any`
	// Defaults to `any`.
	// +kubebuilder:validation:Optional
	PolicyEngineMode *string `json:"policyEngineMode,omitempty" tf:"policy_engine_mode,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	PreAuthenticationFlow *string `json:"preAuthenticationFlow,omitempty" tf:"pre_authentication_flow,omitempty"`

	// Reference to a Flow in authentik to populate preAuthenticationFlow.
	// +kubebuilder:validation:Optional
	PreAuthenticationFlowRef *v1.Reference `json:"preAuthenticationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate preAuthenticationFlow.
	// +kubebuilder:validation:Optional
	PreAuthenticationFlowSelector *v1.Selector `json:"preAuthenticationFlowSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	SLOURL *string `json:"sloUrl,omitempty" tf:"slo_url,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#rsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#rsa-sha512`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha1`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha384`
	// - `http://www.w3.org/2001/04/xmldsig-more#ecdsa-sha512`
	// - `http://www.w3.org/2000/09/xmldsig#dsa-sha1`
	// Defaults to `http://www.w3.org/2001/04/xmldsig-more#rsa-sha256`.
	// +kubebuilder:validation:Optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SigningKp *string `json:"signingKp,omitempty" tf:"signing_kp,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`

	// (String) Defaults to days=1.
	// Defaults to `days=1`.
	// +kubebuilder:validation:Optional
	TemporaryUserDeleteAfter *string `json:"temporaryUserDeleteAfter,omitempty" tf:"temporary_user_delete_after,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `identifier`
	// - `email_link`
	// - `email_deny`
	// - `username_link`
	// - `username_deny`
	// Defaults to `identifier`.
	// +kubebuilder:validation:Optional
	UserMatchingMode *string `json:"userMatchingMode,omitempty" tf:"user_matching_mode,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	// +kubebuilder:validation:Optional
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

// SourceSAMLSpec defines the desired state of SourceSAML
type SourceSAMLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SourceSAMLParameters `json:"forProvider"`
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
	InitProvider SourceSAMLInitParameters `json:"initProvider,omitempty"`
}

// SourceSAMLStatus defines the observed state of SourceSAML.
type SourceSAMLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SourceSAMLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SourceSAML is the Schema for the SourceSAMLs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type SourceSAML struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ssoUrl) || (has(self.initProvider) && has(self.initProvider.ssoUrl))",message="spec.forProvider.ssoUrl is a required parameter"
	Spec   SourceSAMLSpec   `json:"spec"`
	Status SourceSAMLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourceSAMLList contains a list of SourceSAMLs
type SourceSAMLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourceSAML `json:"items"`
}

// Repository type metadata.
var (
	SourceSAML_Kind             = "SourceSAML"
	SourceSAML_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SourceSAML_Kind}.String()
	SourceSAML_KindAPIVersion   = SourceSAML_Kind + "." + CRDGroupVersion.String()
	SourceSAML_GroupVersionKind = CRDGroupVersion.WithKind(SourceSAML_Kind)
)

func init() {
	SchemeBuilder.Register(&SourceSAML{}, &SourceSAMLList{})
}
