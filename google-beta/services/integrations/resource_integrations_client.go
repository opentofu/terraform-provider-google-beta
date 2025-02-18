// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/integrations/Client.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package integrations

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceIntegrationsClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegrationsClientCreate,
		Read:   resourceIntegrationsClientRead,
		Delete: resourceIntegrationsClientDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIntegrationsClientImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location in which client needs to be provisioned.`,
			},
			"cloud_kms_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Cloud KMS config for AuthModule to encrypt/decrypt credentials.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `A Cloud KMS key is a named object containing one or more key versions, along
with metadata for the key. A key exists on exactly one key ring tied to a
specific location.`,
						},
						"kms_location": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Location name of the key ring, e.g. "us-west1".`,
						},
						"kms_ring": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `A key ring organizes keys in a specific Google Cloud location and allows you to
manage access control on groups of keys. A key ring's name does not need to be
unique across a Google Cloud project, but must be unique within a given location.`,
						},
						"key_version": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Each version of a key contains key material used for encryption or signing.
A key's version is represented by an integer, starting at 1. To decrypt data
or verify a signature, you must use the same key version that was used to
encrypt or sign the data.`,
						},
						"kms_project_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The Google Cloud project id of the project where the kms key stored. If empty,
the kms key is stored at the same project as customer's project and ecrypted
with CMEK, otherwise, the kms key is stored in the tenant project and
encrypted with GMEK.`,
						},
					},
				},
			},
			"create_sample_integrations": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Indicates if sample integrations should be created along with provisioning.`,
			},
			"run_as_service_account": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `User input run-as service account, if empty, will bring up a new default service account.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIntegrationsClientCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	cloudKmsConfigProp, err := expandIntegrationsClientCloudKmsConfig(d.Get("cloud_kms_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cloud_kms_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudKmsConfigProp)) && (ok || !reflect.DeepEqual(v, cloudKmsConfigProp)) {
		obj["cloudKmsConfig"] = cloudKmsConfigProp
	}
	createSampleIntegrationsProp, err := expandIntegrationsClientCreateSampleIntegrations(d.Get("create_sample_integrations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("create_sample_integrations"); !tpgresource.IsEmptyValue(reflect.ValueOf(createSampleIntegrationsProp)) && (ok || !reflect.DeepEqual(v, createSampleIntegrationsProp)) {
		obj["createSampleIntegrations"] = createSampleIntegrationsProp
	}
	runAsServiceAccountProp, err := expandIntegrationsClientRunAsServiceAccount(d.Get("run_as_service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("run_as_service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(runAsServiceAccountProp)) && (ok || !reflect.DeepEqual(v, runAsServiceAccountProp)) {
		obj["runAsServiceAccount"] = runAsServiceAccountProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "Client/{{location}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationsBasePath}}projects/{{project}}/locations/{{location}}/clients:provision")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Client: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Client: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// Translate `createSampleIntegrations` to `createSampleWorkflows`
	if val, ok := obj["createSampleIntegrations"]; ok {
		delete(obj, "createSampleIntegrations")
		obj["createSampleWorkflows"] = val
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Client: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Client %q: %#v", d.Id(), res)

	return resourceIntegrationsClientRead(d, meta)
}

func resourceIntegrationsClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationsBasePath}}projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Client: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IntegrationsClient %q", d.Id()))
	}

	res, err = resourceIntegrationsClientDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IntegrationsClient because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Client: %s", err)
	}

	return nil
}

func resourceIntegrationsClientDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Client: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "Client/{{location}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationsBasePath}}projects/{{project}}/locations/{{location}}/clients:deprovision")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Client %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Client")
	}

	log.Printf("[DEBUG] Finished deleting Client %q: %#v", d.Id(), res)
	return nil
}

func resourceIntegrationsClientImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/clients$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)$",
		"^(?P<location>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/clients")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandIntegrationsClientCloudKmsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsLocation, err := expandIntegrationsClientCloudKmsConfigKmsLocation(original["kms_location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsLocation"] = transformedKmsLocation
	}

	transformedKmsRing, err := expandIntegrationsClientCloudKmsConfigKmsRing(original["kms_ring"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsRing); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsRing"] = transformedKmsRing
	}

	transformedKey, err := expandIntegrationsClientCloudKmsConfigKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	transformedKeyVersion, err := expandIntegrationsClientCloudKmsConfigKeyVersion(original["key_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeyVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keyVersion"] = transformedKeyVersion
	}

	transformedKmsProjectId, err := expandIntegrationsClientCloudKmsConfigKmsProjectId(original["kms_project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsProjectId"] = transformedKmsProjectId
	}

	return transformed, nil
}

func expandIntegrationsClientCloudKmsConfigKmsLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientCloudKmsConfigKmsRing(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientCloudKmsConfigKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientCloudKmsConfigKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientCloudKmsConfigKmsProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientCreateSampleIntegrations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationsClientRunAsServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceIntegrationsClientDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Since Client resource doesnt have any properties,
	// Adding this decoder as placeholder else the linter will
	// complain that the returned `res` is never used afterwards.
	return res, nil
}
