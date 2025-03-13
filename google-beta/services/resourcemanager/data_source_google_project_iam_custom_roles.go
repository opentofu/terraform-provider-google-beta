// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package resourcemanager

import (
	"context"
	"fmt"
	"path"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/iam/v1"
)

func DataSourceGoogleProjectIamCustomRoles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectIamCustomRolesRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"show_deleted": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"view": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "BASIC",
				ValidateFunc: validateViewProjectIamCustomRoles,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"permissions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"role_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"stage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func validateViewProjectIamCustomRoles(val interface{}, key string) ([]string, []error) {
	v := val.(string)
	var errs []error

	if v != "BASIC" && v != "FULL" {
		errs = append(errs, fmt.Errorf("%q must be either 'BASIC' or 'FULL', got %q", key, v))
	}

	return nil, errs
}

func dataSourceProjectIamCustomRolesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for custom roles: %s", err)
	}

	roles := make([]map[string]interface{}, 0)

	showDeleted := d.Get("show_deleted").(bool)
	view := d.Get("view").(string)

	request := config.NewIamClient(userAgent).Projects.Roles.List("projects/" + project).ShowDeleted(showDeleted).View(view)

	err = request.Pages(context.Background(), func(roleList *iam.ListRolesResponse) error {
		for _, role := range roleList.Roles {
			var permissions []string

			switch view {
			case "BASIC":
				permissions = []string{}
			case "FULL":
				permissions = role.IncludedPermissions
			default:
				return fmt.Errorf("Unsupported view type: %s", view)
			}

			roles = append(roles, map[string]interface{}{
				"deleted":     role.Deleted,
				"description": role.Description,
				"id":          role.Name,
				"name":        role.Name,
				"permissions": permissions,
				"role_id":     path.Base(role.Name),
				"stage":       role.Stage,
				"title":       role.Title,
			})
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error retrieving project custom roles: %s", err)
	}

	if err := d.Set("roles", roles); err != nil {
		return fmt.Errorf("Error setting project custom roles: %s", err)
	}

	d.SetId("projects/" + project)

	return nil
}
