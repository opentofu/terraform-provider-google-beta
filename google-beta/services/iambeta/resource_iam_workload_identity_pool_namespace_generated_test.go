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

package iambeta_test

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

func TestAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceBasicExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  mode                      = "TRUST_DOMAIN"
}

resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "tf-test-example-namespace%{random_suffix}"
}
`, context)
}

func TestAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceFullExample(context),
			},
			{
				ResourceName:            "google_iam_workload_identity_pool_namespace.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"workload_identity_pool_id", "workload_identity_pool_namespace_id"},
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPoolNamespace_iamWorkloadIdentityPoolNamespaceFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_workload_identity_pool" "pool" {
  provider = google-beta

  workload_identity_pool_id = "tf-test-example-pool%{random_suffix}"
  mode                      = "TRUST_DOMAIN"
}

resource "google_iam_workload_identity_pool_namespace" "example" {
  provider = google-beta

  workload_identity_pool_id           = google_iam_workload_identity_pool.pool.workload_identity_pool_id
  workload_identity_pool_namespace_id = "tf-test-example-namespace%{random_suffix}"
  description                         = "Example Namespace in a Workload Identity Pool"
  disabled                            = true
}
`, context)
}

func testAccCheckIAMBetaWorkloadIdentityPoolNamespaceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_workload_identity_pool_namespace" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IAMBetaBasePath}}projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/namespaces/{{workload_identity_pool_namespace_id}}")
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

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("IAMBetaWorkloadIdentityPoolNamespace still exists at %s", url)
		}

		return nil
	}
}
