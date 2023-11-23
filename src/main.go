package main

import (
	mainConfigsYml "baseapplicationgo/main/configurations/yml"
	"log"
)

func init() {

	mainConfigsYml.GetYmlConfigBean()

}

func main() {

	host := mainConfigsYml.GetYmlConfigBean().Spring.Datasource.PostgresHost

	value := mainConfigsYml.ReplaceEnvNameToValue(host)

	log.Print(value)

}
