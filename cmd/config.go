/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flag variables
var list bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get or set configuration variables",
	Long: `Manage your sonar-admin-cli configuration. Examples:

	sonar-admin-cli config --list  // to list config settings
	sonar-admin-cli config --add env2 --host <hostname> --token <token>
	sonar-admin-cli config --setenv env2`,
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			fmt.Println(viper.Get("sonar.env1.host"))
		}
	},
}

func init() {
	configCmd.Flags().BoolVarP(&list, "list", "l", false, "List current configuration settings")
}
