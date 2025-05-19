package users

import (
	"fmt"
	"os"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

var userSearch = &cobra.Command{
	Use:   "search",
	Short: "Search SonarQube users",
	Run: func(cmd *cobra.Command, args []string) {
		GetUsers(args)
	},
}

func init() {
	userSearch.Flags().StringVarP(&Query, "query", "q", "", "Query users login, name, or email containing STRING")
	userSearch.Flags().BoolVar(&Inactive, "inactive", false, "Set to search for inactive users")
	userSearch.Flags().StringVarP(&GroupName, "group", "g", "", "Filter for users within a given group")
}

func GetUsers(args []string) {
	config := common.GetConfig()
	client := services.NewUserClient(config.URL, config.Token)

	// TODO: if GroupId is provided, use GroupClient.GetGroups to find
	// the id matching the name.

	users, err := client.GetUsers(Query, Inactive, GroupName)
	if err != nil {
		fmt.Println("Encountered an error while querying groups:", err)
		os.Exit(1)
	}

	if len(users) > 0 {
		fmt.Printf("%-16s%-25s%-25s%-12s%-45s\n", "Login", "Name", "Email", "Active", "ID")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		for _, user := range users {
			// TODO: handle long descriptions better
			fmt.Printf("%-16s%-25s%-25s%-12t%-45s\n",
				user.Login,
				user.Name,
				user.Email,
				user.Active,
				user.Id,
			)
		}
	} else {
		fmt.Println("No users found")
	}
}
