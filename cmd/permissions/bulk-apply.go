/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package permissions

import (
	"fmt"
	"log"
	"net/url"

	"github.com/bplong33/gonarqube/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bulkApplyCmd represents the bulkApply command
var bulkApplyCmd = &cobra.Command{
	Use:   "bulk-apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("bulkApply called")
		params := url.Values{}
		params.Add("templateName", TemplateName)
		if Visibility != "" {
			params.Add("visibility", Visibility)
		}

		// get config
		active_env := viper.Get("sonar.active_env")
		host := viper.GetString(fmt.Sprintf("sonar.%s.host", active_env))
		token := viper.GetString(fmt.Sprintf("sonar.%s.token", active_env))

		// parse url
		hostUrl, err := url.Parse(host)
		if err != nil {
			log.Panicf("Invalid hostname. Please verify your config (default location: `%s`).", viper.ConfigFileUsed())
		}

		// build client
		c := services.NewPermissionClient(hostUrl, token)
		statusCode, respStatus := c.BulkApplyTemplate(params)
		if statusCode < 300 {
			fmt.Println("Successfully applied template", TemplateName)
		} else {
			fmt.Println("Failed to apply template", TemplateName, ":", respStatus)
		}
	},
}

func init() {
	bulkApplyCmd.Flags().StringVarP(&TemplateName, "name", "n", "Default template", "Name of Template")
	bulkApplyCmd.Flags().StringVarP(&Visibility, "visibility", "v", "", "Visibility filter [public, private]")
}
