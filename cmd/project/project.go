package project

import (
	"github.com/spf13/cobra"
)

// ProjectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Args:  cobra.NoArgs,
	Short: "A collection of sub commands for interacting with Sonarqube projects",
	Long: `Provide a collection of sub-commands that support CRUD operations on 
on Sonarqube projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ProjectCmd.AddCommand(projectSearch)
}
