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

var (
	Visibility string
	Query      string
	Projects   string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Args:  cobra.NoArgs,
	Short: "Search Sonarqube projects based on various conditions",
	Long: `Search Sonarqube projects based on various conditions. Running the command
with no flags or filters will return all projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving Projects......\n")

		c := services.NewProjectClient()
		opts := map[string]string{}
		if Visibility != "" {
			opts["visibility"] = Visibility
		}
		if Query != "" {
			opts["q"] = Query
		}
		if Projects != "" {
			opts["projects"] = Projects
		}
		// opts := map[string]string{"visibility": "private"}
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
	searchCmd.Flags().StringVarP(&Visibility, "visibility", "v", "",
		"Filter to specific project visibility ('public' or 'private').")
	searchCmd.Flags().StringVarP(&Query, "query", "q", "",
		"Filter only projects whose name or key contain the supplied string")
	searchCmd.Flags().StringVarP(&Projects, "projects", "P", "",
		"A comma-separated list of project keys")
	searchCmd.Flags().BoolP("provisioned", "p", false,
		"Only show projects that have been provisioned.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
