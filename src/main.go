package main

import (
	main_configurations_yml "baseapplicationgo/main/configurations/yml"
	"log"
)

func init() {

	main_configurations_yml.GetYmlConfigBean()

}

func main() {

	host := main_configurations_yml.GetYmlConfigBean().Spring.Datasource.PostgresHost

	value := main_configurations_yml.ReplaceEnvNameToValue(host)

	log.Print(value)

}
