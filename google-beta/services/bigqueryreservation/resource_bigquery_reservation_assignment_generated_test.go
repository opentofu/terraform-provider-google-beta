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

package bigqueryreservation_test

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

func TestAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryReservationReservationAssignmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_reservation_assignment.assignment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "reservation"},
			},
		},
	})
}

func testAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_reservation" "basic" {
  name  = "tf-test-example-reservation%{random_suffix}"
  project = "%{project}"
  location = "us-central1"
  slot_capacity = 0
  ignore_idle_slots = false
}

resource "google_bigquery_reservation_assignment" "assignment" {
  assignee  = "projects/%{project}"
  job_type = "PIPELINE"
  reservation = google_bigquery_reservation.basic.id
}
`, context)
}

func TestAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigqueryReservationReservationAssignmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentFullExample(context),
			},
			{
				ResourceName:            "google_bigquery_reservation_assignment.assignment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "reservation"},
			},
		},
	})
}

func testAccBigqueryReservationReservationAssignment_bigqueryReservationAssignmentFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_reservation" "basic" {
  name  = "tf-test-example-reservation%{random_suffix}"
  project = "%{project}"
  location = "us-central1"
  slot_capacity = 0
  ignore_idle_slots = false
}

resource "google_bigquery_reservation_assignment" "assignment" {
  assignee  = "projects/%{project}"
  job_type = "QUERY"
  location = "us-central1"
  reservation = google_bigquery_reservation.basic.id
}
`, context)
}

func testAccCheckBigqueryReservationReservationAssignmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_reservation_assignment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments")
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
				return fmt.Errorf("BigqueryReservationReservationAssignment still exists at %s", url)
			}
		}

		return nil
	}
}
