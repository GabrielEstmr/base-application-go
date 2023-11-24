package main_configs

import (
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	"log"
)

const _MSG_TERMINATING_APPLICATION_BEANS = "Terminating application's configuration beans."
const _MSG_APPLICATION_BEANS_TERMINATED = "Application configuration beans successfully terminated."

func TerminateConfigBeans() {
	log.Println(_MSG_TERMINATING_APPLICATION_BEANS)
	main_configs_mongo.CloseConnection()
	log.Println(_MSG_APPLICATION_BEANS_TERMINATED)
}
