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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/redis/ClusterUserCreatedConnections.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package redis

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

func ResourceRedisClusterUserCreatedConnections() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisClusterUserCreatedConnectionsCreate,
		Read:   resourceRedisClusterUserCreatedConnectionsRead,
		Update: resourceRedisClusterUserCreatedConnectionsUpdate,
		Delete: resourceRedisClusterUserCreatedConnectionsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisClusterUserCreatedConnectionsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(120 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the Redis cluster these endpoints should be added to.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the region of the Redis cluster these endpoints should be added to.`,
			},
			"cluster_endpoints": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `A list of cluster endpoints`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: ``,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"psc_connection": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Detailed information of a PSC connection that is created by the customer
who owns the cluster.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The IP allocated on the consumer network for the PSC forwarding rule.`,
												},
												"forwarding_rule": {
													Type:     schema.TypeString,
													Required: true,
													Description: `The URI of the consumer side forwarding rule.
Format:
projects/{project}/regions/{region}/forwardingRules/{forwarding_rule}`,
												},
												"network": {
													Type:     schema.TypeString,
													Required: true,
													Description: `The consumer network where the IP address resides, in the form of
projects/{project_id}/global/networks/{network_id}.`,
												},
												"psc_connection_id": {
													Type:     schema.TypeString,
													Required: true,
													Description: `The PSC connection id of the forwarding rule connected to the
service attachment.`,
												},
												"service_attachment": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `The service attachment which is the target of the PSC connection, in the form of projects/{project-id}/regions/{region}/serviceAttachments/{service-attachment-id}.`,
												},
												"project_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Optional:    true,
													Description: `The consumer project_id where the forwarding rule is created from.`,
												},
												"connection_type": {
													Type:     schema.TypeString,
													Computed: true,
													Description: `Output Only. Type of a PSC Connection. 
 Possible values:
 CONNECTION_TYPE_DISCOVERY 
 CONNECTION_TYPE_PRIMARY 
 CONNECTION_TYPE_READER`,
												},
												"psc_connection_status": {
													Type:     schema.TypeString,
													Computed: true,
													Description: `Output Only. The status of the PSC connection: whether a connection exists and ACTIVE or it no longer exists. 
 Possible values:
 ACTIVE 
 NOT_FOUND`,
												},
											},
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

func resourceRedisClusterUserCreatedConnectionsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	clusterEndpointsProp, err := expandRedisClusterUserCreatedConnectionsClusterEndpoints(d.Get("cluster_endpoints"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_endpoints"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterEndpointsProp)) && (ok || !reflect.DeepEqual(v, clusterEndpointsProp)) {
		obj["clusterEndpoints"] = clusterEndpointsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}?updateMask=cluster_endpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ClusterUserCreatedConnections: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClusterUserCreatedConnections: %s", err)
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
		return fmt.Errorf("Error creating ClusterUserCreatedConnections: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = RedisOperationWaitTime(
		config, res, project, "Creating ClusterUserCreatedConnections", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ClusterUserCreatedConnections: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ClusterUserCreatedConnections %q: %#v", d.Id(), res)

	return resourceRedisClusterUserCreatedConnectionsRead(d, meta)
}

func resourceRedisClusterUserCreatedConnectionsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClusterUserCreatedConnections: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("RedisClusterUserCreatedConnections %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ClusterUserCreatedConnections: %s", err)
	}

	if err := d.Set("cluster_endpoints", flattenRedisClusterUserCreatedConnectionsClusterEndpoints(res["clusterEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading ClusterUserCreatedConnections: %s", err)
	}

	return nil
}

func resourceRedisClusterUserCreatedConnectionsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClusterUserCreatedConnections: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	clusterEndpointsProp, err := expandRedisClusterUserCreatedConnectionsClusterEndpoints(d.Get("cluster_endpoints"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_endpoints"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clusterEndpointsProp)) {
		obj["clusterEndpoints"] = clusterEndpointsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}?updateMask=cluster_endpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ClusterUserCreatedConnections %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("cluster_endpoints") {
		updateMask = append(updateMask, "clusterEndpoints")
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
			return fmt.Errorf("Error updating ClusterUserCreatedConnections %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ClusterUserCreatedConnections %q: %#v", d.Id(), res)
		}

		err = RedisOperationWaitTime(
			config, res, project, "Updating ClusterUserCreatedConnections", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceRedisClusterUserCreatedConnectionsRead(d, meta)
}

func resourceRedisClusterUserCreatedConnectionsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ClusterUserCreatedConnections: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	// not setting clusterEndpoints in obj

	url, err := tpgresource.ReplaceVars(d, config, "{{RedisBasePath}}projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ClusterUserCreatedConnections %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	updateMask := []string{}
	updateMask = append(updateMask, "clusterEndpoints")
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

	obj["async_cluster_endpoints_deletion_enabled"] = true

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
			return fmt.Errorf("Error updating ClusterUserCreatedConnections %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ClusterUserCreatedConnections %q: %#v", d.Id(), res)
		}

		err = RedisOperationWaitTime(
			config, res, project, "Updating ClusterUserCreatedConnections", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceRedisClusterUserCreatedConnectionsRead(d, meta)
}

func resourceRedisClusterUserCreatedConnectionsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/clusters/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
		connection := flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnections(original["connections"], d, config)
		if len(connection.([]interface{})) == 0 {
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"connections": connection,
		})
	}
	return transformed
}
func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnections(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
		pscConnections := flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection(original["pscConnection"], d, config)
		if pscConnections == nil {
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"psc_connection": pscConnections,
		})

	}
	return transformed
}
func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["psc_connection_id"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionId(original["pscConnectionId"], d, config)
	transformed["address"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionAddress(original["address"], d, config)
	transformed["forwarding_rule"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionForwardingRule(original["forwardingRule"], d, config)
	transformed["project_id"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionProjectId(original["projectId"], d, config)
	transformed["network"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionNetwork(original["network"], d, config)
	transformed["service_attachment"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionServiceAttachment(original["serviceAttachment"], d, config)
	transformed["psc_connection_status"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionStatus(original["pscConnectionStatus"], d, config)
	transformed["connection_type"] =
		flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionConnectionType(original["connectionType"], d, config)
	return []interface{}{transformed}
}
func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionForwardingRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionServiceAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionConnectionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandRedisClusterUserCreatedConnectionsClusterEndpoints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedConnections, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnections(original["connections"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnections); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["connections"] = transformedConnections
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnections(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPscConnection, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection(original["psc_connection"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPscConnection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pscConnection"] = transformedPscConnection
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPscConnectionId, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionId(original["psc_connection_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPscConnectionId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pscConnectionId"] = transformedPscConnectionId
	}

	transformedAddress, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionAddress(original["address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["address"] = transformedAddress
	}

	transformedForwardingRule, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionForwardingRule(original["forwarding_rule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForwardingRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["forwardingRule"] = transformedForwardingRule
	}

	transformedProjectId, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedNetwork, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedServiceAttachment, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionServiceAttachment(original["service_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachment"] = transformedServiceAttachment
	}

	transformedPscConnectionStatus, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionStatus(original["psc_connection_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPscConnectionStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pscConnectionStatus"] = transformedPscConnectionStatus
	}

	transformedConnectionType, err := expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionConnectionType(original["connection_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectionType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["connectionType"] = transformedConnectionType
	}

	return transformed, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionForwardingRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionPscConnectionStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterUserCreatedConnectionsClusterEndpointsConnectionsPscConnectionConnectionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
