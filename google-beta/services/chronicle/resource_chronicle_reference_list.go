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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/ReferenceList.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package chronicle

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

func ResourceChronicleReferenceList() *schema.Resource {
	return &schema.Resource{
		Create: resourceChronicleReferenceListCreate,
		Read:   resourceChronicleReferenceListRead,
		Update: resourceChronicleReferenceListUpdate,
		Delete: resourceChronicleReferenceListDelete,

		Importer: &schema.ResourceImporter{
			State: resourceChronicleReferenceListImport,
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
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Required. A user-provided description of the reference list.`,
			},
			"entries": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Required. The entries of the reference list.
When listed, they are returned in the order that was specified at creation
or update. The combined size of the values of the reference list may not
exceed 6MB.
This is returned only when the view is REFERENCE_LIST_VIEW_FULL.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Required. The value of the entry. Maximum length is 512 characters.`,
						},
					},
				},
			},
			"instance": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The unique identifier for the Chronicle instance, which is the same as the customer ID.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".`,
			},
			"reference_list_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Required. The ID to use for the reference list. This is also the display name for
the reference list. It must satisfy the following requirements:
- Starts with letter.
- Contains only letters, numbers and underscore.
- Has length < 256.
- Must be unique.`,
			},
			"syntax_type": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Possible values:
REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING
REFERENCE_LIST_SYNTAX_TYPE_REGEX
REFERENCE_LIST_SYNTAX_TYPE_CIDR`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The unique display name of the reference list.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The resource name of the reference list.
Format:
projects/{project}/locations/{location}/instances/{instance}/referenceLists/{reference_list}`,
			},
			"revision_create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The timestamp when the reference list was last updated.`,
			},
			"rule_associations_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Output only. The count of self-authored rules using the reference list.`,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Output only. The resource names for the associated self-authored Rules that use this
reference list.
This is returned only when the view is REFERENCE_LIST_VIEW_FULL.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scope_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `ScopeInfo specifies the scope info of the reference list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reference_list_scope": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `ReferenceListScope specifies the list of scope names of the reference list.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scope_names": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Optional. The list of scope names of the reference list. The scope names should be
full resource names and should be of the format:
"projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope_name}".`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
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

func resourceChronicleReferenceListCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandChronicleReferenceListDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	entriesProp, err := expandChronicleReferenceListEntries(d.Get("entries"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entries"); !tpgresource.IsEmptyValue(reflect.ValueOf(entriesProp)) && (ok || !reflect.DeepEqual(v, entriesProp)) {
		obj["entries"] = entriesProp
	}
	syntaxTypeProp, err := expandChronicleReferenceListSyntaxType(d.Get("syntax_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("syntax_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(syntaxTypeProp)) && (ok || !reflect.DeepEqual(v, syntaxTypeProp)) {
		obj["syntaxType"] = syntaxTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists?referenceListId={{reference_list_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ReferenceList: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReferenceList: %s", err)
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
		return fmt.Errorf("Error creating ReferenceList: %s", err)
	}
	if err := d.Set("name", flattenChronicleReferenceListName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists/{{reference_list_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ReferenceList %q: %#v", d.Id(), res)

	return resourceChronicleReferenceListRead(d, meta)
}

func resourceChronicleReferenceListRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists/{{reference_list_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReferenceList: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ChronicleReferenceList %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}

	if err := d.Set("name", flattenChronicleReferenceListName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("description", flattenChronicleReferenceListDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("entries", flattenChronicleReferenceListEntries(res["entries"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("scope_info", flattenChronicleReferenceListScopeInfo(res["scopeInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("display_name", flattenChronicleReferenceListDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("revision_create_time", flattenChronicleReferenceListRevisionCreateTime(res["revisionCreateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("rules", flattenChronicleReferenceListRules(res["rules"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("syntax_type", flattenChronicleReferenceListSyntaxType(res["syntaxType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}
	if err := d.Set("rule_associations_count", flattenChronicleReferenceListRuleAssociationsCount(res["ruleAssociationsCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReferenceList: %s", err)
	}

	return nil
}

func resourceChronicleReferenceListUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReferenceList: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandChronicleReferenceListDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	entriesProp, err := expandChronicleReferenceListEntries(d.Get("entries"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entries"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entriesProp)) {
		obj["entries"] = entriesProp
	}
	syntaxTypeProp, err := expandChronicleReferenceListSyntaxType(d.Get("syntax_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("syntax_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, syntaxTypeProp)) {
		obj["syntaxType"] = syntaxTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ChronicleBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists/{{reference_list_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ReferenceList %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("entries") {
		updateMask = append(updateMask, "entries")
	}

	if d.HasChange("syntax_type") {
		updateMask = append(updateMask, "syntaxType")
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
			return fmt.Errorf("Error updating ReferenceList %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ReferenceList %q: %#v", d.Id(), res)
		}

	}

	return resourceChronicleReferenceListRead(d, meta)
}

func resourceChronicleReferenceListDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Chronicle ReferenceList resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceChronicleReferenceListImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/instances/(?P<instance>[^/]+)/referenceLists/(?P<reference_list_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<instance>[^/]+)/(?P<reference_list_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<instance>[^/]+)/(?P<reference_list_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{instance}}/referenceLists/{{reference_list_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenChronicleReferenceListName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListEntries(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"value": flattenChronicleReferenceListEntriesValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenChronicleReferenceListEntriesValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListScopeInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["reference_list_scope"] =
		flattenChronicleReferenceListScopeInfoReferenceListScope(original["referenceListScope"], d, config)
	return []interface{}{transformed}
}
func flattenChronicleReferenceListScopeInfoReferenceListScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["scope_names"] =
		flattenChronicleReferenceListScopeInfoReferenceListScopeScopeNames(original["scopeNames"], d, config)
	return []interface{}{transformed}
}
func flattenChronicleReferenceListScopeInfoReferenceListScopeScopeNames(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListRevisionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListSyntaxType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenChronicleReferenceListRuleAssociationsCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandChronicleReferenceListDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleReferenceListEntries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedValue, err := expandChronicleReferenceListEntriesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandChronicleReferenceListEntriesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleReferenceListSyntaxType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
