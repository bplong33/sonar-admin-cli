/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package permissions

import (
	"github.com/spf13/cobra"
)

var hidePrivate bool

// PermissionsCmd represents the permissions command
var PermissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Manage SonarQube permissions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	PermissionsCmd.AddCommand(modifyCmd)
}
