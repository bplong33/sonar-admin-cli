/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package permissions

import (
	"fmt"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

// Usage:
//   sonar-admin-cli permissions modify [ACTION] [FLAGS]
//
//     ACTIONS - Either "add" or "remove"
// 	FLAGS - a permissions flag is required. Should be one of the following:
// 	  - admin, codeviewer, issueadmin, securityhotspotadmin, scan, user

// permissionModify represents the modify command
var permissionModify = &cobra.Command{
	Use:       "modify",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"add", "remove"},
	Short:     "Remove `permission` from `group`",
	Long: `Remove a given permission from a group. If visibility and/or query filters
are provided, this operation will only be done on projects matching these filters.

 sonaradmin permissions modify [ACTION] [FLAGS]

Examples:
	
  # remove user permission from the sonar-user group on all projects
  sonaradmin permissions modify remove -g sonar-user -P user
  
  # remove admin permission from the custom-group group on all projects containing myProject
  sonaradmin permissions modify remove -g custom-group -P admin --query myProject
  
  # remove user permission from the custom-group group on all private projects
  sonaradmin permissions modify remove --group custom-group --permission admin -v private

  # add admin permissions to a custom-group for a 2 specific projects
  sonaradmin permissions modify add --group custom-group --permission admin -p my-project,side-project
`,
	Run: func(cmd *cobra.Command, args []string) {
		ModifyProjectPermissions(args)
	},
}

func init() {
	permissionModify.Flags().StringVarP(&Group, "group", "g", "sonar-user", "Target group (default: sonar-user)")
	permissionModify.Flags().StringVarP(&Permission, "permission", "P", "", `Permission to be removed. Must be one of [admin, codeviewer, issueadmin, securityhotspotadmin, scan, user]`)
	permissionModify.Flags().StringVarP(&Visibility, "visibility", "v", "", "Visibility filter [public, private]")
	permissionModify.Flags().StringVarP(&Query, "query", "q", "", "Filter only projects whose name or key contain the supplied string")
	permissionModify.Flags().StringVarP(&ProjectFilter, "projects", "p", "", "A comma-separated list of project keys")
	permissionModify.MarkFlagRequired("permission")
}

func ModifyProjectPermissions(args []string) {
	// get config
	config := common.GetConfig()

	// call client to remove permissions
	c := services.NewPermissionClient(config.URL, config.Token)
	failed, err := c.BulkModifyPermission(
		args[0], Group, Permission, Visibility, Query, ProjectFilter)
	if err != nil {
		fmt.Println("Error while modifying permissions:", err)
	}
	if len(failed) > 0 {
		fmt.Printf("Failed to modify permissions on %d projects:\n", len(failed))
		for _, proj := range failed {
			fmt.Println("\t", proj.Key)
		}
	}
}
