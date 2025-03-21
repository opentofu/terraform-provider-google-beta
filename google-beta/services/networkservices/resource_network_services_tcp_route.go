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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/TcpRoute.yaml
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

func ResourceNetworkServicesTcpRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesTcpRouteCreate,
		Read:   resourceNetworkServicesTcpRouteRead,
		Update: resourceNetworkServicesTcpRouteUpdate,
		Delete: resourceNetworkServicesTcpRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesTcpRouteImport,
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the TcpRoute resource.`,
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.
If there are multiple rules then the action taken will be the first rule to match.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `A detailed rule defining how to route traffic.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destinations": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The destination services to which traffic should be forwarded. At least one destination service is required.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"service_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: `The URL of a BackendService to route traffic to.`,
												},
												"weight": {
													Type:     schema.TypeInt,
													Optional: true,
													Description: `Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports.
If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend.
If weights are specified for any one service name, they need to be specified for all of them.
If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.`,
												},
											},
										},
									},
									"idle_timeout": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateRegexp(`^(0|[1-9][0-9]*)(\.[0-9]{1,9})?s$`),
										Description: `Specifies the idle timeout for the selected route. The idle timeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection. If not set, the default idle timeout is 30 seconds. If set to 0s, the timeout will be disabled.

A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".`,
									},
									"original_destination": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: `If true, Router will use the destination IP and port of the original connection as the destination of the request.`,
									},
								},
							},
						},
						"matches": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `RouteMatch defines the predicate used to match requests to a given action. Multiple match types are "OR"ed for evaluation.
If no routeMatch field is specified, this rule will unconditionally match traffic.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:     schema.TypeString,
										Required: true,
										Description: `Must be specified in the CIDR range format. A CIDR range consists of an IP Address and a prefix length to construct the subnet mask.
