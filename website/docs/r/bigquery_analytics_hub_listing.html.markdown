---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryanalyticshub/Listing.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "BigQuery Analytics Hub"
description: |-
  A Bigquery Analytics Hub data exchange listing
---

# google_bigquery_analytics_hub_listing

A Bigquery Analytics Hub data exchange listing


To get more information about Listing, see:

* [API documentation](https://cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.dataExchanges.listings)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/bigquery/docs/analytics-hub-introduction)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Basic


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "my_data_exchange"
  display_name     = "my_data_exchange"
  description      = "example data exchange"
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "my_listing"
  display_name     = "my_listing"
  description      = "example data exchange"

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
  }
}

resource "google_bigquery_dataset" "listing" {
  dataset_id                  = "my_listing"
  friendly_name               = "my_listing"
  description                 = "example data exchange"
  location                    = "US"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_restricted&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Restricted


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "my_data_exchange"
  display_name     = "my_data_exchange"
  description      = "example data exchange"
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "my_listing"
  display_name     = "my_listing"
  description      = "example data exchange"

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
  }

  restricted_export_config {
    enabled               = true
    restrict_query_result = true
  }
}

resource "google_bigquery_dataset" "listing" {
  dataset_id                  = "my_listing"
  friendly_name               = "my_listing"
  description                 = "example data exchange"
  location                    = "US"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_dcr&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Dcr


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "dcr_data_exchange"
  display_name     = "dcr_data_exchange"
  description      = "example dcr data exchange"
  sharing_environment_config  {
    dcr_exchange_config {}
  }
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "dcr_listing"
  display_name     = "dcr_listing"
  description      = "example dcr data exchange"

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
    selected_resources {
        table = google_bigquery_table.listing.id
    }
  }

  restricted_export_config {
    enabled                   = true
  }
}

resource "google_bigquery_dataset" "listing" {
  dataset_id                  = "dcr_listing"
  friendly_name               = "dcr_listing"
  description                 = "example dcr data exchange"
  location                    = "US"
}

resource "google_bigquery_table" "listing" {
  deletion_protection = false
  table_id   = "dcr_listing"
  dataset_id = google_bigquery_dataset.listing.dataset_id
  schema = <<EOF
[
  {
    "name": "name",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "post_abbr",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "date",
    "type": "DATE",
    "mode": "NULLABLE"
  }
]
EOF
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_log_linked_dataset_query_user&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Log Linked Dataset Query User


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing_log_email" {
  location         = "US"
  data_exchange_id = "tf_test_log_email_de" 
  display_name     = "tf_test_log_email_de" 
  description      = "Example for log email test"
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing_log_email.data_exchange_id
  listing_id       = "tf_test_log_email_listing" 
  display_name     = "tf_test_log_email_listing" 
  description      = "Example for log email test"
  log_linked_dataset_query_user_email = true

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing_log_email.id
  }
}

resource "google_bigquery_dataset" "listing_log_email" {
  dataset_id                  = "tf_test_log_email_ds" 
  friendly_name               = "tf_test_log_email_ds" 
  description                 = "Example for log email test"
  location                    = "US"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_pubsub&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Pubsub


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "tf_test_pubsub_data_exchange"
  display_name     = "tf_test_pubsub_data_exchange"
  description      = "Example for pubsub topic source"
}

