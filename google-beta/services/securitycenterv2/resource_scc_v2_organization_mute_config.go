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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenterv2/OrganizationMuteConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenterv2

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceSecurityCenterV2OrganizationMuteConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityCenterV2OrganizationMuteConfigCreate,
		Read:   resourceSecurityCenterV2OrganizationMuteConfigRead,
		Update: resourceSecurityCenterV2OrganizationMuteConfigUpdate,
		Delete: resourceSecurityCenterV2OrganizationMuteConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecurityCenterV2OrganizationMuteConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
				Description: `An expression that defines the filter to apply across create/update
events of findings. While creating a filter string, be mindful of
the scope in which the mute configuration is being created. E.g.,
If a filter contains project = X but is created under the
project = Y scope, it might not match any findings.`,
			},
			"mute_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique identifier provided by the client within the parent scope.`,
			},
			"organization": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The organization whose Cloud Security Command Center the Mute
Config lives in.`,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The type of the mute config.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A description of the mute config.`,
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `location Id is provided by organization. If not provided, Use global as default.`,
				Default:     "global",
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The time at which the mute config was created. This field is set by
the server and will be ignored if provided on config creation.`,
			},
			"most_recent_editor": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Email address of the user who last edited the mute config. This
field is set by the server and will be ignored if provided on
config creation or update.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Name of the mute config. Its format is
organizations/{organization}/locations/global/muteConfigs/{configId},
folders/{folder}/locations/global/muteConfigs/{configId},
or projects/{project}/locations/global/muteConfigs/{configId}`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The most recent time at which the mute config was
updated. This field is set by the server and will be ignored if
provided on config creation or update.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecurityCenterV2OrganizationMuteConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2OrganizationMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterV2OrganizationMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterV2OrganizationMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/muteConfigs?muteConfigId={{mute_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationMuteConfig: %#v", obj)
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
		return fmt.Errorf("Error creating OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("name", flattenSecurityCenterV2OrganizationMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationMuteConfig %q: %#v", d.Id(), res)

	return resourceSecurityCenterV2OrganizationMuteConfigRead(d, meta)
}

func resourceSecurityCenterV2OrganizationMuteConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecurityCenterV2OrganizationMuteConfig %q", d.Id()))
	}

	if err := d.Set("name", flattenSecurityCenterV2OrganizationMuteConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("description", flattenSecurityCenterV2OrganizationMuteConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("filter", flattenSecurityCenterV2OrganizationMuteConfigFilter(res["filter"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("create_time", flattenSecurityCenterV2OrganizationMuteConfigCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("update_time", flattenSecurityCenterV2OrganizationMuteConfigUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("most_recent_editor", flattenSecurityCenterV2OrganizationMuteConfigMostRecentEditor(res["mostRecentEditor"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}
	if err := d.Set("type", flattenSecurityCenterV2OrganizationMuteConfigType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationMuteConfig: %s", err)
	}

	return nil
}

func resourceSecurityCenterV2OrganizationMuteConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterV2OrganizationMuteConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandSecurityCenterV2OrganizationMuteConfigFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	typeProp, err := expandSecurityCenterV2OrganizationMuteConfigType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationMuteConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("filter") {
		updateMask = append(updateMask, "filter")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
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
			return fmt.Errorf("Error updating OrganizationMuteConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating OrganizationMuteConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceSecurityCenterV2OrganizationMuteConfigRead(d, meta)
}

func resourceSecurityCenterV2OrganizationMuteConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{SecurityCenterV2BasePath}}organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting OrganizationMuteConfig %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "OrganizationMuteConfig")
	}

	log.Printf("[DEBUG] Finished deleting OrganizationMuteConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceSecurityCenterV2OrganizationMuteConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^organizations/(?P<organization>[^/]+)/locations/(?P<location>[^/]+)/muteConfigs/(?P<mute_config_id>[^/]+)$",
		"^(?P<organization>[^/]+)/(?P<location>[^/]+)/(?P<mute_config_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecurityCenterV2OrganizationMuteConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigFilter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigMostRecentEditor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecurityCenterV2OrganizationMuteConfigType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecurityCenterV2OrganizationMuteConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationMuteConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationMuteConfigType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
