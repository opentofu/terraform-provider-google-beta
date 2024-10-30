// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwprovider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

// TestAccFwProvider_billing_project is a series of acc tests asserting how the PF provider handles billing_project arguments
// It is PF specific because the HCL used provisions PF-implemented resources
// It is a counterpart to TestAccSdkProvider_billing_project
func TestAccFwProvider_billing_project(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		// Configuring the provider using inputs
		"config takes precedence over environment variables":                                           testAccFwProvider_billing_project_configPrecedenceOverEnvironmentVariables,
		"when billing_project is unset in the config, environment variables are used in a given order": testAccFwProvider_billing_project_precedenceOrderEnvironmentVariables, // GOOGLE_BILLING_PROJECT

		// Schema-level validation
		"when billing_project is set to an empty string in the config the value isn't ignored and results in an error": testAccFwProvider_billing_project_emptyStringValidation,

		// Usage
		// TODO: https://github.com/hashicorp/terraform-provider-google/issues/17882
		"GOOGLE_CLOUD_QUOTA_PROJECT environment variable interferes with the billing_account value used": testAccFwProvider_billing_project_affectedByClientLibraryEnv,
		// 1) Usage of billing_account alone is insufficient
		// 2) Usage in combination with user_project_override changes the project where quota is used
		"using billing_account alone doesn't impact provisioning, but using together with user_project_override does": testAccFwProvider_billing_project_useWithAndWithoutUserProjectOverride,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccFwProvider_billing_project_configPrecedenceOverEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	billingProject := "my-billing-project-id"

	// ensure all possible billing_project env vars set; show they aren't used instead
	t.Setenv("GOOGLE_BILLING_PROJECT", billingProject)

	providerBillingProject := "foobar"

	context := map[string]interface{}{
		"billing_project": providerBillingProject,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_billing_project_inProviderBlock(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "billing_project", providerBillingProject),
				)},
		},
	})
}

func testAccFwProvider_billing_project_precedenceOrderEnvironmentVariables(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API
	/*
		These are all the ENVs for billing_project
		GOOGLE_BILLING_PROJECT

		GOOGLE_CLOUD_QUOTA_PROJECT - NOT used by provider, but is in client libraries we use
	*/

	GOOGLE_BILLING_PROJECT := "GOOGLE_BILLING_PROJECT"
	GOOGLE_CLOUD_QUOTA_PROJECT := "GOOGLE_CLOUD_QUOTA_PROJECT"

	context := map[string]interface{}{}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// GOOGLE_BILLING_PROJECT is used if set
				PreConfig: func() {
					t.Setenv("GOOGLE_BILLING_PROJECT", GOOGLE_BILLING_PROJECT) //used
					t.Setenv("GOOGLE_CLOUD_QUOTA_PROJECT", GOOGLE_CLOUD_QUOTA_PROJECT)
				},
				Config: testAccFwProvider_billing_project_inEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "billing_project", GOOGLE_BILLING_PROJECT),
				),
			},
			{
				// GOOGLE_CLOUD_QUOTA_PROJECT is NOT used here
				PreConfig: func() {
					t.Setenv("GOOGLE_BILLING_PROJECT", "")
					t.Setenv("GOOGLE_CLOUD_QUOTA_PROJECT", GOOGLE_CLOUD_QUOTA_PROJECT) // NOT used
				},
				Config: testAccFwProvider_billing_project_inEnvsOnly(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.google_provider_config_plugin_framework.default", "billing_project"),
				),
			},
		},
	})
}

