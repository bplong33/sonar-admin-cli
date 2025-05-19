/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package project

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

// projectSearch represents the search command
var projectSearch = &cobra.Command{
	Use:   "search",
	Args:  cobra.NoArgs,
	Short: "Search Sonarqube projects based on various conditions",
	Long: `Search Sonarqube projects based on various conditions. Running the command
with no flags or filters will return all projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving Projects......\n")

		// get configs
		config := common.GetConfig()
		client := services.NewProjectClient(config.URL, config.Token)

		// check flags and build url.Values
		opts := url.Values{}
		if Visibility != "" {
			opts.Add("visibility", Visibility)
		}
		if Query != "" {
			opts.Add("q", Query)
		}
		if ProjectFilter != "" {
			opts.Add("projects", ProjectFilter)
		}

		// get projects
		projects := client.GetProjects(opts)

		// Print Table Header
		if len(projects) > 0 {
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
		} else {
			fmt.Println("No users found")
		}
	},
}

func init() {
	projectSearch.Flags().StringVarP(
		&Visibility, "visibility", "v", "", "Visibility filter [public, private]",
	)
	projectSearch.Flags().StringVarP(&Query, "query", "q", "",
		"Filter only projects whose name or key contain the supplied string")
	projectSearch.Flags().StringVarP(&ProjectFilter, "projects", "p", "",
		"A comma-separated list of project keys")
	projectSearch.Flags().BoolP("provisioned", "P", false,
		"Only show projects that have been provisioned.")
}
