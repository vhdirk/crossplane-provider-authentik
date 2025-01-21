package provider

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/vhdirk/crossplane-provider-authentik/config/base"
)

const shortGroup = "provider"

type ProviderReference struct {
	// Name is required
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type ProtocolProviderParameters struct {
	// Only one of these fields can be specified
	// +kubebuilder:validation:XValidation:rule="has(self.oauth2Ref) + has(self.ldapRef) + has(self.samlRef) + has(self.scimRef) + has(self.proxyRef) + has(self.radiusRef) + has(self.racRef) + has(self.microsoftEntraRef) + has(self.googleWorkspaceRef) == 1",message="Exactly one provider reference must be specified"

	// OAuth2 provider reference
	// +optional
	OAuth2Ref *ProviderReference `json:"oauth2Ref,omitempty"`

	// LDAP provider reference
	// +optional
	LDAPRef *ProviderReference `json:"ldapRef,omitempty"`

	// SAML provider reference
	// +optional
	SAMLRef *ProviderReference `json:"samlRef,omitempty"`

	// ... other refs
}

type ProtocolProviderObservation struct {
	// ID of the referenced provider
	ID string `json:"id,omitempty"`
}

type ProtocolProviderStatus struct {
	AtProvider ProtocolProviderObservation `json:"atProvider,omitempty"`
}

func ConfigureProtocolProvider(p *config.Provider) {
	p.AddResourceConfigurator("authentik_protocol_provider", func(r *config.Resource) {
		r.Kind = "ProtocolProvider"
		r.ShortGroup = "authentik"

		r.References = map[string]config.Reference{
			"oauth2Ref": {
				TerraformName: "authentik_provider_oauth2",
				Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
			},
			"ldapRef": {
				TerraformName: "authentik_provider_ldap",
				Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
			},
			"samlRef": {
				TerraformName: "authentik_provider_saml",
				Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
			},
			// ... other refs
		}

		// Ensure the ID is stored in status
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if id, ok := attr["id"].(string); ok {
				conn["id"] = []byte(id)
			}
			return conn, nil
		}
	})
}

// Configure configures the provider provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"

		r.References["authorization_flow"] = base.FlowRef
	})

	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OAuth2"

		r.References["authorization_flow"] = base.FlowRef
		r.References["authentication_flow"] = base.FlowRef
		r.References["invalidation_flow"] = base.FlowRef
		r.References["signing_key"] = base.CertificateKeyPairRef
		r.References["encryption_key"] = base.CertificateKeyPairRef

		// r.References["property_mappings"] = base Prop
	})
	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Proxy"

		r.References["authorization_flow"] = base.FlowRef
		r.References["authentication_flow"] = base.FlowRef
		r.References["invalidation_flow"] = base.FlowRef

	})
	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"

		r.References["authentication_flow"] = base.FlowRef
		r.References["authorization_flow"] = base.FlowRef

	})
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Radius"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"

		r.References["authorization_flow"] = base.FlowRef
	})
}
