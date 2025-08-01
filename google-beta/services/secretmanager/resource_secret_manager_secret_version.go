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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/secretmanager/SecretVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package secretmanager

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"google.golang.org/api/googleapi"
)

func setEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) error {
	name := d.Get("name").(string)
	if name == "" {
		return nil
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}
	if v == true {
		url = fmt.Sprintf("%s:enable", url)
	} else {
		url = fmt.Sprintf("%s:disable", url)
	}

	parts := strings.Split(name, "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	return err
}

func ResourceSecretManagerSecretVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretManagerSecretVersionCreate,
		Read:   resourceSecretManagerSecretVersionRead,
		Update: resourceSecretManagerSecretVersionUpdate,
		Delete: resourceSecretManagerSecretVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecretManagerSecretVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"secret_data_wo_version": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `Triggers update of secret data write-only. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)`,
				Default:     0,
			},
			"secret_data": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   `The secret data. Must be no larger than 64KiB.`,
				Sensitive:     true,
				ConflictsWith: []string{},
			},
			"secret_data_wo": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   `The secret data. Must be no larger than 64KiB. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)`,
				WriteOnly:     true,
				ConflictsWith: []string{"secret_data"},
				RequiredWith:  []string{},
			},

			"secret": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Secret Manager secret resource`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `The current state of the SecretVersion.`,
				Default:     true,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Secret was created.`,
			},
			"destroy_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Secret was destroyed. Only present if state is DESTROYED.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the SecretVersion. Format:
'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'`,
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The version of the Secret.`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The deletion policy for the secret version. Setting 'ABANDON' allows the resource
to be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be
disabled rather than deleted. Default is 'DELETE'. Possible values are:
  * DELETE
  * DISABLE
  * ABANDON`,
				Default: "DELETE",
			},
			"is_secret_data_base64": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: `If set to 'true', the secret data is expected to be base64-encoded string and would be sent as is.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecretManagerSecretVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerSecretVersionPayload(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) {
		obj["payload"] = payloadProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{secret}}:addVersion")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecretVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
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
		return fmt.Errorf("Error creating SecretVersion: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceSecretManagerSecretVersionPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	err = setEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished creating SecretVersion %q: %#v", d.Id(), res)

	return resourceSecretManagerSecretVersionRead(d, meta)
}

func resourceSecretManagerSecretVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// Explicitly set the field to default value if unset
	if _, ok := d.GetOkExists("is_secret_data_base64"); !ok {
		if err := d.Set("is_secret_data_base64", false); err != nil {
			return fmt.Errorf("Error setting is_secret_data_base64: %s", err)
		}
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecretManagerSecretVersion %q", d.Id()))
	}

	res, err = resourceSecretManagerSecretVersionDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SecretManagerSecretVersion because it no longer exists.")
		d.SetId("")
		return nil
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "DELETE"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}

	if err := d.Set("enabled", flattenSecretManagerSecretVersionEnabled(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("name", flattenSecretManagerSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("version", flattenSecretManagerSecretVersionVersion(res["version"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("create_time", flattenSecretManagerSecretVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("destroy_time", flattenSecretManagerSecretVersionDestroyTime(res["destroyTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenSecretManagerSecretVersionPayload(res["payload"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading SecretVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}

	return nil
}

func resourceSecretManagerSecretVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	err := setEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	return resourceSecretManagerSecretVersionRead(d, meta)
}

func resourceSecretManagerSecretVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:destroy")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	deletionPolicy := d.Get("deletion_policy")

	if deletionPolicy == "ABANDON" {
		return nil
	} else if deletionPolicy == "DISABLE" {
		url, err = tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:disable")
		if err != nil {
			return err
		}
	}

	log.Printf("[DEBUG] Deleting SecretVersion %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "SecretVersion")
	}

	log.Printf("[DEBUG] Finished deleting SecretVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceSecretManagerSecretVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("(projects/.+/secrets/.+)/versions/.+$")
	versionRegex := regexp.MustCompile("projects/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		return nil, fmt.Errorf("Version name does not fit the format `projects/{{project}}/secrets/{{secret}}/versions/{{version}}`")
	}
	if err := d.Set("secret", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting secret: %s", err)
	}

	parts = versionRegex.FindStringSubmatch(name)

	if err := d.Set("version", parts[3]); err != nil {
		return nil, fmt.Errorf("Error setting version: %s", err)
	}

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_policy", "DELETE"); err != nil {
		return nil, fmt.Errorf("Error setting version: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecretManagerSecretVersionEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v.(string) == "ENABLED" {
		return true
	}

	return false
}

func flattenSecretManagerSecretVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("projects/(.+)/secrets/(.+)/versions/(.+)$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 4 {
		panic(fmt.Sprintf("Version name does not fit the format `projects/{{project}}/secrets/{{secret}}/versions/{{version}}`"))
	}

	return parts[3]
}

func flattenSecretManagerSecretVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionDestroyTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionPayload(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	transformed := make(map[string]interface{})
	// write-only attributes are null on reads, secret_data_wo_version is used instead to return empty transformed that resolves a diff.
	if _, ok := d.GetOkExists("secret_data_wo_version"); ok {
		return []interface{}{transformed}
	}

	// if this secret version is disabled, the api will return an error, as the value cannot be accessed, return what we have
	if d.Get("enabled").(bool) == false {
		transformed["secret_data"] = d.Get("secret_data")
		return []interface{}{transformed}
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:access")
	if err != nil {
		return err
	}

	parts := strings.Split(d.Get("name").(string), "/")
	project := parts[1]

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	accessRes, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return err
	}

	if d.Get("is_secret_data_base64").(bool) {
		transformed["secret_data"] = accessRes["payload"].(map[string]interface{})["data"].(string)
	} else {
		data, err := base64.StdEncoding.DecodeString(accessRes["payload"].(map[string]interface{})["data"].(string))
		if err != nil {
			return err
		}
		transformed["secret_data"] = string(data)
	}
	return []interface{}{transformed}
}

func expandSecretManagerSecretVersionEnabled(_ interface{}, _ tpgresource.TerraformResourceData, _ *transport_tpg.Config) (interface{}, error) {
	return nil, nil
}

func expandSecretManagerSecretVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}
	transformedSecretDataWo, err := expandSecretManagerSecretVersionPayloadSecretDataWo(d.Get("secret_data_wo"), d.(*schema.ResourceData), config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretDataWo); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedSecretDataWo
	}
	return transformed, nil
}

func expandSecretManagerSecretVersionPayloadSecretData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	if d.Get("is_secret_data_base64").(bool) {
		return v, nil
	}
	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
func expandSecretManagerSecretVersionPayloadSecretDataWo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) (interface{}, error) {
	path := cty.GetAttrPath("secret_data_wo")
	woVal, _ := d.GetRawConfigAt(path)
	if !woVal.Type().Equals(cty.String) || woVal.IsNull() {
		return nil, nil
	}
	if d.Get("is_secret_data_base64").(bool) {
		return woVal.AsString(), nil
	}
	return base64.StdEncoding.EncodeToString([]byte(woVal.AsString())), nil
}

func resourceSecretManagerSecretVersionDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DESTROYED" {
		return nil, nil
	}

	return res, nil
}
func resourceSecretManagerSecretVersionPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	res, err := resourceSecretManagerSecretVersionDecoder(d, meta, res)
	if err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}
	if res == nil {
		return fmt.Errorf("decoding response, could not find object")
	}
	if err := d.Set("name", flattenSecretManagerSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
