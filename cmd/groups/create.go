/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package groups

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/bplong33/gonarqube/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new SonarQube group",
	Run: func(cmd *cobra.Command, args []string) {
		AddGroup(args)
	},
}

func init() {
	createCmd.Flags().StringVarP(&GroupDesc, "description", "d", "", "Description of group")
}

func AddGroup(args []string) {
	// get config
	active_env := viper.Get("sonar.active_env")
	host := viper.GetString(fmt.Sprintf("sonar.%s.host", active_env))
	token := viper.GetString(fmt.Sprintf("sonar.%s.token", active_env))

	// parse url
	hostUrl, err := url.Parse(host)
	if err != nil {
		log.Panicf("Invalid hostname. Please verify your config (default location: `%s`).", viper.ConfigFileUsed())
	}
	c := services.NewGroupClient(hostUrl, token)
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
