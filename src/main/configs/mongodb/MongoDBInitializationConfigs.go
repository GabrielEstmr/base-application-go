package main_configs_mongo

import (
	main_configs_mongo_indexes "baseapplicationgo/main/configs/mongodb/indexes"
	"log/slog"
)

const _MSG_INITIALIZING_MONGODB_BEANS = "Initializing MongoDB configuration beans"
const _MSG_MONGODB_BEANS_INITIATED = "MongoDB configuration beans successfully initiated"

func InitMongoConfigBeans() {
	slog.Info(_MSG_INITIALIZING_MONGODB_BEANS)
	GetMongoDBDatabaseBean()
	CreateIndexes(main_configs_mongo_indexes.GetUserIndexes())
	slog.Info(_MSG_MONGODB_BEANS_INITIATED)
}