resource "google_pubsub_topic" "tf_test_pubsub_topic" { 
  name    = "test_pubsub" 
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "tf_test_pubsub_listing"
  display_name     = "tf_test_pubsub_listing"
  description      = "Example for pubsub topic source"

  pubsub_topic {
    topic = google_pubsub_topic.tf_test_pubsub_topic.id
    data_affinity_regions = [
      "us-central1",
      "europe-west1"
    ]
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_dcr_routine&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Dcr Routine


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "dcr_data_exchange_example" {
  provider = google-beta
  location         = "us"
  data_exchange_id = "tf_test_data_exchange"
  display_name     = "tf_test_data_exchange"
  description      = "Example for listing with routine"
  sharing_environment_config {
    dcr_exchange_config {}
  }
}

resource "google_bigquery_dataset" "listing" {
  provider = google-beta
  dataset_id    = "tf_test_dataset"
  friendly_name = "tf_test_dataset"
  description   = "Example for listing with routine"
  location      = "us"
}

resource "google_bigquery_routine" "listing" {
  provider = google-beta
  dataset_id      = google_bigquery_dataset.listing.dataset_id
  routine_id      = "tf_test_routine"
  routine_type    = "TABLE_VALUED_FUNCTION"
  language        = "SQL"
  description     = "A DCR routine example."
  definition_body = <<-EOS
    SELECT 1 + value AS value
  EOS
  arguments {
    name          = "value"
    argument_kind = "FIXED_TYPE"
    data_type     = jsonencode({ "typeKind" : "INT64" })
  }
  return_table_type = jsonencode({
    "columns" : [
      { "name" : "value", "type" : { "typeKind" : "INT64" } },
    ]
  })
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  provider = google-beta
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.dcr_data_exchange_example.data_exchange_id
  listing_id       = "tf_test_listing_routine"
  display_name     = "tf_test_listing_routine"
  description      = "Example for listing with routine"
  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
    selected_resources {
      routine = google_bigquery_routine.listing.id
    }
  }
  restricted_export_config {
    enabled = true
  }
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_public_listing&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Public Listing


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "my_data_exchange"
  display_name     = "my_data_exchange"
  description      = "example public listing"
  discovery_type   = "DISCOVERY_TYPE_PUBLIC"
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "my_listing"
  display_name     = "my_listing"
  description      = "example public listing"
  discovery_type   = "DISCOVERY_TYPE_PUBLIC"
  allow_only_metadata_sharing= false

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
  }
}

