// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package backupdr

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceBackupDRManagementServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupDRManagementServerCreate,
		Read:   resourceBackupDRManagementServerRead,
		Delete: resourceBackupDRManagementServerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBackupDRManagementServerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(40 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the management server (management console)`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of management server (management console)`,
			},
			"networks": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Network details to create management server (management console).`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Network with format 'projects/{{project_id}}/global/networks/{{network_id}}'`,
						},
						"peering_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"PRIVATE_SERVICE_ACCESS", ""}),
							Description:  `Type of Network peeringMode Default value: "PRIVATE_SERVICE_ACCESS" Possible values: ["PRIVATE_SERVICE_ACCESS"]`,
							Default:      "PRIVATE_SERVICE_ACCESS",
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"BACKUP_RESTORE", ""}),
				Description:  `The type of management server (management console). Default value: "BACKUP_RESTORE" Possible values: ["BACKUP_RESTORE"]`,
				Default:      "BACKUP_RESTORE",
			},
			"management_uri": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The management console URI`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The management console api endpoint.`,
						},
						"web_ui": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The management console webUi.`,
						},
					},
				},
			},
			"oauth2_client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The oauth2ClientId of management console.`,
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

func resourceBackupDRManagementServerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	typeProp, err := expandBackupDRManagementServerType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	networksProp, err := expandBackupDRManagementServerNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); !tpgresource.IsEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/managementServers/?management_server_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagementServer: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ManagementServer: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ManagementServer: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/managementServers/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = BackupDROperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ManagementServer", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ManagementServer: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/managementServers/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ManagementServer %q: %#v", d.Id(), res)

	return resourceBackupDRManagementServerRead(d, meta)
}

func resourceBackupDRManagementServerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/managementServers/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ManagementServer: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BackupDRManagementServer %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ManagementServer: %s", err)
	}

	if err := d.Set("type", flattenBackupDRManagementServerType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagementServer: %s", err)
	}
	if err := d.Set("networks", flattenBackupDRManagementServerNetworks(res["networks"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagementServer: %s", err)
	}
	if err := d.Set("oauth2_client_id", flattenBackupDRManagementServerOauth2ClientId(res["oauth2ClientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagementServer: %s", err)
	}
	if err := d.Set("management_uri", flattenBackupDRManagementServerManagementUri(res["managementUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading ManagementServer: %s", err)
	}

	return nil
}

func resourceBackupDRManagementServerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ManagementServer: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BackupDRBasePath}}projects/{{project}}/locations/{{location}}/managementServers/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ManagementServer %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ManagementServer")
	}

	err = BackupDROperationWaitTime(
		config, res, project, "Deleting ManagementServer", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ManagementServer %q: %#v", d.Id(), res)
	return nil
}

func resourceBackupDRManagementServerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/managementServers/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/managementServers/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBackupDRManagementServerType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRManagementServerNetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"network":      flattenBackupDRManagementServerNetworksNetwork(original["network"], d, config),
			"peering_mode": flattenBackupDRManagementServerNetworksPeeringMode(original["peeringMode"], d, config),
		})
	}
	return transformed
}
func flattenBackupDRManagementServerNetworksNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRManagementServerNetworksPeeringMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRManagementServerOauth2ClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRManagementServerManagementUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["web_ui"] =
		flattenBackupDRManagementServerManagementUriWebUi(original["webUi"], d, config)
	transformed["api"] =
		flattenBackupDRManagementServerManagementUriApi(original["api"], d, config)
	return []interface{}{transformed}
}
func flattenBackupDRManagementServerManagementUriWebUi(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBackupDRManagementServerManagementUriApi(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandBackupDRManagementServerType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRManagementServerNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandBackupDRManagementServerNetworksNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		transformedPeeringMode, err := expandBackupDRManagementServerNetworksPeeringMode(original["peering_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPeeringMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["peeringMode"] = transformedPeeringMode
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBackupDRManagementServerNetworksNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRManagementServerNetworksPeeringMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
