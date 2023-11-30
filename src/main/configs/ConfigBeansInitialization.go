package main_configs

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log"
)

const _MSG_INITIALIZING_APPLICATION_BEANS = "Initializing application's configuration beans"
const _MSG_APPLICATION_BEANS_INITIATED = "Application configuration beans successfully initiated"

func InitConfigBeans() {
	log.Println(_MSG_INITIALIZING_APPLICATION_BEANS)
	main_configs_messages.GetMessagesConfigBean()
	main_configs_yml.GetYmlConfigBean()
	main_configs_mongo.InitMongoConfigBeans()
	main_configs_cache.GetRedisClusterBean()
	log.Println(_MSG_APPLICATION_BEANS_INITIATED)
}