func testAccFwProvider_billing_project_emptyStringValidation(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	billingProject := "my-billing-project-id"

	// ensure all billing_project env vars set
	t.Setenv("GOOGLE_BILLING_PROJECT", billingProject)

	context := map[string]interface{}{
		"billing_project": "", // empty string used
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccFwProvider_billing_project_inProviderBlock(context),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("expected a non-empty string"),
			},
		},
	})
}
func testAccFwProvider_billing_project_useWithAndWithoutUserProjectOverride(t *testing.T) {
	acctest.SkipIfVcr(t) // VCR fails to record final interaction that fails with 403

	randomString := acctest.RandString(t, 10)
	contextUserProjectOverrideFalse := map[string]interface{}{
		"org_id":                envvar.GetTestOrgFromEnv(t),
		"billing_account":       envvar.GetTestBillingAccountFromEnv(t),
		"user_project_override": "false", // Used in combo with billing_account
		"random_suffix":         randomString,
	}

	contextUserProjectOverrideTrue := map[string]interface{}{
		"org_id":                envvar.GetTestOrgFromEnv(t),
		"billing_account":       envvar.GetTestBillingAccountFromEnv(t),
		"user_project_override": "true", // Used in combo with billing_account
		"random_suffix":         randomString,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Setup resources
				// Neither user_project_override nor billing_project value used here
				Config: testAccFwProvider_billing_project_useBillingProject_setup(contextUserProjectOverrideFalse),
			},
			{
				// With user_project_override=true the Firebase data source CANNOT be used because quota is consumed
				// from the newly provisioned project, and that project does not have the Firebase Management API enabled.
				// The billing_project is used, leading to the error occurring, because user_project_override=true
				Config:      testAccFwProvider_billing_project_useBillingProject_scenario(contextUserProjectOverrideTrue),
				ExpectError: regexp.MustCompile("Error 403: Firebase Management API has not been used in project"),
			},
			{
				// With user_project_override=false the Firebase data source still cannot be used, but only because we're querying a made up app!
				// This is necessary because we cannot provision anything in the newly provisioned project (to be queried with the datasource)
				// unless we make it a Firebase-enabled project. This enables the API, and that undermines the premise of this test.
				Config:      testAccFwProvider_billing_project_useBillingProject_scenario(contextUserProjectOverrideFalse),
				ExpectError: regexp.MustCompile("Error 403: The caller does not have permission, forbidden"),
			},
		},
	})
}

func testAccFwProvider_billing_project_affectedByClientLibraryEnv(t *testing.T) {
	acctest.SkipIfVcr(t) // VCR fails to record final interaction that fails with 403

	randomString := acctest.RandString(t, 10)
	context := map[string]interface{}{
		"org_id":                envvar.GetTestOrgFromEnv(t),
		"billing_account":       envvar.GetTestBillingAccountFromEnv(t),
		"user_project_override": "true", // Used in combo with billing_account
		"random_suffix":         randomString,
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				// Setup resources
				// Neither user_project_override nor billing_project value used here
				Config: testAccFwProvider_billing_project_useBillingProject_setup(context),
			},
			{
				// This ENV interferes with setting the billing_project,
				// so we get an error mentioning the value
				PreConfig: func() {
					t.Setenv("GOOGLE_CLOUD_QUOTA_PROJECT", "foobar")
				},
				Config:      testAccFwProvider_billing_project_useBillingProject_scenario(context),
				ExpectError: regexp.MustCompile("foobar"),
			},
			{
				// The same config without that ENV present applies without error
				PreConfig: func() {
					t.Setenv("GOOGLE_CLOUD_QUOTA_PROJECT", "")
				},
				Config: testAccFwProvider_billing_project_useBillingProject_scenario(context),
				// We know we're using the billing_account project when we hit this error
				ExpectError: regexp.MustCompile("Error 403: Firebase Management API has not been used in project"),
			},
		},
	})
}

// testAccFwProvider_billing_project_inProviderBlock allows setting the billing_project argument in a provider block.
// This function uses data.google_provider_config_plugin_framework because it is implemented with the plugin-framework
func testAccFwProvider_billing_project_inProviderBlock(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	billing_project = "%{billing_project}"
}

data "google_provider_config_plugin_framework" "default" {}
`, context)
}

// testAccFwProvider_billing_project_inEnvsOnly allows testing when the billing_project argument
// is only supplied via ENVs
func testAccFwProvider_billing_project_inEnvsOnly(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_provider_config_plugin_framework" "default" {}
`, context)
}

func testAccFwProvider_billing_project_useBillingProject_setup(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google-beta" {}

# Create a new project and enable service APIs in those projects
resource "google_project" "project" {
  provider = google-beta
  project_id      = "tf-test-%{random_suffix}"
  name            = "tf-test-%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}
`, context)
}

func testAccFwProvider_billing_project_useBillingProject_scenario(context map[string]interface{}) string {

	// SECOND APPLY
	// This is needed as configuring the provider depends on resources provisioned in the setup step
	return testAccFwProvider_billing_project_useBillingProject_setup(context) + acctest.Nprintf(`
# Set up the usage of
#  - user_project_override
#  - billing_project
provider "google-beta" {
  alias                 = "user_project_override"
  user_project_override = %{user_project_override}
  billing_project       = google_project.project.project_id
  project               = google_project.project.project_id
}

# See if the impersonated SA can interact with the Firebase API in a way that uses
# the newly provisioned project as the source of consumed quota
# NOTE: This data source is implemented with plugin-framework so tests our use of user_project_override + billing_project in a PF specific way.
data "google_firebase_apple_app_config" "my_app_config" {
  provider = google-beta.user_project_override
  app_id = "1:234567891011:ios:1234abcdef1234abcde123" // Made up
}
`, context)
}