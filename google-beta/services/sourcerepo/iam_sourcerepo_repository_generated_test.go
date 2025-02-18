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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/sourcerepo/Repository.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/examples/base_configs/iam_test_file.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package sourcerepo_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccSourceRepoRepositoryIamBindingGenerated(t *testing.T) {
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
				Config: testAccSourceRepoRepositoryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSourceRepoRepositoryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer", envvar.GetTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamMemberGenerated(t *testing.T) {
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
				Config: testAccSourceRepoRepositoryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamPolicyGenerated(t *testing.T) {
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
				Config: testAccSourceRepoRepositoryIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_sourcerepo_repository_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSourceRepoRepositoryIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s", envvar.GetTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSourceRepoRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_member" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSourceRepoRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_sourcerepo_repository_iam_policy" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_sourcerepo_repository_iam_policy" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  depends_on = [
    google_sourcerepo_repository_iam_policy.foo
  ]
}
`, context)
}

func testAccSourceRepoRepositoryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

data "google_iam_policy" "foo" {
}

resource "google_sourcerepo_repository_iam_policy" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
  project = google_sourcerepo_repository.my-repo.project
  repository = google_sourcerepo_repository.my-repo.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
