package groups

import (
	"fmt"
	"os"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

// groupSearch represents the search command
var groupSearch = &cobra.Command{
	Use:   "search",
	Short: "Search SonarQube groups",
	Run: func(cmd *cobra.Command, args []string) {
		ViewGroup(args)
	},
}

func init() {
	groupSearch.Flags().StringVarP(&Query, "query", "q", "", "Query groups containing a STRING")
	groupSearch.Flags().BoolVar(&Managed, "managed", false, "Filter to only show Managd Groups")
}

func ViewGroup(args []string) {
	config := common.GetConfig()
	client := services.NewGroupClient(config.URL, config.Token)

	groups, err := client.GetGroups(Query, Managed)
	if err != nil {
		fmt.Println("Encountered an error while querying groups:", err)
		os.Exit(1)
	}

	if len(groups) > 0 {
		fmt.Printf("%-30s%-45s%-12s%-45s\n", "Name", "Description", "Managed", "ID")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		for _, group := range groups {
			// TODO: handle long descriptions better
			fmt.Printf("%-30s%-45.42s%-12t%-45s\n",
				group.Name,
				group.Description,
				group.Managed,
				group.Id,
			)
		}
	} else {
		fmt.Println("No groups found")
	}
}
