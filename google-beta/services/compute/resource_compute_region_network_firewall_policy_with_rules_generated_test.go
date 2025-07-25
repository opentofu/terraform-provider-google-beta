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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkFirewallPolicyWithRulesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesFullExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_firewall_policy_with_rules.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
}

resource "google_compute_region_network_firewall_policy_with_rules" "primary" {
  name        = "tf-test-fw-policy%{random_suffix}"
  region      = "us-west2"
  description = "Terraform test"

  rule {
    description    = "tcp rule"
    priority       = 1000
    enable_logging = true
    action         = "allow"
    direction      = "EGRESS"

    match {
      dest_ip_ranges            = ["11.100.0.1/32"]
      dest_fqdns                = ["www.yyy.com", "www.zzz.com"]
      dest_region_codes         = ["HK", "IN"]
      dest_threat_intelligences = ["iplist-search-engines-crawlers", "iplist-tor-exit-nodes"]
      dest_address_groups       = [google_network_security_address_group.address_group_1.id]

      layer4_config {
        ip_protocol = "tcp"
        ports       = [8080, 7070]
      }
    }

    target_secure_tag {
      name = google_tags_tag_value.secure_tag_value_1.id
    }
  }

  rule {
    description    = "udp rule"
    rule_name      = "test-rule"
    priority       = 2000
    enable_logging = false
    action         = "deny"
    direction      = "INGRESS"
    disabled       = true

    match {
      src_ip_ranges            = ["0.0.0.0/0"]
      src_fqdns                = ["www.abc.com", "www.def.com"]
      src_region_codes         = ["US", "CA"]
      src_threat_intelligences = ["iplist-known-malicious-ips", "iplist-public-clouds"]
      src_address_groups       = [google_network_security_address_group.address_group_1.id]

      src_secure_tag {
        name = google_tags_tag_value.secure_tag_value_1.id
      }

      layer4_config {
        ip_protocol = "udp"
      }
    }
  }
}

resource "google_network_security_address_group" "address_group_1" {
  name        = "tf-test-address-group%{random_suffix}"
  parent      = data.google_project.project.id
  description = "Regional address group"
  location    = "us-west2"
  items       = ["208.80.154.224/32"]
  type        = "IPV4"
  capacity    = 100
}

resource "google_tags_tag_key" "secure_tag_key_1" {
  description = "Tag key"
  parent      = data.google_project.project.id
  purpose     = "GCE_FIREWALL"
  short_name  = "tf-test-tag-key%{random_suffix}"
  purpose_data = {
    network = "${data.google_project.project.name}/default"
  }
}

resource "google_tags_tag_value" "secure_tag_value_1" {
  description = "Tag value"
  parent      = google_tags_tag_key.secure_tag_key_1.id
  short_name  = "tf-test-tag-value%{random_suffix}"
}
`, context)
}

func TestAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesRoceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkFirewallPolicyWithRulesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesRoceExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_firewall_policy_with_rules.policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkFirewallPolicyWithRules_computeRegionNetworkFirewallPolicyWithRulesRoceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_firewall_policy_with_rules" "policy" {
  provider = google-beta
  name        = "tf-test-rnf-policy%{random_suffix}"
  description = "Terraform test"
  policy_type = "RDMA_ROCE_POLICY"

  rule {
    description    = "deny all rule"
    priority       = 1000
    enable_logging = true
    action         = "deny"
    direction      = "INGRESS"

    match {
      src_ip_ranges            = ["0.0.0.0/0"]

      layer4_config {
        ip_protocol = "all"
      }
    }
  }
}
`, context)
}

func testAccCheckComputeRegionNetworkFirewallPolicyWithRulesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_network_firewall_policy_with_rules" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/firewallPolicies/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeRegionNetworkFirewallPolicyWithRules still exists at %s", url)
			}
		}

		return nil
	}
}
