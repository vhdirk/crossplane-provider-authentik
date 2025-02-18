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

type SAMLInitParameters struct {

	// (String)
	AcsURL *string `json:"acsUrl,omitempty" tf:"acs_url,omitempty"`

	// 5.
	// Defaults to `minutes=-5`.
	AssertionValidNotBefore *string `json:"assertionValidNotBefore,omitempty" tf:"assertion_valid_not_before,omitempty"`

	// (String) Defaults to minutes=5.
	// Defaults to `minutes=5`.
	AssertionValidNotOnOrAfter *string `json:"assertionValidNotOnOrAfter,omitempty" tf:"assertion_valid_not_on_or_after,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// (String)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	AuthorizationFlow *string `json:"authorizationFlow,omitempty" tf:"authorization_flow,omitempty"`

	// Reference to a Flow in authentik to populate authorizationFlow.
	// +kubebuilder:validation:Optional
	AuthorizationFlowRef *v1.Reference `json:"authorizationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authorizationFlow.
	// +kubebuilder:validation:Optional
	AuthorizationFlowSelector *v1.Selector `json:"authorizationFlowSelector,omitempty" tf:"-"`

	// (String) Defaults to “.
	// Defaults to “.
	DefaultRelayState *string `json:"defaultRelayState,omitempty" tf:"default_relay_state,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (String)
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String)
	InvalidationFlow *string `json:"invalidationFlow,omitempty" tf:"invalidation_flow,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	NameIDMapping *string `json:"nameIdMapping,omitempty" tf:"name_id_mapping,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (String) Defaults to minutes=86400.
	// Defaults to `minutes=86400`.
	SessionValidNotOnOrAfter *string `json:"sessionValidNotOnOrAfter,omitempty" tf:"session_valid_not_on_or_after,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SignAssertion *bool `json:"signAssertion,omitempty" tf:"sign_assertion,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	SignResponse *bool `json:"signResponse,omitempty" tf:"sign_response,omitempty"`

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

	// (String) Allowed values:
	// Allowed values:
	// - `redirect`
	// - `post`
	// Defaults to `redirect`.
	SpBinding *string `json:"spBinding,omitempty" tf:"sp_binding,omitempty"`

	// (String) Generated.
	// Generated.
	URLSLOPost *string `json:"urlSloPost,omitempty" tf:"url_slo_post,omitempty"`

	// (String) Generated.
	// Generated.
	URLSLORedirect *string `json:"urlSloRedirect,omitempty" tf:"url_slo_redirect,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoInit *string `json:"urlSsoInit,omitempty" tf:"url_sso_init,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoPost *string `json:"urlSsoPost,omitempty" tf:"url_sso_post,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoRedirect *string `json:"urlSsoRedirect,omitempty" tf:"url_sso_redirect,omitempty"`

	// (String)
	VerificationKp *string `json:"verificationKp,omitempty" tf:"verification_kp,omitempty"`
}

type SAMLObservation struct {

	// (String)
	AcsURL *string `json:"acsUrl,omitempty" tf:"acs_url,omitempty"`

	// 5.
	// Defaults to `minutes=-5`.
	AssertionValidNotBefore *string `json:"assertionValidNotBefore,omitempty" tf:"assertion_valid_not_before,omitempty"`

	// (String) Defaults to minutes=5.
	// Defaults to `minutes=5`.
	AssertionValidNotOnOrAfter *string `json:"assertionValidNotOnOrAfter,omitempty" tf:"assertion_valid_not_on_or_after,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// (String)
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// (String)
	AuthorizationFlow *string `json:"authorizationFlow,omitempty" tf:"authorization_flow,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	DefaultRelayState *string `json:"defaultRelayState,omitempty" tf:"default_relay_state,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (String)
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	InvalidationFlow *string `json:"invalidationFlow,omitempty" tf:"invalidation_flow,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	NameIDMapping *string `json:"nameIdMapping,omitempty" tf:"name_id_mapping,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (String) Defaults to minutes=86400.
	// Defaults to `minutes=86400`.
	SessionValidNotOnOrAfter *string `json:"sessionValidNotOnOrAfter,omitempty" tf:"session_valid_not_on_or_after,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SignAssertion *bool `json:"signAssertion,omitempty" tf:"sign_assertion,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	SignResponse *bool `json:"signResponse,omitempty" tf:"sign_response,omitempty"`

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

	// (String) Allowed values:
	// Allowed values:
	// - `redirect`
	// - `post`
	// Defaults to `redirect`.
	SpBinding *string `json:"spBinding,omitempty" tf:"sp_binding,omitempty"`

	// (String) Generated.
	// Generated.
	URLSLOPost *string `json:"urlSloPost,omitempty" tf:"url_slo_post,omitempty"`

	// (String) Generated.
	// Generated.
	URLSLORedirect *string `json:"urlSloRedirect,omitempty" tf:"url_slo_redirect,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoInit *string `json:"urlSsoInit,omitempty" tf:"url_sso_init,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoPost *string `json:"urlSsoPost,omitempty" tf:"url_sso_post,omitempty"`

	// (String) Generated.
	// Generated.
	URLSsoRedirect *string `json:"urlSsoRedirect,omitempty" tf:"url_sso_redirect,omitempty"`

	// (String)
	VerificationKp *string `json:"verificationKp,omitempty" tf:"verification_kp,omitempty"`
}

type SAMLParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	AcsURL *string `json:"acsUrl,omitempty" tf:"acs_url,omitempty"`

	// 5.
	// Defaults to `minutes=-5`.
	// +kubebuilder:validation:Optional
	AssertionValidNotBefore *string `json:"assertionValidNotBefore,omitempty" tf:"assertion_valid_not_before,omitempty"`

	// (String) Defaults to minutes=5.
	// Defaults to `minutes=5`.
	// +kubebuilder:validation:Optional
	AssertionValidNotOnOrAfter *string `json:"assertionValidNotOnOrAfter,omitempty" tf:"assertion_valid_not_on_or_after,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	AuthenticationFlow *string `json:"authenticationFlow,omitempty" tf:"authentication_flow,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/vhdirk/crossplane-provider-authentik/apis/authentik/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)
	// +kubebuilder:validation:Optional
	AuthorizationFlow *string `json:"authorizationFlow,omitempty" tf:"authorization_flow,omitempty"`

	// Reference to a Flow in authentik to populate authorizationFlow.
	// +kubebuilder:validation:Optional
	AuthorizationFlowRef *v1.Reference `json:"authorizationFlowRef,omitempty" tf:"-"`

	// Selector for a Flow in authentik to populate authorizationFlow.
	// +kubebuilder:validation:Optional
	AuthorizationFlowSelector *v1.Selector `json:"authorizationFlowSelector,omitempty" tf:"-"`

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	DefaultRelayState *string `json:"defaultRelayState,omitempty" tf:"default_relay_state,omitempty"`

	// (String) Allowed values:
	// Allowed values:
	// - `http://www.w3.org/2000/09/xmldsig#sha1`
	// - `http://www.w3.org/2001/04/xmlenc#sha256`
	// - `http://www.w3.org/2001/04/xmldsig-more#sha384`
	// - `http://www.w3.org/2001/04/xmlenc#sha512`
	// Defaults to `http://www.w3.org/2001/04/xmlenc#sha256`.
	// +kubebuilder:validation:Optional
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty" tf:"digest_algorithm,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	EncryptionKp *string `json:"encryptionKp,omitempty" tf:"encryption_kp,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	InvalidationFlow *string `json:"invalidationFlow,omitempty" tf:"invalidation_flow,omitempty"`

	// (String) Defaults to authentik.
	// Defaults to `authentik`.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	NameIDMapping *string `json:"nameIdMapping,omitempty" tf:"name_id_mapping,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (String) Defaults to minutes=86400.
	// Defaults to `minutes=86400`.
	// +kubebuilder:validation:Optional
	SessionValidNotOnOrAfter *string `json:"sessionValidNotOnOrAfter,omitempty" tf:"session_valid_not_on_or_after,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	SignAssertion *bool `json:"signAssertion,omitempty" tf:"sign_assertion,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	SignResponse *bool `json:"signResponse,omitempty" tf:"sign_response,omitempty"`

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

	// (String) Allowed values:
	// Allowed values:
	// - `redirect`
	// - `post`
	// Defaults to `redirect`.
	// +kubebuilder:validation:Optional
	SpBinding *string `json:"spBinding,omitempty" tf:"sp_binding,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	URLSLOPost *string `json:"urlSloPost,omitempty" tf:"url_slo_post,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	URLSLORedirect *string `json:"urlSloRedirect,omitempty" tf:"url_slo_redirect,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	URLSsoInit *string `json:"urlSsoInit,omitempty" tf:"url_sso_init,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	URLSsoPost *string `json:"urlSsoPost,omitempty" tf:"url_sso_post,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	URLSsoRedirect *string `json:"urlSsoRedirect,omitempty" tf:"url_sso_redirect,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	VerificationKp *string `json:"verificationKp,omitempty" tf:"verification_kp,omitempty"`
}

// SAMLSpec defines the desired state of SAML
type SAMLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SAMLParameters `json:"forProvider"`
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
	InitProvider SAMLInitParameters `json:"initProvider,omitempty"`
}

// SAMLStatus defines the observed state of SAML.
type SAMLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SAMLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SAML is the Schema for the SAMLs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type SAML struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acsUrl) || (has(self.initProvider) && has(self.initProvider.acsUrl))",message="spec.forProvider.acsUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.invalidationFlow) || (has(self.initProvider) && has(self.initProvider.invalidationFlow))",message="spec.forProvider.invalidationFlow is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SAMLSpec   `json:"spec"`
	Status SAMLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLList contains a list of SAMLs
type SAMLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SAML `json:"items"`
}

// Repository type metadata.
var (
	SAML_Kind             = "SAML"
	SAML_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SAML_Kind}.String()
	SAML_KindAPIVersion   = SAML_Kind + "." + CRDGroupVersion.String()
	SAML_GroupVersionKind = CRDGroupVersion.WithKind(SAML_Kind)
)

func init() {
	SchemeBuilder.Register(&SAML{}, &SAMLList{})
}
