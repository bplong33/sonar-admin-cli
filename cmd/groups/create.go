package groups

import (
	"fmt"
	"os"

	"github.com/bplong33/gonarqube/services"
	"github.com/bplong33/sonar-admin-cli/common"
	"github.com/spf13/cobra"
)

// groupCreate represents the create command
var groupCreate = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new SonarQube group",
	Run: func(cmd *cobra.Command, args []string) {
		AddGroup(args)
	},
}

func init() {
	groupCreate.Flags().StringVarP(&GroupDesc, "description", "d", "", "Description of group")
}

func AddGroup(args []string) {
	// get config
	config := common.GetConfig()
	c := services.NewGroupClient(config.URL, config.Token)
	statusCode, err := c.CreateGroup(args[0], GroupDesc)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if statusCode >= 300 {
		fmt.Println("Encountered an error while creating the group.")
	} else {
		fmt.Println("Successfully Created Group", args[0])
	}
}
