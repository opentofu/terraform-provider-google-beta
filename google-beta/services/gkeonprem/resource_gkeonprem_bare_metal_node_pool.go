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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkeonprem/BareMetalNodePool.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gkeonprem

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

func ResourceGkeonpremBareMetalNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkeonpremBareMetalNodePoolCreate,
		Read:   resourceGkeonpremBareMetalNodePoolRead,
		Update: resourceGkeonpremBareMetalNodePoolUpdate,
		Delete: resourceGkeonpremBareMetalNodePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkeonpremBareMetalNodePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"bare_metal_cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The cluster this node pool belongs to.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the resource.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The bare metal node pool name.`,
			},
			"node_pool_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Node pool configuration.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_configs": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The list of machine addresses in the Bare Metal Node Pool.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"labels": {
										Type:     schema.TypeMap,
										Optional: true,
										Description: `The map of Kubernetes labels (key/value pairs) to be applied to
each node. These will added in addition to any default label(s)
that Kubernetes may apply to the node. In case of conflict in
label keys, the applied set may differ depending on the Kubernetes
version -- it's best to assume the behavior is undefined and
conflicts should be avoided. For more information, including usage
and the valid values, see:
  - http://kubernetes.io/v1.1/docs/user-guide/labels.html
An object containing a list of "key": value pairs.
For example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"node_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The default IPv4 address for SSH access and Kubernetes node.
Example: 192.168.0.1`,
									},
								},
							},
						},
						"labels": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Description: `The map of Kubernetes labels (key/value pairs) to be applied to
each node. These will added in addition to any default label(s)
that Kubernetes may apply to the node. In case of conflict in
label keys, the applied set may differ depending on the Kubernetes
version -- it's best to assume the behavior is undefined and
conflicts should be avoided. For more information, including usage
and the valid values, see:
  - http://kubernetes.io/v1.1/docs/user-guide/labels.html
An object containing a list of "key": value pairs.
For example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						"operating_system": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `Specifies the nodes operating system (default: LINUX).`,
						},
						"taints": {
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: `The initial taints assigned to nodes of this node pool.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"EFFECT_UNSPECIFIED", "PREFER_NO_SCHEDULE", "NO_EXECUTE", ""}),
										Description:  `Specifies the nodes operating system (default: LINUX). Possible values: ["EFFECT_UNSPECIFIED", "PREFER_NO_SCHEDULE", "NO_EXECUTE"]`,
									},
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Key associated with the effect.`,
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Value associated with the effect.`,
									},
								},
							},
						},
					},
				},
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Annotations on the Bare Metal Node Pool.
This field has the same restrictions as Kubernetes annotations.
The total size of all keys and values combined is limited to 256k.
Key can have 2 segments: prefix (optional) and name (required),
separated by a slash (/).
Prefix must be a DNS subdomain.
Name must be 63 characters or less, begin and end with alphanumerics,
with dashes (-), underscores (_), dots (.), and alphanumerics between.


**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name for the Bare Metal Node Pool.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the cluster was created, in RFC3339 text format.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the cluster was deleted, in RFC3339 text format.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `This checksum is computed by the server based on the value of other
fields, and may be sent on update and delete requests to ensure the
client has an up-to-date value before proceeding.
Allows clients to perform consistent read-modify-writes
through optimistic concurrency control.`,
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If set, there are currently changes in flight to the Bare Metal User Cluster.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of this cluster.`,
			},
			"status": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Specifies detailed node pool status.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conditions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `ResourceConditions provide a standard mechanism for higher-level status reporting from user cluster controller.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Human-readable message indicating details about last transition.`,
									},
									"reason": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Machine-readable message indicating details about last transition.`,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Type of the condition.
