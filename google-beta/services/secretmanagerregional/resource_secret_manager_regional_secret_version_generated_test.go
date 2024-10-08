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

package secretmanagerregional_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionBasicExample(context),
			},
			{
				ResourceName:            "google_secret_manager_regional_secret_version.regional_secret_version_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "secret"},
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  location = "us-central1"
}

resource "google_secret_manager_regional_secret_version" "regional_secret_version_basic" {
  secret = google_secret_manager_regional_secret.secret-basic.id
  secret_data = "tf-test-secret-data%{random_suffix}"
}
`, context)
}

func TestAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionWithBase64DataExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"data":          "./test-fixtures/binary-file.pfx",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionWithBase64DataExample(context),
			},
			{
				ResourceName:            "google_secret_manager_regional_secret_version.regional_secret_version_base64",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"is_secret_data_base64", "location", "secret"},
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionWithBase64DataExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  location = "us-central1"
}

resource "google_secret_manager_regional_secret_version" "regional_secret_version_base64" {
  secret = google_secret_manager_regional_secret.secret-basic.id
  secret_data = filebase64("%{data}")
  is_secret_data_base64 = true
}
`, context)
}

func TestAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDisabledExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDisabledExample(context),
			},
			{
				ResourceName:            "google_secret_manager_regional_secret_version.regional_secret_version_disabled",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "secret"},
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDisabledExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  location = "us-central1"
}

resource "google_secret_manager_regional_secret_version" "regional_secret_version_disabled" {
  secret = google_secret_manager_regional_secret.secret-basic.id
  secret_data = "tf-test-secret-data%{random_suffix}"
  enabled = false
}
`, context)
}

func TestAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyAbandonExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyAbandonExample(context),
			},
			{
				ResourceName:            "google_secret_manager_regional_secret_version.regional_secret_version_deletion_policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "location", "secret"},
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyAbandonExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  location = "us-central1"
}

resource "google_secret_manager_regional_secret_version" "regional_secret_version_deletion_policy" {
  secret = google_secret_manager_regional_secret.secret-basic.id
  secret_data = "tf-test-secret-data%{random_suffix}"
  deletion_policy = "ABANDON"
}
`, context)
}

func TestAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyDisableExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyDisableExample(context),
			},
			{
				ResourceName:            "google_secret_manager_regional_secret_version.regional_secret_version_deletion_policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "location", "secret"},
			},
		},
	})
}

func testAccSecretManagerRegionalRegionalSecretVersion_regionalSecretVersionDeletionPolicyDisableExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_regional_secret" "secret-basic" {
  secret_id = "tf-test-secret-version%{random_suffix}"
  location = "us-central1"
}

resource "google_secret_manager_regional_secret_version" "regional_secret_version_deletion_policy" {
  secret = google_secret_manager_regional_secret.secret-basic.id
  secret_data = "tf-test-secret-data%{random_suffix}"
  deletion_policy = "DISABLE"
}
`, context)
}

func testAccCheckSecretManagerRegionalRegionalSecretVersionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secret_manager_regional_secret_version" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecretManagerRegionalBasePath}}{{name}}")
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
				return fmt.Errorf("SecretManagerRegionalRegionalSecretVersion still exists at %s", url)
			}
		}

		return nil
	}
}
