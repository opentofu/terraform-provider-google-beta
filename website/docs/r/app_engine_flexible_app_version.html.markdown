---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This code is generated by Magic Modules using the following:
#
#     Configuration: https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/appengine/FlexibleAppVersion.yaml
#     Template:      https:#github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.html.markdown.tmpl
#
#     DO NOT EDIT this file directly. Any changes made to this file will be
#     overwritten during the next generation cycle.
#
# ----------------------------------------------------------------------------
subcategory: "App Engine"
description: |-
  Flexible App Version resource to create a new version of flexible GAE Application.
---

# google_app_engine_flexible_app_version

Flexible App Version resource to create a new version of flexible GAE Application. Based on Google Compute Engine,
the App Engine flexible environment automatically scales your app up and down while also balancing the load.
Learn about the differences between the standard environment and the flexible environment
at https://cloud.google.com/appengine/docs/the-appengine-environments.

~> **Note:** The App Engine flexible environment service account uses the member ID `service-[YOUR_PROJECT_NUMBER]@gae-api-prod.google.com.iam.gserviceaccount.com`
It should have the App Engine Flexible Environment Service Agent role, which will be applied when the `appengineflex.googleapis.com` service is enabled.


To get more information about FlexibleAppVersion, see:

* [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/appengine/docs/flexible)

## Example Usage - App Engine Flexible App Version


```hcl
resource "google_project" "my_project" {
  name = "appeng-flex"
  project_id = "appeng-flex"
  org_id = "123456789"
  billing_account = "000000-0000000-0000000-000000"
  deletion_policy = "DELETE"
}

resource "google_app_engine_application" "app" {
  project     = google_project.my_project.project_id
  location_id = "us-central"
}

resource "google_project_service" "service" {
  project = google_project.my_project.project_id
  service = "appengineflex.googleapis.com"

  disable_dependent_services = false
}

resource "google_service_account" "custom_service_account" {
  project      = google_project_service.service.project
  account_id   = "my-account"
  display_name = "Custom Service Account"
}

resource "google_project_iam_member" "gae_api" {
  project = google_project_service.service.project
  role    = "roles/compute.networkUser"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = google_project_service.service.project
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_project_iam_member" "storage_viewer" {
  project = google_project_service.service.project
  role    = "roles/storage.objectViewer"
  member  = "serviceAccount:${google_service_account.custom_service_account.email}"
}

resource "google_app_engine_flexible_app_version" "myapp_v1" {
  version_id = "v1"
  project    = google_project_iam_member.gae_api.project
  service    = "default"
  runtime    = "nodejs"
  flexible_runtime_settings {
    operating_system = "ubuntu22"
    runtime_version = "20"
  }

  entrypoint {
    shell = "node ./app.js"
  }

  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }
  }

  liveness_check {
    path = "/"
  }

  readiness_check {
    path = "/"
  }

  env_variables = {
    port = "8080"
  }

  handlers {
    url_regex        = ".*\\/my-path\\/*"
    security_level   = "SECURE_ALWAYS"
    login            = "LOGIN_REQUIRED"
    auth_fail_action = "AUTH_FAIL_ACTION_REDIRECT"

    static_files {
      path = "my-other-path"
      upload_path_regex = ".*\\/my-path\\/*"
    }
  }

  automatic_scaling {
    cool_down_period = "120s"
    cpu_utilization {
      target_utilization = 0.5
    }
  }

  noop_on_destroy = true
  service_account = google_service_account.custom_service_account.email
}

resource "google_storage_bucket" "bucket" {
  project  = google_project.my_project.project_id
  name     = "appengine-static-content"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/hello-world.zip"
}
```

## Argument Reference

The following arguments are supported:


* `runtime` -
  (Required)
  Desired runtime. Example python27.

