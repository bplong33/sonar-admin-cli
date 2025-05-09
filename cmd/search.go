/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving Projects......\n")

		c := services.NewProjectClient()
		opts := map[string]string{"visibility": "private"}
		projects := c.GetProjects(opts)

		// Print Table Header
		fmt.Printf("%-30s%-30s%-15s%-20s\n", "Name", "Key", "Visibility", "LastAnalysisDate")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		for _, p := range projects {
			fmt.Printf("%-30s%-30s%-15s%-20s\n",
				p.Name,
				p.Key,
				p.Visibility,
				p.LastAnalysisDate,
			)
		}
	},
}

func init() {
	projectCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
