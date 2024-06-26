package main_configs

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_configs_ff "baseapplicationgo/main/configs/ff"
	main_configs_log "baseapplicationgo/main/configs/log"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log/slog"
)

const _MSG_INITIALIZING_APPLICATION_BEANS = "Initializing application's configuration beans"
const _MSG_APPLICATION_BEANS_INITIATED = "Application configuration beans successfully initiated"

func InitConfigBeans() {
	slog.Info(_MSG_INITIALIZING_APPLICATION_BEANS)
	main_configs_messages.GetMessagesConfigBean()
	main_configs_yml.GetYmlConfigBean()
	main_configs_mongo.InitMongoConfigBeans()
	main_configs_cache.GetRedisClusterBean()
	main_configs_ff.GetFfConfigDataBean()
	main_configs_rabbitmq.SetAmqpConfig()
	main_configs_log.GetLogConfigBean()
	slog.Info(_MSG_APPLICATION_BEANS_INITIATED)
}
