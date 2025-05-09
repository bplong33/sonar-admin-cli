module github.com/bplong33/sonar-admin-cli

go 1.23.0

replace github.com/bplong33/gonarqube => ../gonarqube/

require (
	github.com/bplong33/gonarqube v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)
