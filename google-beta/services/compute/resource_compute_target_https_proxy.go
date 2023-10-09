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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceComputeTargetHttpsProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetHttpsProxyCreate,
		Read:   resourceComputeTargetHttpsProxyRead,
		Update: resourceComputeTargetHttpsProxyUpdate,
		Delete: resourceComputeTargetHttpsProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetHttpsProxyImport,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"url_map": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `A reference to the UrlMap resource that defines the mapping from URL
to the BackendService.`,
			},
			"certificate_manager_certificates": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description: `A list of Certificate Manager certificate URLs that are used to authenticate
connections between users and the load balancer. At least one resource must be specified.
Accepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificates/{resourceName}' or just the self_link projects/{project}/locations/{location}/certificates/{resourceName}`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ConflictsWith: []string{"ssl_certificates"},
			},
			"certificate_map": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A reference to the CertificateMap resource uri that identifies a certificate map
associated with the given target proxy. This field can only be set for global target proxies.
Accepted format is '//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}'.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"http_keep_alive_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `Specifies how long to keep a connection open, after completing a response,
while there is no matching traffic (in seconds). If an HTTP keepalive is
not specified, a default value (610 seconds) will be used. For Global
external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
the maximum allowed value is 1200 seconds. For Global external HTTP(S)
load balancer (classic), this option is not available publicly.`,
			},
			"proxy_bind": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `This field only applies when the forwarding rule that references
this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.`,
			},
			"quic_override": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"NONE", "ENABLE", "DISABLE", ""}),
				Description: `Specifies the QUIC override policy for this resource. This determines
whether the load balancer will attempt to negotiate QUIC with clients
or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
specified, Google manages whether QUIC is used. Default value: "NONE" Possible values: ["NONE", "ENABLE", "DISABLE"]`,
				Default: "NONE",
			},
			"server_tls_policy": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `A URL referring to a networksecurity.ServerTlsPolicy
