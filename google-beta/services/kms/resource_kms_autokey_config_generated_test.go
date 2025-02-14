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

package kms_test

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

func TestAccKMSAutokeyConfig_kmsAutokeyConfigAllExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckKMSAutokeyConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccKMSAutokeyConfig_kmsAutokeyConfigAllExample(context),
			},
			{
				ResourceName:            "google_kms_autokey_config.example-autokeyconfig",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"folder"},
			},
		},
	})
}

func testAccKMSAutokeyConfig_kmsAutokeyConfigAllExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# Create Folder in GCP Organization
resource "google_folder" "autokms_folder" {
  provider     = google-beta
  display_name = "tf-test-my-folder%{random_suffix}"
  parent       = "organizations/%{org_id}"
  deletion_protection = false
}

# Create the key project
resource "google_project" "key_project" {
  provider        = google-beta
  project_id      = "tf-test-key-proj%{random_suffix}"
  name            = "tf-test-key-proj%{random_suffix}"
  folder_id       = google_folder.autokms_folder.folder_id
  billing_account = "%{billing_account}"
  depends_on      = [google_folder.autokms_folder]
  deletion_policy = "DELETE"
}

# Enable the Cloud KMS API
resource "google_project_service" "kms_api_service" {
  provider                   = google-beta
  service                    = "cloudkms.googleapis.com"
  project                    = google_project.key_project.project_id
  disable_on_destroy         = false
  disable_dependent_services = true
  depends_on                 = [google_project.key_project]
}

# Wait delay after enabling APIs
resource "time_sleep" "wait_enable_service_api" {
  depends_on       = [google_project_service.kms_api_service]
  create_duration  = "30s"
}

#Create KMS Service Agent
resource "google_project_service_identity" "kms_service_agent" {
  provider   = google-beta
  service    = "cloudkms.googleapis.com"
  project    = google_project.key_project.number
  depends_on = [time_sleep.wait_enable_service_api]
}

# Wait delay after creating service agent.
resource "time_sleep" "wait_service_agent" {
  depends_on       = [google_project_service_identity.kms_service_agent]
  create_duration  = "10s"
}

#Grant the KMS Service Agent the Cloud KMS Admin role
resource "google_project_iam_member" "autokey_project_admin" {
  provider   = google-beta
  project    = google_project.key_project.project_id
  role       = "roles/cloudkms.admin"
  member     = "serviceAccount:service-${google_project.key_project.number}@gcp-sa-cloudkms.iam.gserviceaccount.com"
  depends_on = [time_sleep.wait_service_agent]
}

# Wait delay after granting IAM permissions
resource "time_sleep" "wait_srv_acc_permissions" {
  create_duration = "10s"
  depends_on      = [google_project_iam_member.autokey_project_admin]
}

resource "google_kms_autokey_config" "example-autokeyconfig" {
  provider    = google-beta
  folder      = google_folder.autokms_folder.id
  key_project = "projects/${google_project.key_project.project_id}"
  depends_on  = [time_sleep.wait_srv_acc_permissions]
}

# Wait delay after setting AutokeyConfig, to prevent diffs on reapply,
# because setting the config takes a little to fully propagate.
resource "time_sleep" "wait_autokey_propagation" {
  create_duration = "30s"
  depends_on      = [google_kms_autokey_config.example-autokeyconfig]
}
`, context)
}

func testAccCheckKMSAutokeyConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_kms_autokey_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{KMSBasePath}}folders/{{folder}}/autokeyConfig")
			url = strings.Replace(url, "folders/folders/", "folders/", 1)
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return nil
			}

			if v := res["key_project"]; v != nil {
				return fmt.Errorf("AutokeyConfig still exists at %s", url)
			}

			return nil
		}

		return nil
	}
}
