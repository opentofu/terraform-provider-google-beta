---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/website/docs/d/storage_insights_dataset_config.html.markdown
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Storage Insights"
description: |-
  Represents a Storage Insights DatasetConfig.
---

# google_storage_insights_dataset_config

Use this data source to get information about a Storage Insights Dataset Config resource.
See [the official documentation](https://cloud.google.com/storage/docs/insights/datasets)
and
[API](https://cloud.google.com/storage/docs/insights/reference/rest/v1/projects.locations.datasetConfigs).


## Example Usage

```hcl
data "google_storage_insights_dataset_config" "sample-config" {
  project = "sample_project"
  location = "sample_location"
  dataset_config_id = "sample_dataset_config_id"
}
```

## Argument Reference

The following arguments are supported:

* `project` - (Optional) The name of the GCP project in which dataset config exists. Can be configured through config as well.
* `location` - (Required) The location of the Dataset Config.
* `dataset_config_id` - (Required) The user-defined ID of the DatasetConfig


## Attributes Reference

See [google_storage_insights_dataset_config](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_insights_dataset_config#argument-reference) resource for details of the available attributes.
