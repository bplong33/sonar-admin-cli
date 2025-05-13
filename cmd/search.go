/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flag variables
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

		// TODO: Make a util function for getting configs
		// get configs
		active_env := viper.Get("sonar.active_env")
		host := viper.GetString(fmt.Sprintf("sonar.%s.host", active_env))
		token := viper.GetString(fmt.Sprintf("sonar.%s.token", active_env))

		// parse url
		hostUrl, err := url.Parse(host)
		if err != nil {
			log.Panicln("Invalid hostname. Please verify your config (default location: `~/.sonar-admin-cli.toml`).")
		}

		// build client
		c := services.NewProjectClient(hostUrl, token)

		// check flags and build url.Values
		opts := url.Values{}
		if Visibility != "" {
			opts.Add("visibility", Visibility)
		}
		if Query != "" {
			opts.Add("q", Query)
		}
		if Projects != "" {
			opts.Add("projects", Projects)
		}

		// get projects
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

	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
