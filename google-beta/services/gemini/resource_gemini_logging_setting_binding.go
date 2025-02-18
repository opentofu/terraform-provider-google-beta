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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gemini/LoggingSettingBinding.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gemini

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceGeminiLoggingSettingBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceGeminiLoggingSettingBindingCreate,
		Read:   resourceGeminiLoggingSettingBindingRead,
		Update: resourceGeminiLoggingSettingBindingUpdate,
		Delete: resourceGeminiLoggingSettingBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGeminiLoggingSettingBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Update: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(90 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"logging_setting_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"setting_binding_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Id of the setting binding.`,
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Target of the binding.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels as key value pairs.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"product": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GEMINI_CLOUD_ASSIST", "GEMINI_CODE_ASSIST", ""}),
				Description:  `Product type of the setting binding. Possible values: ["GEMINI_CLOUD_ASSIST", "GEMINI_CODE_ASSIST"]`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Create time stamp.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Identifier. Name of the resource.
Format:projects/{project}/locations/{location}/loggingSettings/{setting}/settingBindings/{setting_binding}`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Update time stamp.`,
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

func resourceGeminiLoggingSettingBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	targetProp, err := expandGeminiLoggingSettingBindingTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	productProp, err := expandGeminiLoggingSettingBindingProduct(d.Get("product"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("product"); !tpgresource.IsEmptyValue(reflect.ValueOf(productProp)) && (ok || !reflect.DeepEqual(v, productProp)) {
		obj["product"] = productProp
	}
	labelsProp, err := expandGeminiLoggingSettingBindingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings?settingBindingId={{setting_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new LoggingSettingBinding: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for LoggingSettingBinding: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating LoggingSettingBinding: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GeminiOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating LoggingSettingBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create LoggingSettingBinding: %s", err)
	}

	if err := d.Set("name", flattenGeminiLoggingSettingBindingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating LoggingSettingBinding %q: %#v", d.Id(), res)

	return resourceGeminiLoggingSettingBindingRead(d, meta)
}

func resourceGeminiLoggingSettingBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for LoggingSettingBinding: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GeminiLoggingSettingBinding %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}

	if err := d.Set("labels", flattenGeminiLoggingSettingBindingLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("target", flattenGeminiLoggingSettingBindingTarget(res["target"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("product", flattenGeminiLoggingSettingBindingProduct(res["product"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("name", flattenGeminiLoggingSettingBindingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("create_time", flattenGeminiLoggingSettingBindingCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("update_time", flattenGeminiLoggingSettingBindingUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("terraform_labels", flattenGeminiLoggingSettingBindingTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}
	if err := d.Set("effective_labels", flattenGeminiLoggingSettingBindingEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading LoggingSettingBinding: %s", err)
	}

	return nil
}

func resourceGeminiLoggingSettingBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for LoggingSettingBinding: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	targetProp, err := expandGeminiLoggingSettingBindingTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	productProp, err := expandGeminiLoggingSettingBindingProduct(d.Get("product"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("product"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, productProp)) {
		obj["product"] = productProp
	}
	labelsProp, err := expandGeminiLoggingSettingBindingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating LoggingSettingBinding %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("target") {
		updateMask = append(updateMask, "target")
	}

	if d.HasChange("product") {
		updateMask = append(updateMask, "product")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating LoggingSettingBinding %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating LoggingSettingBinding %q: %#v", d.Id(), res)
		}

		err = GeminiOperationWaitTime(
			config, res, project, "Updating LoggingSettingBinding", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceGeminiLoggingSettingBindingRead(d, meta)
}

func resourceGeminiLoggingSettingBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for LoggingSettingBinding: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting LoggingSettingBinding %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "LoggingSettingBinding")
	}

	err = GeminiOperationWaitTime(
		config, res, project, "Deleting LoggingSettingBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting LoggingSettingBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceGeminiLoggingSettingBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/loggingSettings/(?P<logging_setting_id>[^/]+)/settingBindings/(?P<setting_binding_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<logging_setting_id>[^/]+)/(?P<setting_binding_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<logging_setting_id>[^/]+)/(?P<setting_binding_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/loggingSettings/{{logging_setting_id}}/settingBindings/{{setting_binding_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGeminiLoggingSettingBindingLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGeminiLoggingSettingBindingTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiLoggingSettingBindingProduct(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiLoggingSettingBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiLoggingSettingBindingCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiLoggingSettingBindingUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiLoggingSettingBindingTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGeminiLoggingSettingBindingEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGeminiLoggingSettingBindingTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiLoggingSettingBindingProduct(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiLoggingSettingBindingEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
