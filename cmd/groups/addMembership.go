package groups

import (
	"fmt"
	"os"
	"strings"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

// groupAddMembership represents the addMembership command
var groupAddMembership = &cobra.Command{
	Use:  "add-user",
	Args: cobra.ExactArgs(2),
	// Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Short: "Add a user to the group",
	Long:  `sonaradmin group add-user [USERQUERY] [GROUPQUERY]`,
	Run: func(cmd *cobra.Command, args []string) {
		addUser(args)
	},
}

func init() {
}

func addUser(args []string) {
	config := common.GetConfig()
	userClient := services.NewUserClient(config.URL, config.Token)
	groupClient := services.NewGroupClient(config.URL, config.Token)

	var userId string
	var groupId string

	users, err := userClient.GetUsers(args[0], true, "")
	if err != nil {
		fmt.Println("Encountered an error while searching users:", err)
	}

	// get user
	if len(users) == 0 {
		fmt.Println("User not found")
		os.Exit(1)
	} else if len(users) > 1 {
		fmt.Printf("Multiple users found matching the given search `%s`:\n", args[0])
		// TODO: Print out users
		fmt.Printf("%-15s%-25s\n", "Login", "Name")
		fmt.Println(strings.Repeat("-", 50))
		for _, user := range users {
			fmt.Printf("%-15s%-25s\n", user.Login, user.Name)
		}
		os.Exit(1)
	} else {
		userId = users[0].Id
	}

	// get group
	groups, err := groupClient.GetGroups(args[1], false)
	if err != nil {
		fmt.Println("Encountered an error while searching groups:", err)
	}

	if len(groups) == 0 {
		fmt.Println("User not found")
		os.Exit(1)
	} else if len(groups) > 1 {
		fmt.Printf("Multiple groups found matching the given search `%s`:\n\n", args[0])
		// TODO: Print out users
		fmt.Printf("%-20s%.25s\n", "Name", "Description")
		fmt.Println(strings.Repeat("-", 50))
		for _, group := range groups {
			fmt.Printf("%-20s%.25s\n", group.Name, group.Description)
		}
		os.Exit(1)
	} else {
		groupId = groups[0].Id
	}

	// fmt.Println(userId)
	// fmt.Println(groupId)

	status, err := groupClient.AddUserToGroup(userId, groupId)
	if err != nil {
		fmt.Println("Encountered an error while adding user to group:", err)
		os.Exit(1)
	}
	if status >= 300 {
		fmt.Printf("Failed to add user %s to group %s - Status Code %d\n",
			users[0].Login,
			groups[0].Name,
			status,
		)
		os.Exit(1)
	}

	fmt.Println("Successfully added user to group.")
}
