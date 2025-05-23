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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/RegionCommitment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

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

func ResourceComputeRegionCommitment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionCommitmentCreate,
		Read:   resourceComputeRegionCommitmentRead,
		Delete: resourceComputeRegionCommitmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionCommitmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. The name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"plan": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TWELVE_MONTH", "THIRTY_SIX_MONTH"}),
				Description: `The plan for this commitment, which determines duration and discount rate.
The currently supported plans are TWELVE_MONTH (1 year), and THIRTY_SIX_MONTH (3 years). Possible values: ["TWELVE_MONTH", "THIRTY_SIX_MONTH"]`,
			},
			"auto_renew": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Specifies whether to enable automatic renewal for the commitment.
The default value is false if not specified.
If the field is set to true, the commitment will be automatically renewed for either
one or three years according to the terms of the existing commitment.`,
			},
			"category": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"LICENSE", "MACHINE", ""}),
				Description: `The category of the commitment. Category MACHINE specifies commitments composed of
machine resources such as VCPU or MEMORY, listed in resources. Category LICENSE
specifies commitments composed of software licenses, listed in licenseResources.
Note that only MACHINE commitments should have a Type specified. Possible values: ["LICENSE", "MACHINE"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"existing_reservations": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the already existing reservations to attach to the Commitment.`,
			},
			"license_resource": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The license specification required as part of a license commitment.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `Any applicable license URI.`,
						},
						"amount": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The number of licenses purchased.`,
						},
						"cores_per_license": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Specifies the core range of the instance for which this license applies.`,
						},
					},
				},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `URL of the region where this commitment may be used.`,
			},
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `A list of commitment amounts for particular resources.
Note that VCPU and MEMORY resource commitments must occur together.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accelerator_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Name of the accelerator type resource. Applicable only when the type is ACCELERATOR.`,
						},
						"amount": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The amount of the resource purchased (in a type-dependent unit,
such as bytes). For vCPUs, this can just be an integer. For memory,
this must be provided in MB. Memory must be a multiple of 256 MB,
with up to 6.5GB of memory per every vCPU.`,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Type of resource for which this commitment applies.
Possible values are VCPU, MEMORY, LOCAL_SSD, and ACCELERATOR.`,
						},
					},
				},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The type of commitment, which affects the discount rate and the eligible resources.
The type could be one of the following value: 'MEMORY_OPTIMIZED', 'ACCELERATOR_OPTIMIZED',
'GENERAL_PURPOSE_N1', 'GENERAL_PURPOSE_N2', 'GENERAL_PURPOSE_N2D', 'GENERAL_PURPOSE_E2',
'GENERAL_PURPOSE_T2D', 'GENERAL_PURPOSE_C3', 'COMPUTE_OPTIMIZED_C2', 'COMPUTE_OPTIMIZED_C2D' and
'GRAPHICS_OPTIMIZED_G2'`,
			},
			"commitment_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Unique identifier for the resource.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"end_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Commitment end time in RFC3339 text format.`,
			},
			"start_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Commitment start time in RFC3339 text format.`,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Status of the commitment with regards to eventual expiration
