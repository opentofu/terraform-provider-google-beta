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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/artifactregistry/VPCSCConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package artifactregistry

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
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceArtifactRegistryVPCSCConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceArtifactRegistryVPCSCConfigCreate,
		Read:   resourceArtifactRegistryVPCSCConfigRead,
		Update: resourceArtifactRegistryVPCSCConfigUpdate,
		Delete: resourceArtifactRegistryVPCSCConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceArtifactRegistryVPCSCConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the location this config is located in.`,
			},
			"vpcsc_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DENY", "ALLOW", ""}),
				Description:  `The VPC SC policy for project and location. Possible values: ["DENY", "ALLOW"]`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The name of the project's VPC SC Config.
Always of the form: projects/{project}/location/{location}/vpcscConfig`,
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

func resourceArtifactRegistryVPCSCConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	vpcscPolicyProp, err := expandArtifactRegistryVPCSCConfigVpcscPolicy(d.Get("vpcsc_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpcsc_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcscPolicyProp)) && (ok || !reflect.DeepEqual(v, vpcscPolicyProp)) {
		obj["vpcscPolicy"] = vpcscPolicyProp
	}

	obj, err = resourceArtifactRegistryVPCSCConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VPCSCConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VPCSCConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating VPCSCConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating VPCSCConfig %q: %#v", d.Id(), res)

	return resourceArtifactRegistryVPCSCConfigRead(d, meta)
}

func resourceArtifactRegistryVPCSCConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VPCSCConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ArtifactRegistryVPCSCConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VPCSCConfig: %s", err)
	}

	if err := d.Set("vpcsc_policy", flattenArtifactRegistryVPCSCConfigVpcscPolicy(res["vpcscPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading VPCSCConfig: %s", err)
	}
	if err := d.Set("name", flattenArtifactRegistryVPCSCConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading VPCSCConfig: %s", err)
	}

	return nil
}

func resourceArtifactRegistryVPCSCConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VPCSCConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	vpcscPolicyProp, err := expandArtifactRegistryVPCSCConfigVpcscPolicy(d.Get("vpcsc_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpcsc_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, vpcscPolicyProp)) {
		obj["vpcscPolicy"] = vpcscPolicyProp
	}

	obj, err = resourceArtifactRegistryVPCSCConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ArtifactRegistryBasePath}}projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating VPCSCConfig %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

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
		return fmt.Errorf("Error updating VPCSCConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating VPCSCConfig %q: %#v", d.Id(), res)
	}

	return resourceArtifactRegistryVPCSCConfigRead(d, meta)
}

func resourceArtifactRegistryVPCSCConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] ArtifactRegistry VPCSCConfig resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceArtifactRegistryVPCSCConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/vpcscConfig/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenArtifactRegistryVPCSCConfigVpcscPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenArtifactRegistryVPCSCConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandArtifactRegistryVPCSCConfigVpcscPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceArtifactRegistryVPCSCConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	if _, ok := d.GetOk("location"); !ok {
		location, err := tpgresource.GetRegion(d, config)
		if err != nil {
			return nil, fmt.Errorf("Cannot determine location: set in this resource, or set provider-level 'region' or 'zone'.")
		}
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	return obj, nil
}