* `readiness_check` -
  (Required)
  Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
  Structure is [documented below](#nested_readiness_check).

* `liveness_check` -
  (Required)
  Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
  Structure is [documented below](#nested_liveness_check).

* `service` -
  (Required)
  AppEngine service resource. Can contain numbers, letters, and hyphens.


* `version_id` -
  (Optional)
  Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
  Reserved names,"default", "latest", and any name with the prefix "ah-".

* `inbound_services` -
  (Optional)
  A list of the types of messages that this application is able to receive.
  Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.

* `instance_class` -
  (Optional)
  Instance class that is used to run this version. Valid values are
  AutomaticScaling: F1, F2, F4, F4_1G
  ManualScaling: B1, B2, B4, B8, B4_1G
  Defaults to F1 for AutomaticScaling and B1 for ManualScaling.

* `network` -
  (Optional)
  Extra network settings
  Structure is [documented below](#nested_network).

* `resources` -
  (Optional)
  Machine resources for a version.
  Structure is [documented below](#nested_resources).

* `runtime_channel` -
  (Optional)
  The channel of the runtime to use. Only available for some runtimes.

* `flexible_runtime_settings` -
  (Optional)
  Runtime settings for App Engine flexible environment.
  Structure is [documented below](#nested_flexible_runtime_settings).

* `beta_settings` -
  (Optional)
  Metadata settings that are supplied to this version to enable beta runtime features.

* `serving_status` -
  (Optional)
  Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
  Default value is `SERVING`.
  Possible values are: `SERVING`, `STOPPED`.

* `runtime_api_version` -
  (Optional)
  The version of the API in the given runtime environment.
  Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
  Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.

* `handlers` -
  (Optional)
  An ordered list of URL-matching patterns that should be applied to incoming requests.
  The first matching URL handles the request and other request handlers are not attempted.
  Structure is [documented below](#nested_handlers).

* `runtime_main_executable_path` -
  (Optional)
  The path or name of the app's main executable.

* `service_account` -
  (Optional)
  The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as
  default if this field is neither provided in app.yaml file nor through CLI flag.

* `api_config` -
  (Optional)
  Serving configuration for Google Cloud Endpoints.
  Structure is [documented below](#nested_api_config).

* `env_variables` -
  (Optional)
  Environment variables available to the application.  As these are not returned in the API request, Terraform will not detect any changes made outside of the Terraform config.

* `default_expiration` -
  (Optional)
  Duration that static files should be cached by web proxies and browsers.
  Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.

* `nobuild_files_regex` -
  (Optional)
  Files that match this pattern will not be built into this version. Only applicable for Go runtimes.

* `deployment` -
  (Optional)
  Code and application artifacts that make up this version.
  Structure is [documented below](#nested_deployment).

* `endpoints_api_service` -
  (Optional)
  Code and application artifacts that make up this version.
  Structure is [documented below](#nested_endpoints_api_service).

* `entrypoint` -
  (Optional)
  The entrypoint for the application.
  Structure is [documented below](#nested_entrypoint).

* `vpc_access_connector` -
  (Optional)
  Enables VPC connectivity for standard apps.
  Structure is [documented below](#nested_vpc_access_connector).

* `automatic_scaling` -
  (Optional)
  Automatic scaling is based on request rate, response latencies, and other application metrics.
  Structure is [documented below](#nested_automatic_scaling).

* `manual_scaling` -
  (Optional)
  A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
  Structure is [documented below](#nested_manual_scaling).

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.

* `noop_on_destroy` - (Optional) If set to `true`, the application version will not be deleted.

* `delete_service_on_destroy` - (Optional) If set to `true`, the service will be deleted if it is the last version.



<a name="nested_readiness_check"></a>The `readiness_check` block supports:

* `path` -
  (Required)
  The request path.

* `host` -
  (Optional)
  Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com"

* `failure_threshold` -
  (Optional)
  Number of consecutive failed checks required before removing traffic. Default: 2.

* `success_threshold` -
  (Optional)
  Number of consecutive successful checks required before receiving traffic. Default: 2.

* `check_interval` -
  (Optional)
  Interval between health checks.  Default: "5s".

* `timeout` -
  (Optional)
  Time before the check is considered failed. Default: "4s"

* `app_start_timeout` -
  (Optional)
  A maximum time limit on application initialization, measured from moment the application successfully
  replies to a healthcheck until it is ready to serve traffic. Default: "300s"

<a name="nested_liveness_check"></a>The `liveness_check` block supports:

* `path` -
  (Required)
  The request path.

* `host` -
  (Optional)
  Host header to send when performing a HTTP Readiness check. Example: "myapp.appspot.com"

* `failure_threshold` -
  (Optional)
  Number of consecutive failed checks required before considering the VM unhealthy. Default: 4.

* `success_threshold` -
  (Optional)
  Number of consecutive successful checks required before considering the VM healthy. Default: 2.

* `check_interval` -
  (Optional)
  Interval between health checks.

* `timeout` -
  (Optional)
  Time before the check is considered failed. Default: "4s"

* `initial_delay` -
  (Optional)
  The initial delay before starting to execute the checks. Default: "300s"

<a name="nested_network"></a>The `network` block supports:

* `forwarded_ports` -
  (Optional)
  List of ports, or port pairs, to forward from the virtual machine to the application container.

* `instance_ip_mode` -
  (Optional, [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html))
  Prevent instances from receiving an ephemeral external IP address.
  Possible values are: `EXTERNAL`, `INTERNAL`.

* `instance_tag` -
  (Optional)
  Tag to apply to the instance during creation.

* `name` -
  (Required)
  Google Compute Engine network where the virtual machines are created. Specify the short name, not the resource path.

* `subnetwork` -
  (Optional)
  Google Cloud Platform sub-network where the virtual machines are created. Specify the short name, not the resource path.
  If the network that the instance is being created in is a Legacy network, then the IP address is allocated from the IPv4Range.
  If the network that the instance is being created in is an auto Subnet Mode Network, then only network name should be specified (not the subnetworkName) and the IP address is created from the IPCidrRange of the subnetwork that exists in that zone for that network.
  If the network that the instance is being created in is a custom Subnet Mode Network, then the subnetworkName must be specified and the IP address is created from the IPCidrRange of the subnetwork.
  If specified, the subnetwork must exist in the same region as the App Engine flexible environment application.

* `session_affinity` -
  (Optional)
  Enable session affinity.

<a name="nested_resources"></a>The `resources` block supports:

* `cpu` -
  (Optional)
  Number of CPU cores needed.

* `disk_gb` -
  (Optional)
  Disk size (GB) needed.

* `memory_gb` -
  (Optional)
  Memory (GB) needed.

* `volumes` -
  (Optional)
  List of ports, or port pairs, to forward from the virtual machine to the application container.
  Structure is [documented below](#nested_resources_volumes).


<a name="nested_resources_volumes"></a>The `volumes` block supports:

* `name` -
  (Required)
  Unique name for the volume.

* `volume_type` -
  (Required)
  Underlying volume type, e.g. 'tmpfs'.

* `size_gb` -
  (Required)
  Volume size in gigabytes.

<a name="nested_flexible_runtime_settings"></a>The `flexible_runtime_settings` block supports:

* `operating_system` -
  (Optional)
  Operating System of the application runtime.

* `runtime_version` -
  (Optional)
  The runtime version of an App Engine flexible application.

<a name="nested_handlers"></a>The `handlers` block supports:

* `url_regex` -
  (Optional)
  URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.
  All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.

* `security_level` -
  (Optional)
  Security (HTTPS) enforcement for this URL.
  Possible values are: `SECURE_DEFAULT`, `SECURE_NEVER`, `SECURE_OPTIONAL`, `SECURE_ALWAYS`.

* `login` -
  (Optional)
  Methods to restrict access to a URL based on login status.
  Possible values are: `LOGIN_OPTIONAL`, `LOGIN_ADMIN`, `LOGIN_REQUIRED`.

* `auth_fail_action` -
  (Optional)
  Actions to take when the user is not logged in.
  Possible values are: `AUTH_FAIL_ACTION_REDIRECT`, `AUTH_FAIL_ACTION_UNAUTHORIZED`.

* `redirect_http_response_code` -
  (Optional)
  30x code to use when performing redirects for the secure field.
  Possible values are: `REDIRECT_HTTP_RESPONSE_CODE_301`, `REDIRECT_HTTP_RESPONSE_CODE_302`, `REDIRECT_HTTP_RESPONSE_CODE_303`, `REDIRECT_HTTP_RESPONSE_CODE_307`.

* `script` -
  (Optional)
  Executes a script to handle the requests that match this URL pattern.
  Only the auto value is supported for Node.js in the App Engine standard environment, for example "script:" "auto".
  Structure is [documented below](#nested_handlers_handlers_script).

* `static_files` -
  (Optional)
  Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files.
  Static file handlers describe which files in the application directory are static files, and which URLs serve them.
  Structure is [documented below](#nested_handlers_handlers_static_files).


<a name="nested_handlers_handlers_script"></a>The `script` block supports:

* `script_path` -
  (Required)
  Path to the script from the application root directory.

<a name="nested_handlers_handlers_static_files"></a>The `static_files` block supports:

* `path` -
  (Optional)
  Path to the static files matched by the URL pattern, from the application root directory.
  The path can refer to text matched in groupings in the URL pattern.

* `upload_path_regex` -
  (Optional)
  Regular expression that matches the file paths for all files that should be referenced by this handler.

* `http_headers` -
  (Optional)
  HTTP headers to use for all responses from these URLs.
  An object containing a list of "key:value" value pairs.".

* `mime_type` -
  (Optional)
  MIME type used to serve all files served by this handler.
  Defaults to file-specific MIME types, which are derived from each file's filename extension.

* `expiration` -
  (Optional)
  Time a static file served by this handler should be cached by web proxies and browsers.
  A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".
  Default is '0s'

* `require_matching_file` -
  (Optional)
  Whether this handler should match the request if the file referenced by the handler does not exist.

* `application_readable` -
  (Optional)
  Whether files should also be uploaded as code data. By default, files declared in static file handlers are
  uploaded as static data and are only served to end users; they cannot be read by the application. If enabled,
  uploads are charged against both your code and static data storage resource quotas.

<a name="nested_api_config"></a>The `api_config` block supports:

* `auth_fail_action` -
  (Optional)
  Action to take when users access resources that require authentication.
  Default value is `AUTH_FAIL_ACTION_REDIRECT`.
  Possible values are: `AUTH_FAIL_ACTION_REDIRECT`, `AUTH_FAIL_ACTION_UNAUTHORIZED`.

* `login` -
  (Optional)
  Level of login required to access this resource.
  Default value is `LOGIN_OPTIONAL`.
  Possible values are: `LOGIN_OPTIONAL`, `LOGIN_ADMIN`, `LOGIN_REQUIRED`.

* `script` -
  (Required)
  Path to the script from the application root directory.

* `security_level` -
  (Optional)
  Security (HTTPS) enforcement for this URL.
  Possible values are: `SECURE_DEFAULT`, `SECURE_NEVER`, `SECURE_OPTIONAL`, `SECURE_ALWAYS`.

* `url` -
  (Optional)
  URL to serve the endpoint at.

<a name="nested_deployment"></a>The `deployment` block supports:

* `zip` -
  (Optional)
  Zip File
  Structure is [documented below](#nested_deployment_zip).

* `files` -
  (Optional)
  Manifest of the files stored in Google Cloud Storage that are included as part of this version.
  All files must be readable using the credentials supplied with this call.
  Structure is [documented below](#nested_deployment_files).

* `container` -
  (Optional)
  The Docker image for the container that runs the version.
  Structure is [documented below](#nested_deployment_container).

* `cloud_build_options` -
  (Optional)
  Options for the build operations performed as a part of the version deployment. Only applicable when creating a version using source code directly.
  Structure is [documented below](#nested_deployment_cloud_build_options).


<a name="nested_deployment_zip"></a>The `zip` block supports:

* `source_url` -
  (Required)
  Source URL

* `files_count` -
  (Optional)
  files count

<a name="nested_deployment_files"></a>The `files` block supports:

* `name` - (Required) The identifier for this object. Format specified above.

* `sha1_sum` -
  (Optional)
  SHA1 checksum of the file

* `source_url` -
  (Required)
  Source URL

<a name="nested_deployment_container"></a>The `container` block supports:

* `image` -
  (Required)
  URI to the hosted container image in Google Container Registry. The URI must be fully qualified and include a tag or digest.
  Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"

<a name="nested_deployment_cloud_build_options"></a>The `cloud_build_options` block supports:

* `app_yaml_path` -
  (Required)
  Path to the yaml file used in deployment, used to determine runtime configuration details.

* `cloud_build_timeout` -
  (Optional)
  The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.
  A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".

<a name="nested_endpoints_api_service"></a>The `endpoints_api_service` block supports:

* `name` -
  (Required)
  Endpoints service name which is the name of the "service" resource in the Service Management API.
  For example "myapi.endpoints.myproject.cloud.goog"

* `config_id` -
  (Optional)
  Endpoints service configuration ID as specified by the Service Management API. For example "2016-09-19r1".
  By default, the rollout strategy for Endpoints is "FIXED". This means that Endpoints starts up with a particular configuration ID.
  When a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID
  and is required in this case.
  Endpoints also has a rollout strategy called "MANAGED". When using this, Endpoints fetches the latest configuration and does not need
  the configuration ID. In this case, configId must be omitted.

* `rollout_strategy` -
  (Optional)
  Endpoints rollout strategy. If FIXED, configId must be specified. If MANAGED, configId must be omitted.
  Default value is `FIXED`.
  Possible values are: `FIXED`, `MANAGED`.

* `disable_trace_sampling` -
  (Optional)
  Enable or disable trace sampling. By default, this is set to false for enabled.

<a name="nested_entrypoint"></a>The `entrypoint` block supports:

* `shell` -
  (Required)
  The format should be a shell command that can be fed to bash -c.

<a name="nested_vpc_access_connector"></a>The `vpc_access_connector` block supports:

* `name` -
  (Required)
  Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.

<a name="nested_automatic_scaling"></a>The `automatic_scaling` block supports:

* `cool_down_period` -
  (Optional)
  The time period that the Autoscaler should wait before it starts collecting information from a new instance.
  This prevents the autoscaler from collecting information when the instance is initializing,
  during which the collected usage would not be reliable. Default: 120s

* `cpu_utilization` -
  (Required)
  Target scaling by CPU usage.
  Structure is [documented below](#nested_automatic_scaling_cpu_utilization).

* `max_concurrent_requests` -
  (Optional)
  Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.
  Defaults to a runtime-specific value.

* `max_idle_instances` -
  (Optional)
  Maximum number of idle instances that should be maintained for this version.

* `max_total_instances` -
  (Optional)
  Maximum number of instances that should be started to handle requests for this version. Default: 20

* `max_pending_latency` -
  (Optional)
  Maximum amount of time that a request should wait in the pending queue before starting a new instance to handle it.

* `min_idle_instances` -
  (Optional)
  Minimum number of idle instances that should be maintained for this version. Only applicable for the default version of a service.

* `min_total_instances` -
  (Optional)
  Minimum number of running instances that should be maintained for this version. Default: 2

* `min_pending_latency` -
  (Optional)
  Minimum amount of time a request should wait in the pending queue before starting a new instance to handle it.

* `request_utilization` -
  (Optional)
  Target scaling by request utilization.
  Structure is [documented below](#nested_automatic_scaling_request_utilization).

* `disk_utilization` -
  (Optional)
  Target scaling by disk usage.
  Structure is [documented below](#nested_automatic_scaling_disk_utilization).

* `network_utilization` -
  (Optional)
  Target scaling by network usage.
  Structure is [documented below](#nested_automatic_scaling_network_utilization).


<a name="nested_automatic_scaling_cpu_utilization"></a>The `cpu_utilization` block supports:

* `aggregation_window_length` -
  (Optional)
  Period of time over which CPU utilization is calculated.

* `target_utilization` -
  (Required)
  Target CPU utilization ratio to maintain when scaling. Must be between 0 and 1.

<a name="nested_automatic_scaling_request_utilization"></a>The `request_utilization` block supports:

* `target_request_count_per_second` -
  (Optional)
  Target requests per second.

* `target_concurrent_requests` -
  (Optional)
  Target number of concurrent requests.

<a name="nested_automatic_scaling_disk_utilization"></a>The `disk_utilization` block supports:

* `target_write_bytes_per_second` -
  (Optional)
  Target bytes written per second.

* `target_write_ops_per_second` -
  (Optional)
  Target ops written per second.

* `target_read_bytes_per_second` -
  (Optional)
  Target bytes read per second.

* `target_read_ops_per_second` -
  (Optional)
  Target ops read per seconds.

<a name="nested_automatic_scaling_network_utilization"></a>The `network_utilization` block supports:

* `target_sent_bytes_per_second` -
  (Optional)
  Target bytes sent per second.

* `target_sent_packets_per_second` -
  (Optional)
  Target packets sent per second.

* `target_received_bytes_per_second` -
  (Optional)
  Target bytes received per second.

* `target_received_packets_per_second` -
  (Optional)
  Target packets received per second.

<a name="nested_manual_scaling"></a>The `manual_scaling` block supports:

* `instances` -
  (Required)
  Number of instances to assign to the service at the start.
  **Note:** When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2
  Modules API set_num_instances() you must use `lifecycle.ignore_changes = ["manual_scaling"[0].instances]` to prevent drift detection.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `apps/{{project}}/services/{{service}}/versions/{{version_id}}`

* `name` -
  Full path to the Version resource in the API. Example, "v1".


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


FlexibleAppVersion can be imported using any of these accepted formats:

* `apps/{{project}}/services/{{service}}/versions/{{version_id}}`
* `{{project}}/{{service}}/{{version_id}}`
* `{{service}}/{{version_id}}`


In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import FlexibleAppVersion using one of the formats above. For example:

```tf
import {
  id = "apps/{{project}}/services/{{service}}/versions/{{version_id}}"
  to = google_app_engine_flexible_app_version.default
}
```

When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), FlexibleAppVersion can be imported using one of the formats above. For example:

```
$ terraform import google_app_engine_flexible_app_version.default apps/{{project}}/services/{{service}}/versions/{{version_id}}
$ terraform import google_app_engine_flexible_app_version.default {{project}}/{{service}}/{{version_id}}
$ terraform import google_app_engine_flexible_app_version.default {{service}}/{{version_id}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
