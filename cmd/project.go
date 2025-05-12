/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Args:  cobra.NoArgs,
	Short: "A collection of sub commands for interacting with Sonarqube projects",
	Long: `Provide a collection of sub-commands that support CRUD operations on 
on Sonarqube projects.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("project called")
	// },
}

func init() {
	projectCmd.AddCommand(searchCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
