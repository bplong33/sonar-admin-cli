/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package groups

import (
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var GroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage and View group information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	GroupCmd.AddCommand(groupCreate)
	GroupCmd.AddCommand(groupSearch)
	GroupCmd.AddCommand(groupAddMembership)
}
