/*
Copyright © 2025 Brandon Long <bplong96@gmail.com>
*/
package main

import (
	"log"

	"github.com/bplong33/sonar-admin-cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load environment variables")
	}

	cmd.Execute()
}
