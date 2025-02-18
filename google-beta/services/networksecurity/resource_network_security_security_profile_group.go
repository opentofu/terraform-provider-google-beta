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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/SecurityProfileGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

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
)

func ResourceNetworkSecuritySecurityProfileGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecuritySecurityProfileGroupCreate,
		Read:   resourceNetworkSecuritySecurityProfileGroupRead,
		Update: resourceNetworkSecuritySecurityProfileGroupUpdate,
		Delete: resourceNetworkSecuritySecurityProfileGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecuritySecurityProfileGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the security profile group resource.`,
			},
			"custom_intercept_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Reference to a SecurityProfile with the CustomIntercept configuration.`,
			},
			"custom_mirroring_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of the profile. The Max length is 512 characters.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `A map of key/value label pairs to assign to the resource.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The location of the security profile group.
The default value is 'global'.`,
				Default: "global",
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The name of the parent this security profile group belongs to.
Format: organizations/{organization_id}.`,
			},
			"threat_prevention_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the security profile group was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `This checksum is computed by the server based on the value of other fields,
and may be sent on update and delete requests to ensure the client has an up-to-date
value before proceeding.`,
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
				Description: `Time the security profile group was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecuritySecurityProfileGroupCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileGroupThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(threatPreventionProfileProp)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	customMirroringProfileProp, err := expandNetworkSecuritySecurityProfileGroupCustomMirroringProfile(d.Get("custom_mirroring_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_mirroring_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(customMirroringProfileProp)) && (ok || !reflect.DeepEqual(v, customMirroringProfileProp)) {
		obj["customMirroringProfile"] = customMirroringProfileProp
	}
	customInterceptProfileProp, err := expandNetworkSecuritySecurityProfileGroupCustomInterceptProfile(d.Get("custom_intercept_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_intercept_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(customInterceptProfileProp)) && (ok || !reflect.DeepEqual(v, customInterceptProfileProp)) {
		obj["customInterceptProfile"] = customInterceptProfileProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups?securityProfileGroupId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecurityProfileGroup: %#v", obj)
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
		return fmt.Errorf("Error creating SecurityProfileGroup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating SecurityProfileGroup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SecurityProfileGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SecurityProfileGroup %q: %#v", d.Id(), res)

	return resourceNetworkSecuritySecurityProfileGroupRead(d, meta)
}

func resourceNetworkSecuritySecurityProfileGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecuritySecurityProfileGroup %q", d.Id()))
	}

	if err := d.Set("create_time", flattenNetworkSecuritySecurityProfileGroupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecuritySecurityProfileGroupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("etag", flattenNetworkSecuritySecurityProfileGroupEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecuritySecurityProfileGroupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecuritySecurityProfileGroupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("threat_prevention_profile", flattenNetworkSecuritySecurityProfileGroupThreatPreventionProfile(res["threatPreventionProfile"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("custom_mirroring_profile", flattenNetworkSecuritySecurityProfileGroupCustomMirroringProfile(res["customMirroringProfile"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("custom_intercept_profile", flattenNetworkSecuritySecurityProfileGroupCustomInterceptProfile(res["customInterceptProfile"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecuritySecurityProfileGroupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecuritySecurityProfileGroupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfileGroup: %s", err)
	}

	return nil
}

func resourceNetworkSecuritySecurityProfileGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileGroupThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	customMirroringProfileProp, err := expandNetworkSecuritySecurityProfileGroupCustomMirroringProfile(d.Get("custom_mirroring_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_mirroring_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customMirroringProfileProp)) {
		obj["customMirroringProfile"] = customMirroringProfileProp
	}
	customInterceptProfileProp, err := expandNetworkSecuritySecurityProfileGroupCustomInterceptProfile(d.Get("custom_intercept_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_intercept_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customInterceptProfileProp)) {
		obj["customInterceptProfile"] = customInterceptProfileProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SecurityProfileGroup %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("threat_prevention_profile") {
		updateMask = append(updateMask, "threatPreventionProfile")
	}

	if d.HasChange("custom_mirroring_profile") {
		updateMask = append(updateMask, "customMirroringProfile")
	}

	if d.HasChange("custom_intercept_profile") {
		updateMask = append(updateMask, "customInterceptProfile")
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
			return fmt.Errorf("Error updating SecurityProfileGroup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SecurityProfileGroup %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating SecurityProfileGroup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecuritySecurityProfileGroupRead(d, meta)
}

func resourceNetworkSecuritySecurityProfileGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SecurityProfileGroup %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "SecurityProfileGroup")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting SecurityProfileGroup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SecurityProfileGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecuritySecurityProfileGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/locations/(?P<location>[^/]+)/securityProfileGroups/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecuritySecurityProfileGroupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecuritySecurityProfileGroupThreatPreventionProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupCustomMirroringProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupCustomInterceptProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileGroupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecuritySecurityProfileGroupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecuritySecurityProfileGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupThreatPreventionProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupCustomMirroringProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupCustomInterceptProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
