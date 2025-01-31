// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccComputeNetworkFirewallPolicyPacketMirroringRule_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkFirewallPolicyPacketMirroringRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkFirewallPolicyPacketMirroringRule_basic(context),
			},
			{
				ResourceName:            "google_compute_network_firewall_policy_packet_mirroring_rule.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"firewall_policy"},
			},
			{
				Config: testAccComputeNetworkFirewallPolicyPacketMirroringRule_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_compute_network_firewall_policy_packet_mirroring_rule.primary", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_compute_network_firewall_policy_packet_mirroring_rule.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"firewall_policy"},
			},
		},
	})
}

func testAccComputeNetworkFirewallPolicyPacketMirroringRule_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network_firewall_policy" "basic_network_firewall_policy" {
  provider    = google-beta
  name        = "tf-test-fw-policy%{random_suffix}"
  description = "Sample global network firewall policy"
  project     = "%{project_name}"
}

resource "google_compute_network_firewall_policy_packet_mirroring_rule" "primary" {
  provider                = google-beta
  action                  = "do_not_mirror"
  direction               = "INGRESS"
  firewall_policy         = google_compute_network_firewall_policy.basic_network_firewall_policy.name
  priority                = 1000

  match {
    src_ip_ranges = ["10.100.0.1/32"]
    layer4_configs {
      ip_protocol = "all"
    }
  }
}

`, context)
}

func testAccComputeNetworkFirewallPolicyPacketMirroringRule_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network_firewall_policy" "basic_network_firewall_policy" {
  provider    = google-beta
  name        = "tf-test-fw-policy%{random_suffix}"
  description = "Sample global network firewall policy"
  project     = "%{project_name}"
}

resource "google_compute_network_firewall_policy_packet_mirroring_rule" "primary" {
  provider                = google-beta
  action                  = "mirror"
  description             = "This is a simple packet mirroring rule description"
  direction               = "INGRESS"
  disabled                = true
  firewall_policy         = google_compute_network_firewall_policy.basic_network_firewall_policy.name
  priority                = 1000
  rule_name               = "test-rule"

  match {
    src_ip_ranges = ["192.168.0.1/32"]
    layer4_configs {
      ip_protocol = "tcp"
    }
  }

  security_profile_group = "//networksecurity.googleapis.com/${google_network_security_security_profile_group.security_profile_group_1.id}"

  target_secure_tags {
    name = "tagValues/${google_tags_tag_value.secure_tag_value_1.name}"
  }

}

resource "google_network_security_mirroring_deployment_group" "default" {
  provider                      = google-beta
  mirroring_deployment_group_id = "tf-test-deployment-group%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.default.id
}

resource "google_network_security_mirroring_endpoint_group" "default" {
  provider                      = google-beta
  mirroring_endpoint_group_id   = "tf-test-endpoint-group%{random_suffix}"
  location                      = "global"
  mirroring_deployment_group    = google_network_security_mirroring_deployment_group.default.id
}

resource "google_network_security_security_profile" "default" {
  provider    = google-beta
  name        = "tf-test-sec-profile%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "my description"
  type        = "CUSTOM_MIRRORING"

  custom_mirroring_profile {
    mirroring_endpoint_group = google_network_security_mirroring_endpoint_group.default.id
  }
}

resource "google_network_security_security_profile_group" "security_profile_group_1" {
  provider                 = google-beta
  name                     = "tf-test-sec-profile-group%{random_suffix}"
  parent                   = "organizations/%{org_id}"
  description              = "my description"
  custom_mirroring_profile = google_network_security_security_profile.default.id
}

resource "google_tags_tag_key" "secure_tag_key_1" {
  provider    = google-beta
  description = "Test tag key description"
  parent      = "organizations/%{org_id}"
  purpose     = "GCE_FIREWALL"
  short_name  = "tf-test-tag-key%{random_suffix}"
  purpose_data = {
    network = "%{project_name}/${google_compute_network.default.name}"
  }
}

resource "google_tags_tag_value" "secure_tag_value_1" {
  provider    = google-beta
  description = "Test tag value description."
  parent      = google_tags_tag_key.secure_tag_key_1.id
  short_name  = "tf-test-tag-value%{random_suffix}"
}
`, context)
}
