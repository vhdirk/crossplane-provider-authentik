package propertymapping

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const shortGroup = "propertymapping"

var GetOptionalFieldValueInto = func(paved *fieldpath.Paved, path string, out any) error {
	if err := paved.GetValueInto(path, &out); err != nil {
		if !fieldpath.IsNotFound(err) {
			return err
		}
	}
	return nil
}

var DeleteOptionalField = func(paved *fieldpath.Paved, path string) error {
	if err := paved.DeleteField(path); err != nil {
		if !fieldpath.IsNotFound(err) {
			return err
		}
	}
	return nil
}

var HasField = func(paved *fieldpath.Paved, path string, defaultValue any) (bool, error) {
	val, err := paved.GetValue(path)

	if val != nil && val != defaultValue {
		return true, nil
	}

	if err != nil && fieldpath.IsNotFound(err) {
		return false, nil
	}

	return false, err
}

var forProviderField = func(field string) string {
	return fmt.Sprintf("spec.forProvider.%s", field)
}

// var MappingProviderInitializer = func(conflictsWith []string) config.NewInitializerFn {
// 	return func(client client.Client) managed.Initializer {
// 		return managed.InitializerFn(func(ctx context.Context, mg xpresource.Managed) error {

// 			paved, err := fieldpath.PaveObject(mg)
// 			if err != nil {
// 				return errors.Wrap(err, "cannot pave object")
// 			}

// 			managedFieldPath := forProviderField("managed")
// 			hasManaged, err := HasField(paved, managedFieldPath, "")
// 			if err != nil {
// 				return errors.Wrapf(err, "cannot unmarshal field 'managed'")
// 			}

// 			for _, fieldName := range conflictsWith {
// 				fieldPath := forProviderField(fieldName)
// 				hasField, err := HasField(paved, fieldPath, "")
// 				if err != nil {
// 					return errors.Wrapf(err, "cannot unmarshal field '%s'", fieldName)
// 				}

// 				if hasManaged && hasField {
// 					return errors.New("The field '%s' is mutually exclusive with 'managed'")
// 				}
// 			}

// 			if hasManaged {
// 				for _, fieldName := range conflictsWith {
// 					fieldPath := forProviderField(fieldName)

// 					if err := DeleteOptionalField(paved, fieldPath); err != nil {
// 						return errors.Wrapf(err, "could not delete field: %s", fieldName)
// 					}
// 				}
// 			}

// 			pavedData, err := paved.MarshalJSON()
// 			if err != nil {
// 				return errors.Wrapf(err, "could not marshal modified spec into JSON")
// 			}
// 			if err := json.Unmarshal(pavedData, mg); err != nil {
// 				return errors.Wrapf(err, "could not unmarshal modified role spec into managed resource interface")
// 			}
// 			if err := client.Update(ctx, mg); err != nil {
// 				return errors.Wrapf(err, "could not update managed resource")
// 			}
// 			return nil
// 		})
// 	}
// }

// Configure configures the propertymapping provider.
func Configure(p *config.Provider) {
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Notification"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderGoogleWorkspace"
		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderMicrosoftEntra"
		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderRAC"

		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

		// r.TerraformResource.Schema["name"] = &schema.Schema{
		// 	Type:     schema.TypeString,
		// 	Optional: true,
		// 	Computed: false,
		// }

		// r.InitializerFns = append(r.InitializerFns, MappingProviderInitializer([]string{"name", "expression"}))
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderRadius"

		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

		// r.InitializerFns = append(r.InitializerFns, MappingProviderInitializer([]string{"name", "settings", "expression"}))

	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderSAML"
		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

		// r.InitializerFns = append(r.InitializerFns, MappingProviderInitializer([]string{"name", "saml_name", "friendly_name", "expression"}))
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderSCIM"
		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

		// r.InitializerFns = append(r.InitializerFns, MappingProviderInitializer([]string{"name", "expression"}))
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderScope"
		r.TerraformResource.Schema["managed"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: false,
		}

		for _, field := range []string{"name", "scope_name", "description", "expression"} {
			if s, ok := r.TerraformResource.Schema[field]; ok {
				s.Optional = true
				s.Computed = false
			}
		}
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceKerberos"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceLDAP"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceOAuth"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_plex", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourcePlex"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSAML"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSCIM"
	})

}
