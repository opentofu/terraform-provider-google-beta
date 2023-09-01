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

package corebilling

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceCoreBillingProjectInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceCoreBillingProjectInfoCreate,
		Read:   resourceCoreBillingProjectInfoRead,
		Update: resourceCoreBillingProjectInfoUpdate,
		Delete: resourceCoreBillingProjectInfoDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCoreBillingProjectInfoImport,
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
			"billing_account": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The ID of the billing account associated with the project, if
any. Set to empty string to disable billing for the project.
For example, '"012345-567890-ABCDEF"' or '""'.`,
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

func resourceCoreBillingProjectInfoCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	billing_accountProp, err := expandCoreBillingProjectInfoBillingAccount(d.Get("billing_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(billing_accountProp)) && (ok || !reflect.DeepEqual(v, billing_accountProp)) {
		obj["billing_account"] = billing_accountProp
	}

	obj, err = resourceCoreBillingProjectInfoEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CoreBillingBasePath}}projects/{{project}}/billingInfo/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectInfo: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectInfo: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ProjectInfo: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/billingInfo")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ProjectInfo %q: %#v", d.Id(), res)

	return resourceCoreBillingProjectInfoRead(d, meta)
}

func resourceCoreBillingProjectInfoRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CoreBillingBasePath}}projects/{{project}}/billingInfo/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectInfo: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("CoreBillingProjectInfo %q", d.Id()))
	}

	res, err = resourceCoreBillingProjectInfoDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing CoreBillingProjectInfo because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectInfo: %s", err)
	}

	if err := d.Set("billing_account", flattenCoreBillingProjectInfoBillingAccount(res["billing_account"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectInfo: %s", err)
	}

	return nil
}

func resourceCoreBillingProjectInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectInfo: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	billing_accountProp, err := expandCoreBillingProjectInfoBillingAccount(d.Get("billing_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, billing_accountProp)) {
		obj["billing_account"] = billing_accountProp
	}

	obj, err = resourceCoreBillingProjectInfoEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{CoreBillingBasePath}}projects/{{project}}/billingInfo/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ProjectInfo %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ProjectInfo %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ProjectInfo %q: %#v", d.Id(), res)
	}

	return resourceCoreBillingProjectInfoRead(d, meta)
}

func resourceCoreBillingProjectInfoDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectInfo: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{CoreBillingBasePath}}projects/{{project}}/billingInfo/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ProjectInfo %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ProjectInfo")
	}

	log.Printf("[DEBUG] Finished deleting ProjectInfo %q: %#v", d.Id(), res)
	return nil
}

func resourceCoreBillingProjectInfoImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/billingInfo")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCoreBillingProjectInfoBillingAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandCoreBillingProjectInfoBillingAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceCoreBillingProjectInfoEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	ba := d.Get("billing_account").(string)
	if ba == "" {
		obj["billingAccountName"] = ""
	} else {
		obj["billingAccountName"] = "billingAccounts/" + ba
	}
	delete(obj, "billing_account")
	return obj, nil
}

func resourceCoreBillingProjectInfoDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	res["billing_account"] = strings.TrimPrefix(res["billingAccountName"].(string), "billingAccounts/")
	return res, nil
}
