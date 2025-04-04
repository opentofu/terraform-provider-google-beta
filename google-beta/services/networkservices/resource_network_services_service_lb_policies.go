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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/ServiceLbPolicies.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkservices

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

func ResourceNetworkServicesServiceLbPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesServiceLbPoliciesCreate,
		Read:   resourceNetworkServicesServiceLbPoliciesRead,
		Update: resourceNetworkServicesServiceLbPoliciesUpdate,
		Delete: resourceNetworkServicesServiceLbPoliciesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesServiceLbPoliciesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The location of the service lb policy.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the ServiceLbPolicy resource. It matches pattern 'projects/{project}/locations/{location}/serviceLbPolicies/{service_lb_policy_name}'.`,
			},
			"auto_capacity_drain": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Option to specify if an unhealthy MIG/NEG should be considered for global load balancing and traffic routing.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Optional. If set to 'True', an unhealthy MIG/NEG will be set as drained. - An MIG/NEG is considered unhealthy if less than 25% of the instances/endpoints in the MIG/NEG are healthy. - This option will never result in draining more than 50% of the configured IGs/NEGs for the Backend Service.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"failover_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Option to specify health based failover behavior. This is not related to Network load balancer FailoverPolicy.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failover_health_threshold": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `Optional. The percentage threshold that a load balancer will begin to send traffic to failover backends. If the percentage of endpoints in a MIG/NEG is smaller than this value, traffic would be sent to failover backends if possible. This field should be set to a value between 1 and 99. The default value is 50 for Global external HTTP(S) load balancer (classic) and Proxyless service mesh, and 70 for others.`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the ServiceLbPolicy resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"load_balancing_algorithm": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"SPRAY_TO_REGION", "SPRAY_TO_WORLD", "WATERFALL_BY_REGION", "WATERFALL_BY_ZONE", ""}),
				Description:  `The type of load balancing algorithm to be used. The default behavior is WATERFALL_BY_REGION. Possible values: ["SPRAY_TO_REGION", "SPRAY_TO_WORLD", "WATERFALL_BY_REGION", "WATERFALL_BY_ZONE"]`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the ServiceLbPolicy was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: `Time the ServiceLbPolicy was updated in UTC.`,
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

func resourceNetworkServicesServiceLbPoliciesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesServiceLbPoliciesDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	loadBalancingAlgorithmProp, err := expandNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(d.Get("load_balancing_algorithm"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancingAlgorithmProp)) && (ok || !reflect.DeepEqual(v, loadBalancingAlgorithmProp)) {
		obj["loadBalancingAlgorithm"] = loadBalancingAlgorithmProp
	}
	autoCapacityDrainProp, err := expandNetworkServicesServiceLbPoliciesAutoCapacityDrain(d.Get("auto_capacity_drain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_capacity_drain"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoCapacityDrainProp)) && (ok || !reflect.DeepEqual(v, autoCapacityDrainProp)) {
		obj["autoCapacityDrain"] = autoCapacityDrainProp
	}
	failoverConfigProp, err := expandNetworkServicesServiceLbPoliciesFailoverConfig(d.Get("failover_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(failoverConfigProp)) && (ok || !reflect.DeepEqual(v, failoverConfigProp)) {
		obj["failoverConfig"] = failoverConfigProp
	}
	labelsProp, err := expandNetworkServicesServiceLbPoliciesEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/serviceLbPolicies?serviceLbPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceLbPolicies: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceLbPolicies: %s", err)
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
		return fmt.Errorf("Error creating ServiceLbPolicies: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating ServiceLbPolicies", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceLbPolicies: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceLbPolicies %q: %#v", d.Id(), res)

	return resourceNetworkServicesServiceLbPoliciesRead(d, meta)
}

func resourceNetworkServicesServiceLbPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceLbPolicies: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkServicesServiceLbPolicies %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkServicesServiceLbPoliciesCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesServiceLbPoliciesUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesServiceLbPoliciesLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesServiceLbPoliciesDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("load_balancing_algorithm", flattenNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(res["loadBalancingAlgorithm"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("auto_capacity_drain", flattenNetworkServicesServiceLbPoliciesAutoCapacityDrain(res["autoCapacityDrain"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("failover_config", flattenNetworkServicesServiceLbPoliciesFailoverConfig(res["failoverConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkServicesServiceLbPoliciesTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkServicesServiceLbPoliciesEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceLbPolicies: %s", err)
	}

	return nil
}

func resourceNetworkServicesServiceLbPoliciesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceLbPolicies: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesServiceLbPoliciesDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	loadBalancingAlgorithmProp, err := expandNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(d.Get("load_balancing_algorithm"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, loadBalancingAlgorithmProp)) {
		obj["loadBalancingAlgorithm"] = loadBalancingAlgorithmProp
	}
	autoCapacityDrainProp, err := expandNetworkServicesServiceLbPoliciesAutoCapacityDrain(d.Get("auto_capacity_drain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("auto_capacity_drain"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoCapacityDrainProp)) {
		obj["autoCapacityDrain"] = autoCapacityDrainProp
	}
	failoverConfigProp, err := expandNetworkServicesServiceLbPoliciesFailoverConfig(d.Get("failover_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, failoverConfigProp)) {
		obj["failoverConfig"] = failoverConfigProp
	}
	labelsProp, err := expandNetworkServicesServiceLbPoliciesEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceLbPolicies %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("load_balancing_algorithm") {
		updateMask = append(updateMask, "loadBalancingAlgorithm")
	}

	if d.HasChange("auto_capacity_drain") {
		updateMask = append(updateMask, "autoCapacityDrain")
	}

	if d.HasChange("failover_config") {
		updateMask = append(updateMask, "failoverConfig")
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
			return fmt.Errorf("Error updating ServiceLbPolicies %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ServiceLbPolicies %q: %#v", d.Id(), res)
		}

		err = NetworkServicesOperationWaitTime(
			config, res, project, "Updating ServiceLbPolicies", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkServicesServiceLbPoliciesRead(d, meta)
}

func resourceNetworkServicesServiceLbPoliciesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceLbPolicies: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting ServiceLbPolicies %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ServiceLbPolicies")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting ServiceLbPolicies", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServiceLbPolicies %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesServiceLbPoliciesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/serviceLbPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceLbPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesServiceLbPoliciesCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceLbPoliciesUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceLbPoliciesLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesServiceLbPoliciesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceLbPoliciesAutoCapacityDrain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable"] =
		flattenNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(original["enable"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesServiceLbPoliciesFailoverConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["failover_health_threshold"] =
		flattenNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(original["failoverHealthThreshold"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesServiceLbPoliciesTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesServiceLbPoliciesEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkServicesServiceLbPoliciesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesLoadBalancingAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesAutoCapacityDrain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnable, err := expandNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandNetworkServicesServiceLbPoliciesAutoCapacityDrainEnable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesFailoverConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFailoverHealthThreshold, err := expandNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(original["failover_health_threshold"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFailoverHealthThreshold); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["failoverHealthThreshold"] = transformedFailoverHealthThreshold
	}

	return transformed, nil
}

func expandNetworkServicesServiceLbPoliciesFailoverConfigFailoverHealthThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesServiceLbPoliciesEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
