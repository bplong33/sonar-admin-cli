package main

import (
	"github.com/bplong33/sonar-admin-cli/cmd"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Unable to load environment variables")
	// }

	cmd.Execute()
}