resource "google_bigquery_dataset" "listing" {
  dataset_id                  = "my_listing"
  friendly_name               = "my_listing"
  description                 = "example public listing"
  location                    = "US"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_listing_marketplace&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Listing Marketplace


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "listing" {
  location         = "US"
  data_exchange_id = "my_data_exchange"
  display_name     = "my_data_exchange"
  description      = "example data exchange"
}

resource "google_bigquery_analytics_hub_listing" "listing" {
  location         = "US"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.listing.data_exchange_id
  listing_id       = "my_listing"
  display_name     = "my_listing"
  description      = "example data exchange"
  delete_commercial = true

  bigquery_dataset {
    dataset = google_bigquery_dataset.listing.id
  }

}

resource "google_bigquery_dataset" "listing" {
  dataset_id                  = "my_listing"
  friendly_name               = "my_listing"
  description                 = "example data exchange"
  location                    = "US"
}
```

## Argument Reference

The following arguments are supported:


* `data_exchange_id` -
  (Required)
  The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.

* `listing_id` -
  (Required)
  The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.

* `location` -
  (Required)
  The name of the location this data exchange listing.

* `display_name` -
  (Required)
  Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.


* `description` -
  (Optional)
  Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).

* `primary_contact` -
  (Optional)
  Email or URL of the primary point of contact of the listing.

* `documentation` -
  (Optional)
  Documentation describing the listing.

* `icon` -
  (Optional)
  Base64 encoded image representing the listing.

* `request_access` -
  (Optional)
  Email or URL of the request access of the listing. Subscribers can use this reference to request access.

* `data_provider` -
  (Optional)
  Details of the data provider who owns the source data.
  Structure is [documented below](#nested_data_provider).

* `publisher` -
  (Optional)
  Details of the publisher who owns the listing and who can share the source data.
  Structure is [documented below](#nested_publisher).

* `categories` -
  (Optional)
  Categories of the listing. Up to two categories are allowed.

* `bigquery_dataset` -
  (Optional)
  Shared dataset i.e. BigQuery dataset source.
  Structure is [documented below](#nested_bigquery_dataset).

* `pubsub_topic` -
  (Optional)
  Pub/Sub topic source.
  Structure is [documented below](#nested_pubsub_topic).

* `restricted_export_config` -
  (Optional)
  If set, restricted export configuration will be propagated and enforced on the linked dataset.
  Structure is [documented below](#nested_restricted_export_config).

* `log_linked_dataset_query_user_email` -
  (Optional)
  If true, subscriber email logging is enabled and all queries on the linked dataset will log the email address of the querying user. Once enabled, this setting cannot be turned off.

* `discovery_type` -
  (Optional)
  Specifies the type of discovery on the discovery page. Cannot be set for a restricted listing. Note that this does not control the visibility of the exchange/listing which is defined by IAM permission.
  Possible values are: `DISCOVERY_TYPE_PRIVATE`, `DISCOVERY_TYPE_PUBLIC`.

* `allow_only_metadata_sharing` -
  (Optional)
  If true, the listing is only available to get the resource metadata. Listing is non subscribable.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `delete_commercial` - (Optional) If the listing is commercial then this field must be set to true, otherwise a failure is thrown. This acts as a safety guard to avoid deleting commercial listings accidentally.


<a name="nested_data_provider"></a>The `data_provider` block supports:

* `name` -
  (Required)
  Name of the data provider.

* `primary_contact` -
  (Optional)
  Email or URL of the data provider.

<a name="nested_publisher"></a>The `publisher` block supports:

* `name` -
  (Required)
  Name of the listing publisher.

* `primary_contact` -
  (Optional)
  Email or URL of the listing publisher.

<a name="nested_bigquery_dataset"></a>The `bigquery_dataset` block supports:

* `dataset` -
  (Required)
  Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123

* `selected_resources` -
  (Optional)
  Resource in this dataset that is selectively shared. This field is required for data clean room exchanges.
  Structure is [documented below](#nested_bigquery_dataset_selected_resources).


<a name="nested_bigquery_dataset_selected_resources"></a>The `selected_resources` block supports:

* `table` -
  (Optional)
  Format: For table: projects/{projectId}/datasets/{datasetId}/tables/{tableId} Example:"projects/test_project/datasets/test_dataset/tables/test_table"

* `routine` -
  (Optional, [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html))
  Format: For routine: projects/{projectId}/datasets/{datasetId}/routines/{routineId} Example:"projects/test_project/datasets/test_dataset/routines/test_routine"

<a name="nested_pubsub_topic"></a>The `pubsub_topic` block supports:

* `topic` -
  (Required)
  Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId

* `data_affinity_regions` -
  (Optional)
  Region hint on where the data might be published. Data affinity regions are modifiable.
  See https://cloud.google.com/about/locations for full listing of possible Cloud regions.

<a name="nested_restricted_export_config"></a>The `restricted_export_config` block supports:

* `enabled` -
  (Optional)
  If true, enable restricted export.

* `restrict_direct_table_access` -
  (Output)
  If true, restrict direct table access(read api/tabledata.list) on linked table.

* `restrict_query_result` -
  (Optional)
  If true, restrict export of query result derived from restricted linked dataset table.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}`

* `name` -
  The resource name of the listing. e.g. "projects/myproject/locations/US/dataExchanges/123/listings/456"

* `state` -
  Current state of the listing.

* `commercial_info` -
  Commercial info contains the information about the commercial data products associated with the listing.
  Structure is [documented below](#nested_commercial_info).


<a name="nested_commercial_info"></a>The `commercial_info` block contains:

* `cloud_marketplace` -
  (Output)
  Details of the Marketplace Data Product associated with the Listing.
  Structure is [documented below](#nested_commercial_info_cloud_marketplace).


<a name="nested_commercial_info_cloud_marketplace"></a>The `cloud_marketplace` block contains:

* `service` -
  (Output)
  Resource name of the commercial service associated with the Marketplace Data Product. e.g. example.com

* `commercial_state` -
  (Output)
  Commercial state of the Marketplace Data Product.
  Possible values: COMMERCIAL_STATE_UNSPECIFIED, ONBOARDING, ACTIVE

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Listing can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}`
* `{{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}`
* `{{location}}/{{data_exchange_id}}/{{listing_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Listing using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}"
  to = google_bigquery_analytics_hub_listing.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), Listing can be imported using one of the formats above. For example:

```
$ terraform import google_bigquery_analytics_hub_listing.default projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
$ terraform import google_bigquery_analytics_hub_listing.default {{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}
$ terraform import google_bigquery_analytics_hub_listing.default {{location}}/{{data_exchange_id}}/{{listing_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