resource that describes how the proxy should authenticate inbound
traffic. serverTlsPolicy only applies to a global TargetHttpsProxy
attached to globalForwardingRules with the loadBalancingScheme
set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.
For details which ServerTlsPolicy resources are accepted with
INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED
loadBalancingScheme consult ServerTlsPolicy documentation.
If left blank, communications are not encrypted.`,
			},
			"ssl_certificates": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of SslCertificate resource URLs that are used to authenticate
connections between users and the load balancer. At least one resource must be specified.`,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				},
				ConflictsWith: []string{"certificate_manager_certificates"},
			},
			"ssl_policy": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `A reference to the SslPolicy resource that will be associated with
the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
resource will not have any SSL policy configured.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"proxy_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
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

func resourceComputeTargetHttpsProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetHttpsProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeTargetHttpsProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	quicOverrideProp, err := expandComputeTargetHttpsProxyQuicOverride(d.Get("quic_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quic_override"); !tpgresource.IsEmptyValue(reflect.ValueOf(quicOverrideProp)) && (ok || !reflect.DeepEqual(v, quicOverrideProp)) {
		obj["quicOverride"] = quicOverrideProp
	}
	certificateManagerCertificatesProp, err := expandComputeTargetHttpsProxyCertificateManagerCertificates(d.Get("certificate_manager_certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate_manager_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateManagerCertificatesProp)) && (ok || !reflect.DeepEqual(v, certificateManagerCertificatesProp)) {
		obj["certificateManagerCertificates"] = certificateManagerCertificatesProp
	}
	sslCertificatesProp, err := expandComputeTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslCertificatesProp)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
		obj["sslCertificates"] = sslCertificatesProp
	}
	certificateMapProp, err := expandComputeTargetHttpsProxyCertificateMap(d.Get("certificate_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateMapProp)) && (ok || !reflect.DeepEqual(v, certificateMapProp)) {
		obj["certificateMap"] = certificateMapProp
	}
	sslPolicyProp, err := expandComputeTargetHttpsProxySslPolicy(d.Get("ssl_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(sslPolicyProp)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
		obj["sslPolicy"] = sslPolicyProp
	}
	urlMapProp, err := expandComputeTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("url_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	proxyBindProp, err := expandComputeTargetHttpsProxyProxyBind(d.Get("proxy_bind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("proxy_bind"); !tpgresource.IsEmptyValue(reflect.ValueOf(proxyBindProp)) && (ok || !reflect.DeepEqual(v, proxyBindProp)) {
		obj["proxyBind"] = proxyBindProp
	}
	httpKeepAliveTimeoutSecProp, err := expandComputeTargetHttpsProxyHttpKeepAliveTimeoutSec(d.Get("http_keep_alive_timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_keep_alive_timeout_sec"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpKeepAliveTimeoutSecProp)) && (ok || !reflect.DeepEqual(v, httpKeepAliveTimeoutSecProp)) {
		obj["httpKeepAliveTimeoutSec"] = httpKeepAliveTimeoutSecProp
	}
	serverTlsPolicyProp, err := expandComputeTargetHttpsProxyServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}

	obj, err = resourceComputeTargetHttpsProxyEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetHttpsProxy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetHttpsProxy: %s", err)
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
		return fmt.Errorf("Error creating TargetHttpsProxy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating TargetHttpsProxy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TargetHttpsProxy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TargetHttpsProxy %q: %#v", d.Id(), res)

	return resourceComputeTargetHttpsProxyRead(d, meta)
}

func resourceComputeTargetHttpsProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetHttpsProxy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeTargetHttpsProxy %q", d.Id()))
	}

	res, err = resourceComputeTargetHttpsProxyDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeTargetHttpsProxy because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeTargetHttpsProxyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeTargetHttpsProxyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeTargetHttpsProxyProxyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeTargetHttpsProxyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("quic_override", flattenComputeTargetHttpsProxyQuicOverride(res["quicOverride"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("certificate_manager_certificates", flattenComputeTargetHttpsProxyCertificateManagerCertificates(res["certificateManagerCertificates"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("ssl_certificates", flattenComputeTargetHttpsProxySslCertificates(res["sslCertificates"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("certificate_map", flattenComputeTargetHttpsProxyCertificateMap(res["certificateMap"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("ssl_policy", flattenComputeTargetHttpsProxySslPolicy(res["sslPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("url_map", flattenComputeTargetHttpsProxyUrlMap(res["urlMap"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("proxy_bind", flattenComputeTargetHttpsProxyProxyBind(res["proxyBind"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("http_keep_alive_timeout_sec", flattenComputeTargetHttpsProxyHttpKeepAliveTimeoutSec(res["httpKeepAliveTimeoutSec"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("server_tls_policy", flattenComputeTargetHttpsProxyServerTlsPolicy(res["serverTlsPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}

	return nil
}

func resourceComputeTargetHttpsProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetHttpsProxy: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("quic_override") {
		obj := make(map[string]interface{})

		quicOverrideProp, err := expandComputeTargetHttpsProxyQuicOverride(d.Get("quic_override"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("quic_override"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, quicOverrideProp)) {
			obj["quicOverride"] = quicOverrideProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}/setQuicOverride")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("certificate_manager_certificates") || d.HasChange("ssl_certificates") {
		obj := make(map[string]interface{})

		certificateManagerCertificatesProp, err := expandComputeTargetHttpsProxyCertificateManagerCertificates(d.Get("certificate_manager_certificates"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("certificate_manager_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, certificateManagerCertificatesProp)) {
			obj["certificateManagerCertificates"] = certificateManagerCertificatesProp
		}
		sslCertificatesProp, err := expandComputeTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_certificates"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
			obj["sslCertificates"] = sslCertificatesProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/targetHttpsProxies/{{name}}/setSslCertificates")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("certificate_map") {
		obj := make(map[string]interface{})

		certificateMapProp, err := expandComputeTargetHttpsProxyCertificateMap(d.Get("certificate_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("certificate_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, certificateMapProp)) {
			obj["certificateMap"] = certificateMapProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}/setCertificateMap")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("ssl_policy") {
		obj := make(map[string]interface{})

		sslPolicyProp, err := expandComputeTargetHttpsProxySslPolicy(d.Get("ssl_policy"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
			obj["sslPolicy"] = sslPolicyProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}/setSslPolicy")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("url_map") {
		obj := make(map[string]interface{})

		urlMapProp, err := expandComputeTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("url_map"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
			obj["urlMap"] = urlMapProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/targetHttpsProxies/{{name}}/setUrlMap")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TargetHttpsProxy %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating TargetHttpsProxy", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeTargetHttpsProxyRead(d, meta)
}

func resourceComputeTargetHttpsProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TargetHttpsProxy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TargetHttpsProxy %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "TargetHttpsProxy")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting TargetHttpsProxy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TargetHttpsProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeTargetHttpsProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/global/targetHttpsProxies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeTargetHttpsProxyCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyProxyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeTargetHttpsProxyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyQuicOverride(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return "NONE"
	}

	return v
}

func flattenComputeTargetHttpsProxyCertificateManagerCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxySslCertificates(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeTargetHttpsProxyCertificateMap(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxySslPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeTargetHttpsProxyUrlMap(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeTargetHttpsProxyProxyBind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyHttpKeepAliveTimeoutSec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenComputeTargetHttpsProxyServerTlsPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func expandComputeTargetHttpsProxyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyQuicOverride(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyCertificateManagerCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for certificate_manager_certificates: nil")
		}
		if strings.HasPrefix(raw.(string), "//") || strings.HasPrefix(raw.(string), "https://") {
			// Any full URL will be passed to the API request (regardless of the resource type). This is to allow self_links of CertificateManagerCeritificate resources.
			// If the full URL is an invalid reference, that should be handled by the API.
			req = append(req, raw.(string))
		} else if reg, _ := regexp.Compile("projects/(.*)/locations/(.*)/certificates/(.*)"); reg.MatchString(raw.(string)) {
			// If the input is the id pattern of CertificateManagerCertificate resource, a prefix will be added to construct the full URL before constructing the API request.
			self_link := "https://certificatemanager.googleapis.com/v1/" + raw.(string)
			req = append(req, self_link)
		} else {
			return nil, fmt.Errorf("Invalid value for certificate_manager_certificates: %v is an invalid format for a certificateManagerCertificate resource", raw.(string))
		}
	}
	return req, nil
}

func expandComputeTargetHttpsProxySslCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: nil")
		}
		f, err := tpgresource.ParseGlobalFieldValue("sslCertificates", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeTargetHttpsProxyCertificateMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxySslPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("sslPolicies", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for ssl_policy: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetHttpsProxyUrlMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("urlMaps", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetHttpsProxyProxyBind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyHttpKeepAliveTimeoutSec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyServerTlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeTargetHttpsProxyEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	if _, ok := obj["certificateManagerCertificates"]; ok {
		// The field certificateManagerCertificates should not be included in the API request, and it should be renamed to `sslCertificates`
		// The API does not allow using both certificate manager certificates and sslCertificates. If that changes
		// in the future, the encoder logic should change accordingly because this will mean that both fields are no longer mutual exclusive.
		log.Printf("[DEBUG] converting the field CertificateManagerCertificates to sslCertificates before sending the request")
		obj["sslCertificates"] = obj["certificateManagerCertificates"]
		delete(obj, "certificateManagerCertificates")
	}
	return obj, nil
}

func resourceComputeTargetHttpsProxyDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Since both sslCertificates and certificateManagerCertificates maps to the same API field (sslCertificates), we need to check the types
	// of certificates that exist in the array and decide whether to change the field to certificateManagerCertificate or not.
	// The decoder logic depends on the fact that the API does not allow mixed type of certificates and it returns
	// certificate manager certificates in the format of //certificatemanager.googleapis.com/projects/*/locations/*/certificates/*
	if sslCertificates, ok := res["sslCertificates"].([]interface{}); ok && len(sslCertificates) > 0 {
		regPat, _ := regexp.Compile("//certificatemanager.googleapis.com/projects/(.*)/locations/(.*)/certificates/(.*)")

		if regPat.MatchString(sslCertificates[0].(string)) {
			// It is enough to check only the type of one of the provided certificates beacuse all the certificates should be the same type.
			log.Printf("[DEBUG] The field sslCertificates contains certificateManagerCertificates, the field name will be converted to certificateManagerCertificates")
			res["certificateManagerCertificates"] = res["sslCertificates"]
			delete(res, "sslCertificates")
		}
	}
	return res, nil
}