(each commitment has an end date defined).`,
			},
			"status_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A human-readable explanation of the status.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRegionCommitmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionCommitmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionCommitmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	planProp, err := expandComputeRegionCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("plan"); !tpgresource.IsEmptyValue(reflect.ValueOf(planProp)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	resourcesProp, err := expandComputeRegionCommitmentResources(d.Get("resources"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourcesProp)) && (ok || !reflect.DeepEqual(v, resourcesProp)) {
		obj["resources"] = resourcesProp
	}
	typeProp, err := expandComputeRegionCommitmentType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	categoryProp, err := expandComputeRegionCommitmentCategory(d.Get("category"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("category"); !tpgresource.IsEmptyValue(reflect.ValueOf(categoryProp)) && (ok || !reflect.DeepEqual(v, categoryProp)) {
		obj["category"] = categoryProp
	}
	licenseResourceProp, err := expandComputeRegionCommitmentLicenseResource(d.Get("license_resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("license_resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(licenseResourceProp)) && (ok || !reflect.DeepEqual(v, licenseResourceProp)) {
		obj["licenseResource"] = licenseResourceProp
	}
	autoRenewProp, err := expandComputeRegionCommitmentAutoRenew(d.Get("auto_renew"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_renew"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoRenewProp)) && (ok || !reflect.DeepEqual(v, autoRenewProp)) {
		obj["autoRenew"] = autoRenewProp
	}
	existingReservationsProp, err := expandComputeRegionCommitmentExistingReservations(d.Get("existing_reservations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("existing_reservations"); !tpgresource.IsEmptyValue(reflect.ValueOf(existingReservationsProp)) && (ok || !reflect.DeepEqual(v, existingReservationsProp)) {
		obj["existingReservations"] = existingReservationsProp
	}
	regionProp, err := expandComputeRegionCommitmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/commitments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionCommitment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionCommitment: %s", err)
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
		return fmt.Errorf("Error creating RegionCommitment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/commitments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating RegionCommitment", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionCommitment: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionCommitment %q: %#v", d.Id(), res)

	return resourceComputeRegionCommitmentRead(d, meta)
}

func resourceComputeRegionCommitmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/commitments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionCommitment: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeRegionCommitment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}

	if err := d.Set("commitment_id", flattenComputeRegionCommitmentCommitmentId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeRegionCommitmentCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionCommitmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionCommitmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("status", flattenComputeRegionCommitmentStatus(res["status"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("status_message", flattenComputeRegionCommitmentStatusMessage(res["statusMessage"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("plan", flattenComputeRegionCommitmentPlan(res["plan"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("start_timestamp", flattenComputeRegionCommitmentStartTimestamp(res["startTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("end_timestamp", flattenComputeRegionCommitmentEndTimestamp(res["endTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("resources", flattenComputeRegionCommitmentResources(res["resources"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("type", flattenComputeRegionCommitmentType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("category", flattenComputeRegionCommitmentCategory(res["category"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("license_resource", flattenComputeRegionCommitmentLicenseResource(res["licenseResource"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("auto_renew", flattenComputeRegionCommitmentAutoRenew(res["autoRenew"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("existing_reservations", flattenComputeRegionCommitmentExistingReservations(res["existingReservations"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionCommitmentRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionCommitment: %s", err)
	}

	return nil
}

func resourceComputeRegionCommitmentDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Compute RegionCommitment resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceComputeRegionCommitmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/commitments/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/commitments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionCommitmentCommitmentId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeRegionCommitmentCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentStatusMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentPlan(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentStartTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentEndTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"type":             flattenComputeRegionCommitmentResourcesType(original["type"], d, config),
			"amount":           flattenComputeRegionCommitmentResourcesAmount(original["amount"], d, config),
			"accelerator_type": flattenComputeRegionCommitmentResourcesAcceleratorType(original["acceleratorType"], d, config),
		})
	}
	return transformed
}
func flattenComputeRegionCommitmentResourcesType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentResourcesAmount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentResourcesAcceleratorType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentCategory(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentLicenseResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["license"] =
		flattenComputeRegionCommitmentLicenseResourceLicense(original["license"], d, config)
	transformed["amount"] =
		flattenComputeRegionCommitmentLicenseResourceAmount(original["amount"], d, config)
	transformed["cores_per_license"] =
		flattenComputeRegionCommitmentLicenseResourceCoresPerLicense(original["coresPerLicense"], d, config)
	return []interface{}{transformed}
}
func flattenComputeRegionCommitmentLicenseResourceLicense(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentLicenseResourceAmount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentLicenseResourceCoresPerLicense(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentAutoRenew(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentExistingReservations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeRegionCommitmentRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.GetResourceNameFromSelfLink(v.(string))
}

func expandComputeRegionCommitmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentPlan(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandComputeRegionCommitmentResourcesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedAmount, err := expandComputeRegionCommitmentResourcesAmount(original["amount"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAmount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["amount"] = transformedAmount
		}

		transformedAcceleratorType, err := expandComputeRegionCommitmentResourcesAcceleratorType(original["accelerator_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["acceleratorType"] = transformedAcceleratorType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionCommitmentResourcesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResourcesAmount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResourcesAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentCategory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLicense, err := expandComputeRegionCommitmentLicenseResourceLicense(original["license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["license"] = transformedLicense
	}

	transformedAmount, err := expandComputeRegionCommitmentLicenseResourceAmount(original["amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAmount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["amount"] = transformedAmount
	}

	transformedCoresPerLicense, err := expandComputeRegionCommitmentLicenseResourceCoresPerLicense(original["cores_per_license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCoresPerLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["coresPerLicense"] = transformedCoresPerLicense
	}

	return transformed, nil
}

func expandComputeRegionCommitmentLicenseResourceLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResourceAmount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResourceCoresPerLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentAutoRenew(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentExistingReservations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
