/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package groups

import (
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var GroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Create a new SonarQube group",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GroupsCmd.AddCommand(createCmd)
}