By default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: "10.0.0.1" - matches against this exact IP address. "10.0.0.0/8" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. "0.0.0.0/0" - matches against any IP address'.`,
									},
									"port": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Specifies the destination port to match against.`,
									},
								},
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"gateways": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
Each gateway reference should match the pattern: projects/*/locations/global/gateways/<gateway_name>`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the TcpRoute resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"meshes": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>
The attached Mesh should be of a type SIDECAR`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the TcpRoute was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
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
				Description: `Time the TcpRoute was updated in UTC.`,
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

func resourceNetworkServicesTcpRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesTcpRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	meshesProp, err := expandNetworkServicesTcpRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("meshes"); ok || !reflect.DeepEqual(v, meshesProp) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesTcpRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateways"); ok || !reflect.DeepEqual(v, gatewaysProp) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesTcpRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); ok || !reflect.DeepEqual(v, rulesProp) {
		obj["rules"] = rulesProp
	}
	labelsProp, err := expandNetworkServicesTcpRouteEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tcpRoutes?tcpRouteId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TcpRoute: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TcpRoute: %s", err)
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
		return fmt.Errorf("Error creating TcpRoute: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating TcpRoute", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TcpRoute: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TcpRoute %q: %#v", d.Id(), res)

	return resourceNetworkServicesTcpRouteRead(d, meta)
}

func resourceNetworkServicesTcpRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TcpRoute: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkServicesTcpRoute %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}

	if err := d.Set("self_link", flattenNetworkServicesTcpRouteSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkServicesTcpRouteCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesTcpRouteUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesTcpRouteLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesTcpRouteDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("meshes", flattenNetworkServicesTcpRouteMeshes(res["meshes"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("gateways", flattenNetworkServicesTcpRouteGateways(res["gateways"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("rules", flattenNetworkServicesTcpRouteRules(res["rules"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkServicesTcpRouteTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkServicesTcpRouteEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading TcpRoute: %s", err)
	}

	return nil
}

func resourceNetworkServicesTcpRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TcpRoute: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesTcpRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	meshesProp, err := expandNetworkServicesTcpRouteMeshes(d.Get("meshes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("meshes"); ok || !reflect.DeepEqual(v, meshesProp) {
		obj["meshes"] = meshesProp
	}
	gatewaysProp, err := expandNetworkServicesTcpRouteGateways(d.Get("gateways"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateways"); ok || !reflect.DeepEqual(v, gatewaysProp) {
		obj["gateways"] = gatewaysProp
	}
	rulesProp, err := expandNetworkServicesTcpRouteRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); ok || !reflect.DeepEqual(v, rulesProp) {
		obj["rules"] = rulesProp
	}
	labelsProp, err := expandNetworkServicesTcpRouteEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TcpRoute %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("meshes") {
		updateMask = append(updateMask, "meshes")
	}

	if d.HasChange("gateways") {
		updateMask = append(updateMask, "gateways")
	}

	if d.HasChange("rules") {
		updateMask = append(updateMask, "rules")
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
			return fmt.Errorf("Error updating TcpRoute %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TcpRoute %q: %#v", d.Id(), res)
		}

		err = NetworkServicesOperationWaitTime(
			config, res, project, "Updating TcpRoute", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkServicesTcpRouteRead(d, meta)
}

func resourceNetworkServicesTcpRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TcpRoute: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting TcpRoute %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "TcpRoute")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting TcpRoute", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TcpRoute %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesTcpRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/global/tcpRoutes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/tcpRoutes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesTcpRouteSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesTcpRouteDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteMeshes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteGateways(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"matches": flattenNetworkServicesTcpRouteRulesMatches(original["matches"], d, config),
			"action":  flattenNetworkServicesTcpRouteRulesAction(original["action"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTcpRouteRulesMatches(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"address": flattenNetworkServicesTcpRouteRulesMatchesAddress(original["address"], d, config),
			"port":    flattenNetworkServicesTcpRouteRulesMatchesPort(original["port"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTcpRouteRulesMatchesAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteRulesMatchesPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteRulesAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["destinations"] =
		flattenNetworkServicesTcpRouteRulesActionDestinations(original["destinations"], d, config)
	transformed["original_destination"] =
		flattenNetworkServicesTcpRouteRulesActionOriginalDestination(original["originalDestination"], d, config)
	transformed["idle_timeout"] =
		flattenNetworkServicesTcpRouteRulesActionIdleTimeout(original["idleTimeout"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkServicesTcpRouteRulesActionDestinations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"service_name": flattenNetworkServicesTcpRouteRulesActionDestinationsServiceName(original["serviceName"], d, config),
			"weight":       flattenNetworkServicesTcpRouteRulesActionDestinationsWeight(original["weight"], d, config),
		})
	}
	return transformed
}
func flattenNetworkServicesTcpRouteRulesActionDestinationsServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteRulesActionDestinationsWeight(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesTcpRouteRulesActionOriginalDestination(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteRulesActionIdleTimeout(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesTcpRouteTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesTcpRouteEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkServicesTcpRouteDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteMeshes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteGateways(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMatches, err := expandNetworkServicesTcpRouteRulesMatches(original["matches"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["matches"] = transformedMatches
		}

		transformedAction, err := expandNetworkServicesTcpRouteRulesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesMatches(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAddress, err := expandNetworkServicesTcpRouteRulesMatchesAddress(original["address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["address"] = transformedAddress
		}

		transformedPort, err := expandNetworkServicesTcpRouteRulesMatchesPort(original["port"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["port"] = transformedPort
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesMatchesAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesMatchesPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDestinations, err := expandNetworkServicesTcpRouteRulesActionDestinations(original["destinations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinations"] = transformedDestinations
	}

	transformedOriginalDestination, err := expandNetworkServicesTcpRouteRulesActionOriginalDestination(original["original_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOriginalDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["originalDestination"] = transformedOriginalDestination
	}

	transformedIdleTimeout, err := expandNetworkServicesTcpRouteRulesActionIdleTimeout(original["idle_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdleTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idleTimeout"] = transformedIdleTimeout
	}

	return transformed, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedServiceName, err := expandNetworkServicesTcpRouteRulesActionDestinationsServiceName(original["service_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceName"] = transformedServiceName
		}

		transformedWeight, err := expandNetworkServicesTcpRouteRulesActionDestinationsWeight(original["weight"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["weight"] = transformedWeight
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinationsServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionDestinationsWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionOriginalDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteRulesActionIdleTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesTcpRouteEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
