package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bplong33/sonar-admin-cli/cmd/config"
	"github.com/bplong33/sonar-admin-cli/cmd/groups"
	"github.com/bplong33/sonar-admin-cli/cmd/permissions"
	"github.com/bplong33/sonar-admin-cli/cmd/project"
	"github.com/bplong33/sonar-admin-cli/cmd/users"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "sonar-admin-cli",
	Aliases: []string{"sonaradmin"},
	Short:   "SonarQube CLI Admin Tool",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var CfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(permissions.PermissionsCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(groups.GroupCmd)
	rootCmd.AddCommand(users.UserCmd)
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/.sonar-admin-cli.toml)")
}

func initConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".sonar-admin-cli")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read in config:", err)
		os.Exit(1)
	}
}
