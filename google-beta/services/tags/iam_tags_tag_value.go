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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/tags/TagValue.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/iam_policy.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package tags

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var TagsTagValueIamSchema = map[string]*schema.Schema{
	"tag_value": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type TagsTagValueIamUpdater struct {
	tagValue string
	d        tpgresource.TerraformResourceData
	Config   *transport_tpg.Config
}

func TagsTagValueIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("tag_value"); ok {
		values["tag_value"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"tagValues/(?P<tag_value>[^/]+)", "(?P<tag_value>[^/]+)"}, d, config, d.Get("tag_value").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &TagsTagValueIamUpdater{
		tagValue: values["tag_value"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("tag_value", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting tag_value: %s", err)
	}

	return u, nil
}

func TagsTagValueIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := tpgresource.GetImportIdQualifiers([]string{"tagValues/(?P<tag_value>[^/]+)", "(?P<tag_value>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &TagsTagValueIamUpdater{
		tagValue: values["tag_value"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("tag_value", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting tag_value: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *TagsTagValueIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTagValueUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *TagsTagValueIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTagValueUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *TagsTagValueIamUpdater) qualifyTagValueUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{TagsBasePath}}%s:%s", fmt.Sprintf("tagValues/%s", u.tagValue), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *TagsTagValueIamUpdater) GetResourceId() string {
	return fmt.Sprintf("tagValues/%s", u.tagValue)
}

func (u *TagsTagValueIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-tags-tagvalue-%s", u.GetResourceId())
}

func (u *TagsTagValueIamUpdater) DescribeResource() string {
	return fmt.Sprintf("tags tagvalue %q", u.GetResourceId())
}
