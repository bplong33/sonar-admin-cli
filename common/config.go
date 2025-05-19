package common

import (
	"fmt"
	"log"
	"net/url"

	"github.com/spf13/viper"
)

type Config struct {
	URL   *url.URL
	Token string
}

func GetConfig() *Config {
	active_env := viper.Get("sonar.active_env")
	host := viper.GetString(fmt.Sprintf("sonar.%s.host", active_env))
	token := viper.GetString(fmt.Sprintf("sonar.%s.token", active_env))

	// parse url
	hostUrl, err := url.Parse(host)
	if err != nil {
		log.Panicln("Invalid hostname. Please verify your config (default location: `~/.sonar-admin-cli.toml`).")
	}

	return &Config{URL: hostUrl, Token: token}
}
