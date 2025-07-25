---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryanalyticshub/DataExchangeSubscription.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "BigQuery Analytics Hub"
description: |-
  A Bigquery Analytics Hub Data Exchange subscription
---

# google_bigquery_analytics_hub_data_exchange_subscription

A Bigquery Analytics Hub Data Exchange subscription

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about DataExchangeSubscription, see:

* [API documentation](https://cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.subscriptions)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/bigquery/docs/analytics-hub-introduction)

~> **Note:** When importing the resource with `terraform import`, provide the destination/subscriber's project and location
in the format projects/{{subscriber_project}}/locations/{{subscriber_location}}/subscriptions/{{subscription_id}}
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md&cloudshell_working_dir=bigquery_analyticshub_dataexchange_subscription_basic&open_in_editor=main.tf" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Bigquery Analyticshub Dataexchange Subscription Basic


```hcl
resource "google_bigquery_analytics_hub_data_exchange" "subscription" {
  provider = google-beta
  location            = "us"
  data_exchange_id    = "my_test_dataexchange"
  display_name        = "my_test_dataexchange"
  description         = "Test Data Exchange"
  sharing_environment_config {
    dcr_exchange_config {}
  }
}

resource "google_bigquery_dataset" "subscription" {
  provider = google-beta
  dataset_id    = "listing_src_dataset"
  friendly_name = "listing_src_dataset"
  description   = "Dataset for Listing"
  location      = "us"
}

resource "google_bigquery_table" "subscription" {
  provider = google-beta
  deletion_protection = false
  table_id            = "listing_src_table"
  dataset_id          = google_bigquery_dataset.subscription.dataset_id
  schema              = <<EOF
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

resource "google_bigquery_analytics_hub_listing" "subscription" {
  provider = google-beta
  location             = "us"
  data_exchange_id     = google_bigquery_analytics_hub_data_exchange.subscription.data_exchange_id
  listing_id           = "my_test_listing"
  display_name         = "my_test_listing"
  description          = "Test Listing"

  restricted_export_config {
    enabled = true
  }

  bigquery_dataset {
    dataset = google_bigquery_dataset.subscription.id
    selected_resources {
      table = google_bigquery_table.subscription.id
    }
  }
}

resource "google_bigquery_analytics_hub_data_exchange_subscription" "subscription" {
  provider = google-beta
  project                = google_bigquery_dataset.subscription.project #Subscriber's project
  location               = "us"

  data_exchange_project  = google_bigquery_analytics_hub_data_exchange.subscription.project
  data_exchange_location = google_bigquery_analytics_hub_data_exchange.subscription.location
  data_exchange_id       = google_bigquery_analytics_hub_data_exchange.subscription.data_exchange_id

  subscription_id    = "my_subscription_id"
  subscriber_contact = "testuser@example.com"

  destination_dataset {
    location = "us"

    dataset_reference {
      project_id = google_bigquery_dataset.subscription.project #Subscriber's project
      dataset_id = "subscribed_dest_dataset"
    }
    friendly_name = "Subscribed Destination Dataset"
    description   = "Destination dataset for subscription"
    labels = {
      environment = "development"
      owner       = "team-a"
    }
  }

  refresh_policy="ON_READ"
}
```

## Argument Reference

The following arguments are supported:


* `data_exchange_id` -
  (Required)
  The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.

* `data_exchange_project` -
  (Required)
  The ID of the Google Cloud project where the Data Exchange is located.

* `data_exchange_location` -
  (Required)
  The name of the location of the Data Exchange.

* `location` -
  (Required)
  The geographic location where the Subscription (and its linked dataset) should reside.
  This is the subscriber's desired location for the created resources.
  See https://cloud.google.com/bigquery/docs/locations for supported locations.

* `subscription_id` -
  (Required)
  Name of the subscription to create.


* `subscriber_contact` -
  (Optional)
  Email of the subscriber.

* `destination_dataset` -
  (Optional)
  BigQuery destination dataset to create for the subscriber.
  Structure is [documented below](#nested_destination_dataset).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `refresh_policy` - (Optional) Controls when the subscription is automatically refreshed by the provider.
* `ON_READ`: Default value if not specified. The subscription will be refreshed every time Terraform performs a read operation (e.g., `terraform plan`, `terraform apply`, `terraform refresh`). This ensures the state is always up-to-date.
* `ON_STALE`: The subscription will only be refreshed when its reported `state` (an output-only field from the API) is `STATE_STALE` during a Terraform read operation.
* `NEVER`: The provider will not automatically refresh the subscription.


<a name="nested_destination_dataset"></a>The `destination_dataset` block supports:

* `location` -
  (Required)
  The geographic location where the dataset should reside.
  See https://cloud.google.com/bigquery/docs/locations for supported locations.

* `dataset_reference` -
  (Required)
  A reference that identifies the destination dataset.
  Structure is [documented below](#nested_destination_dataset_dataset_reference).

* `friendly_name` -
  (Optional)
  A descriptive name for the dataset.

* `description` -
  (Optional)
  A user-friendly description of the dataset.

* `labels` -
  (Optional)
  The labels associated with this dataset. You can use these to
  organize and group your datasets.


<a name="nested_destination_dataset_dataset_reference"></a>The `dataset_reference` block supports:

* `dataset_id` -
  (Required)
  A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters.

* `project_id` -
  (Required)
  The ID of the project containing this dataset.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}`

* `name` -
  The resource name of the subscription. e.g. "projects/myproject/locations/us/subscriptions/123"

* `creation_time` -
  Timestamp when the subscription was created.

* `last_modify_time` -
  Timestamp when the subscription was last modified.

* `organization_id` -
  Organization of the project this subscription belongs to.

* `organization_display_name` -
  Display name of the project of this subscription.

* `state` -
  Current state of the subscription.

* `resource_type` -
  Listing shared asset type.

* `linked_dataset_map` -
  Output only. Map of listing resource names to associated linked resource,
  e.g. projects/123/locations/us/dataExchanges/456/listings/789 -> projects/123/datasets/my_dataset
  For Data Exchange subscriptions, this map may contain multiple entries if the Data Exchange has multiple listings.
  Structure is [documented below](#nested_linked_dataset_map).

* `linked_resources` -
  Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.
  Structure is [documented below](#nested_linked_resources).

* `data_exchange` -
  Output only. Resource name of the source Data Exchange. e.g. projects/123/locations/us/dataExchanges/456

* `log_linked_dataset_query_user_email` -
  Output only. By default, false. If true, the Subscriber agreed to the email sharing mandate that is enabled for DataExchange/Listing.


<a name="nested_linked_dataset_map"></a>The `linked_dataset_map` block contains:

* `resource_name` - (Required) The identifier for this object. Format specified above.

* `listing` -
  (Output)
  Output only. Listing for which linked resource is created.

* `linked_dataset` -
  (Output)
  Output only. Name of the linked dataset, e.g. projects/subscriberproject/datasets/linkedDataset

* `linked_pubsub_subscription` -
  (Output)
  Output only. Name of the Pub/Sub subscription, e.g. projects/subscriberproject/subscriptions/subscriptions/sub_id

<a name="nested_linked_resources"></a>The `linked_resources` block contains:

* `listing` -
  (Output)
  Output only. Listing for which linked resource is created.

* `linked_dataset` -
  (Output)
  Output only. Name of the linked dataset, e.g. projects/subscriberproject/datasets/linkedDataset

## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


DataExchangeSubscription can be imported using any of these accepted formats:

* `projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}`
* `{{project}}/{{location}}/{{subscription_id}}`
* `{{location}}/{{subscription_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DataExchangeSubscription using one of the formats above. For example:

```tf
import {
  id = "projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}"
  to = google_bigquery_analytics_hub_data_exchange_subscription.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), DataExchangeSubscription can be imported using one of the formats above. For example:

```
$ terraform import google_bigquery_analytics_hub_data_exchange_subscription.default projects/{{project}}/locations/{{location}}/subscriptions/{{subscription_id}}
$ terraform import google_bigquery_analytics_hub_data_exchange_subscription.default {{project}}/{{location}}/{{subscription_id}}
$ terraform import google_bigquery_analytics_hub_data_exchange_subscription.default {{location}}/{{subscription_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
