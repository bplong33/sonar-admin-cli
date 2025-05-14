/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get or set configuration variables",
	Long: `Manage your sonar-admin-cli configuration.

Examples:
	sonar-admin-cli config --list  // to list config settings
	sonar-admin-cli config --add env2 --host <hostname> --token <token>
	sonar-admin-cli config --setenv env2`,
	Run: func(cmd *cobra.Command, args []string) {
		if List {
			PrintConfig()
		}
	},
}

func init() {
	ConfigCmd.Flags().BoolVar(&List, "list", false, "List current configuration settings")
}

// util functions
func PrintConfig() {
	fmt.Println(viper.Get("sonar.env1.host"))
}
