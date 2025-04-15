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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryreservation/ReservationAssignment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigqueryreservation

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceBigqueryReservationReservationAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationReservationAssignmentCreate,
		Read:   resourceBigqueryReservationReservationAssignmentRead,
		Delete: resourceBigqueryReservationReservationAssignmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationReservationAssignmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"assignee": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.`,
			},
			"job_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY, CONTINUOUS`,
			},
			"reservation": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The reservation for the resource`,
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The resource name of the assignment.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active.
Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE`,
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

func resourceBigqueryReservationReservationAssignmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	assigneeProp, err := expandNestedBigqueryReservationReservationAssignmentAssignee(d.Get("assignee"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("assignee"); !tpgresource.IsEmptyValue(reflect.ValueOf(assigneeProp)) && (ok || !reflect.DeepEqual(v, assigneeProp)) {
		obj["assignee"] = assigneeProp
	}
	jobTypeProp, err := expandNestedBigqueryReservationReservationAssignmentJobType(d.Get("job_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("job_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(jobTypeProp)) && (ok || !reflect.DeepEqual(v, jobTypeProp)) {
		obj["jobType"] = jobTypeProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ReservationAssignment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if _, ok := d.GetOkExists("location"); !ok {
		// Extract location from parent reservation.
		reservation := d.Get("reservation").(string)

		tableRef := regexp.MustCompile("projects/(.+)/locations/(.+)/reservations/(.+)")
		if parts := tableRef.FindStringSubmatch(reservation); parts != nil {
			err := d.Set("location", parts[2])
			if err != nil {
				return err
			}
		}

		if strings.Contains(url, "locations//") {
			// re-compute url now that location must be set
			url = strings.ReplaceAll(url, "/locations//", "/locations/"+d.Get("location").(string)+"/")
			if err != nil {
				return err
			}
		}
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
		return fmt.Errorf("Error creating ReservationAssignment: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceBigqueryReservationReservationAssignmentPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ReservationAssignment %q: %#v", d.Id(), res)

	return resourceBigqueryReservationReservationAssignmentRead(d, meta)
}

func resourceBigqueryReservationReservationAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationReservationAssignment %q", d.Id()))
	}

	res, err = flattenNestedBigqueryReservationReservationAssignment(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing BigqueryReservationReservationAssignment because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}

	if err := d.Set("name", flattenNestedBigqueryReservationReservationAssignmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}
	if err := d.Set("assignee", flattenNestedBigqueryReservationReservationAssignmentAssignee(res["assignee"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}
	if err := d.Set("job_type", flattenNestedBigqueryReservationReservationAssignmentJobType(res["jobType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}
	if err := d.Set("state", flattenNestedBigqueryReservationReservationAssignmentState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}

	return nil
}

func resourceBigqueryReservationReservationAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ReservationAssignment %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ReservationAssignment")
	}

	log.Printf("[DEBUG] Finished deleting ReservationAssignment %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationReservationAssignmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/reservations/(?P<reservation>[^/]+)/assignments/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<reservation>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<reservation>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedBigqueryReservationReservationAssignmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenNestedBigqueryReservationReservationAssignmentAssignee(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigqueryReservationReservationAssignmentJobType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedBigqueryReservationReservationAssignmentState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedBigqueryReservationReservationAssignmentAssignee(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedBigqueryReservationReservationAssignmentJobType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenNestedBigqueryReservationReservationAssignment(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["assignments"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value assignments. Actual value: %v", v)
	}

	_, item, err := resourceBigqueryReservationReservationAssignmentFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceBigqueryReservationReservationAssignmentFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName := d.Get("name")
	expectedFlattenedName := flattenNestedBigqueryReservationReservationAssignmentName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenNestedBigqueryReservationReservationAssignmentName(item["name"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemName)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceBigqueryReservationReservationAssignmentPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("name", flattenNestedBigqueryReservationReservationAssignmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