(e.g., ClusterRunning, NodePoolRunning or ServerSidePreflightReady)`,
									},
									"last_transition_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Last time the condition transit from one status to another.`,
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The lifecycle state of the condition.`,
									},
								},
							},
						},
						"error_message": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Human-friendly representation of the error message from the user cluster
controller. The error message can be temporary as the user cluster
controller creates a cluster or node pool. If the error message persists
for a longer period of time, it can be used to surface error message to
indicate real problems requiring user intervention.`,
						},
					},
				},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The unique identifier of the Bare Metal Node Pool.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the cluster was last updated, in RFC3339 text format.`,
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

func resourceGkeonpremBareMetalNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandGkeonpremBareMetalNodePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodePoolConfigProp, err := expandGkeonpremBareMetalNodePoolNodePoolConfig(d.Get("node_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodePoolConfigProp)) && (ok || !reflect.DeepEqual(v, nodePoolConfigProp)) {
		obj["nodePoolConfig"] = nodePoolConfigProp
	}
	annotationsProp, err := expandGkeonpremBareMetalNodePoolEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools?bare_metal_node_pool_id={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new BareMetalNodePool: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BareMetalNodePool: %s", err)
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
		return fmt.Errorf("Error creating BareMetalNodePool: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = GkeonpremOperationWaitTime(
		config, res, project, "Creating BareMetalNodePool", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {

		return fmt.Errorf("Error waiting to create BareMetalNodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating BareMetalNodePool %q: %#v", d.Id(), res)

	return resourceGkeonpremBareMetalNodePoolRead(d, meta)
}

func resourceGkeonpremBareMetalNodePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BareMetalNodePool: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GkeonpremBareMetalNodePool %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}

	if err := d.Set("display_name", flattenGkeonpremBareMetalNodePoolDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("annotations", flattenGkeonpremBareMetalNodePoolAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("node_pool_config", flattenGkeonpremBareMetalNodePoolNodePoolConfig(res["nodePoolConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("status", flattenGkeonpremBareMetalNodePoolStatus(res["status"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("uid", flattenGkeonpremBareMetalNodePoolUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("state", flattenGkeonpremBareMetalNodePoolState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("reconciling", flattenGkeonpremBareMetalNodePoolReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("create_time", flattenGkeonpremBareMetalNodePoolCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("update_time", flattenGkeonpremBareMetalNodePoolUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("delete_time", flattenGkeonpremBareMetalNodePoolDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("etag", flattenGkeonpremBareMetalNodePoolEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}
	if err := d.Set("effective_annotations", flattenGkeonpremBareMetalNodePoolEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading BareMetalNodePool: %s", err)
	}

	return nil
}

func resourceGkeonpremBareMetalNodePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BareMetalNodePool: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandGkeonpremBareMetalNodePoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodePoolConfigProp, err := expandGkeonpremBareMetalNodePoolNodePoolConfig(d.Get("node_pool_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_pool_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodePoolConfigProp)) {
		obj["nodePoolConfig"] = nodePoolConfigProp
	}
	annotationsProp, err := expandGkeonpremBareMetalNodePoolEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating BareMetalNodePool %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("node_pool_config") {
		updateMask = append(updateMask, "nodePoolConfig")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
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
			return fmt.Errorf("Error updating BareMetalNodePool %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating BareMetalNodePool %q: %#v", d.Id(), res)
		}

		err = GkeonpremOperationWaitTime(
			config, res, project, "Updating BareMetalNodePool", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceGkeonpremBareMetalNodePoolRead(d, meta)
}

func resourceGkeonpremBareMetalNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for BareMetalNodePool: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting BareMetalNodePool %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "BareMetalNodePool")
	}

	err = GkeonpremOperationWaitTime(
		config, res, project, "Deleting BareMetalNodePool", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting BareMetalNodePool %q: %#v", d.Id(), res)
	return nil
}

func resourceGkeonpremBareMetalNodePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/bareMetalClusters/(?P<bare_metal_cluster>[^/]+)/bareMetalNodePools/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<bare_metal_cluster>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<bare_metal_cluster>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGkeonpremBareMetalNodePoolDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["node_configs"] =
		flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(original["nodeConfigs"], d, config)
	transformed["operating_system"] =
		flattenGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(original["operatingSystem"], d, config)
	transformed["taints"] =
		flattenGkeonpremBareMetalNodePoolNodePoolConfigTaints(original["taints"], d, config)
	transformed["labels"] =
		flattenGkeonpremBareMetalNodePoolNodePoolConfigLabels(original["labels"], d, config)
	return []interface{}{transformed}
}
func flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"node_ip": flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(original["nodeIp"], d, config),
			"labels":  flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(original["labels"], d, config),
		})
	}
	return transformed
}
func flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigTaints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"key":    flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(original["key"], d, config),
			"value":  flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(original["value"], d, config),
			"effect": flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(original["effect"], d, config),
		})
	}
	return transformed
}
func flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolNodePoolConfigLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["error_message"] =
		flattenGkeonpremBareMetalNodePoolStatusErrorMessage(original["errorMessage"], d, config)
	transformed["conditions"] =
		flattenGkeonpremBareMetalNodePoolStatusConditions(original["conditions"], d, config)
	return []interface{}{transformed}
}
func flattenGkeonpremBareMetalNodePoolStatusErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatusConditions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"type":                 flattenGkeonpremBareMetalNodePoolStatusConditionsType(original["type"], d, config),
			"reason":               flattenGkeonpremBareMetalNodePoolStatusConditionsReason(original["reason"], d, config),
			"message":              flattenGkeonpremBareMetalNodePoolStatusConditionsMessage(original["message"], d, config),
			"last_transition_time": flattenGkeonpremBareMetalNodePoolStatusConditionsLastTransitionTime(original["lastTransitionTime"], d, config),
			"state":                flattenGkeonpremBareMetalNodePoolStatusConditionsState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenGkeonpremBareMetalNodePoolStatusConditionsType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatusConditionsReason(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatusConditionsMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatusConditionsLastTransitionTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolStatusConditionsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGkeonpremBareMetalNodePoolEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGkeonpremBareMetalNodePoolDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNodeConfigs, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(original["node_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeConfigs"] = transformedNodeConfigs
	}

	transformedOperatingSystem, err := expandGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(original["operating_system"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperatingSystem); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operatingSystem"] = transformedOperatingSystem
	}

	transformedTaints, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaints(original["taints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTaints); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["taints"] = transformedTaints
	}

	transformedLabels, err := expandGkeonpremBareMetalNodePoolNodePoolConfigLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNodeIp, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(original["node_ip"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNodeIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nodeIp"] = transformedNodeIp
		}

		transformedLabels, err := expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(original["labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["labels"] = transformedLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsNodeIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigNodeConfigsLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigOperatingSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValue, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		transformedEffect, err := expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigTaintsEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGkeonpremBareMetalNodePoolNodePoolConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGkeonpremBareMetalNodePoolEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
