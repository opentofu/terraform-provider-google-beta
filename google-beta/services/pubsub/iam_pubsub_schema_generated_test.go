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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/pubsub/Schema.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccPubsubSchemaIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSchemaIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_schema_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/schemas/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-schema%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccPubsubSchemaIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_schema_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/schemas/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-schema%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubSchemaIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccPubsubSchemaIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_pubsub_schema_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/schemas/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-schema%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPubsubSchemaIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPubsubSchemaIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_pubsub_schema_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_pubsub_schema_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/schemas/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-schema%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPubsubSchemaIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_pubsub_schema_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/schemas/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("tf-test-example-schema%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPubsubSchemaIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "tf-test-example-schema%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

resource "google_pubsub_schema_iam_member" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccPubsubSchemaIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "tf-test-example-schema%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_pubsub_schema_iam_policy" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_pubsub_schema_iam_policy" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  depends_on = [
    google_pubsub_schema_iam_policy.foo
  ]
}
`, context)
}

func testAccPubsubSchemaIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "tf-test-example-schema%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

data "google_iam_policy" "foo" {
}

resource "google_pubsub_schema_iam_policy" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccPubsubSchemaIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "tf-test-example-schema%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

resource "google_pubsub_schema_iam_binding" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccPubsubSchemaIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_pubsub_schema" "example" {
  name = "tf-test-example-schema%{random_suffix}"
  type = "AVRO"
  definition = "{\n  \"type\" : \"record\",\n  \"name\" : \"Avro\",\n  \"fields\" : [\n    {\n      \"name\" : \"StringField\",\n      \"type\" : \"string\"\n    },\n    {\n      \"name\" : \"IntField\",\n      \"type\" : \"int\"\n    }\n  ]\n}\n"
}

resource "google_pubsub_schema_iam_binding" "foo" {
  project = google_pubsub_schema.example.project
  schema = google_pubsub_schema.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
