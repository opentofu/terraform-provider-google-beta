// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/sourcerepo/source_repo_utils.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package sourcerepo

import (
	"regexp"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsub"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func expandSourceRepoRepositoryPubsubConfigsTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (string, error) {
	// short-circuit if the topic is a full uri so we don't need to GetProject
	ok, err := regexp.MatchString(pubsub.PubsubTopicRegex, v.(string))
	if err != nil {
		return "", err
	}

	if ok {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return "", err
	}

	return pubsub.GetComputedTopicName(project, v.(string)), err
}
