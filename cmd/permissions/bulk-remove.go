/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package permissions

import (
	"fmt"
	"log"
	"net/url"

	"github.com/bplong33/gonarqube/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bulkRemoveCmd represents the bulkRemove command
var bulkRemoveCmd = &cobra.Command{
	Use:   "bulk-remove",
	Short: "Remove `permission` from `group`",
	Long: `Remove a given permission from a group. If visibility and/or query filters
are provided, this operation will only be done on projects matching these filters.

Examples:
	# remove user permission from the sonar-user group on all projects
	sonar-admin-cli permissions bulk-remove -g sonar-user -p user
	
	# remove admin permission from the custom-group group on all projects containing myProject
	sonar-admin-cli permissions bulk-remove -g custom-group -p admin --query myProject

	# remove user permission from the custom-group group on all private projects
	sonar-admin-cli permissions bulk-remove --group custom-group --permission admin -v private
`,
	Run: func(cmd *cobra.Command, args []string) {
		// get config
		active_env := viper.Get("sonar.active_env")
		host := viper.GetString(fmt.Sprintf("sonar.%s.host", active_env))
		token := viper.GetString(fmt.Sprintf("sonar.%s.token", active_env))

		// parse url
		hostUrl, err := url.Parse(host)
		if err != nil {
			log.Panicf("Invalid hostname. Please verify your config (default location: `%s`).", viper.ConfigFileUsed())
		}

		// call client to remove permissions
		c := services.NewPermissionClient(hostUrl, token)
		failed, err := c.BulkRemovePermission(Group, Permission, Visibility, Query)
		if err != nil {
			fmt.Println("Error while removing permissions:", err)
		}
		if len(failed) > 0 {
			fmt.Printf("Failed to modify permissions on %d projects:\n", len(failed))
			for _, proj := range failed {
				fmt.Println("\t", proj.Key)
			}
		}
	},
}

func init() {
	bulkRemoveCmd.Flags().StringVarP(
		&Group, "group", "g", "sonar-user", "Target group (default: sonar-user)",
	)
	bulkRemoveCmd.Flags().StringVarP(
		&Permission, "permission", "p", "",
		`Permission to be removed. Must be one of [admin, codeviewer, issueadmin, securityhotspotadmin, scan, user]`,
	)
	bulkRemoveCmd.MarkFlagRequired("permission")
	bulkRemoveCmd.Flags().StringVarP(
		&Visibility, "visibility", "v", "", "Visibility filter [public, private]",
	)
	bulkRemoveCmd.Flags().StringVarP(
		&Query, "query", "q", "",
		"Filter only projects whose name or key contain the supplied string",
	)
}
