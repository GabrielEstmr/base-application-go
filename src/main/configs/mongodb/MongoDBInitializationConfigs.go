package main_configs_mongo

import (
	main_configs_mongo_indexes "baseapplicationgo/main/configs/mongodb/indexes"
	"log"
)

const _MSG_INITIALIZING_MONGODB_BEANS = "Initializing MongoDB configuration beans"
const _MSG_MONGODB_BEANS_INITIATED = "MongoDB configuration beans successfully initiated"

func InitConfigBeans() {
	log.Println(_MSG_INITIALIZING_MONGODB_BEANS)
	GetMongoDbBean()
	CreateIndexes(main_configs_mongo_indexes.GetUserIndexes())
	log.Println(_MSG_MONGODB_BEANS_INITIATED)
}
