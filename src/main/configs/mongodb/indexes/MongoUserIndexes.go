package main_configs_mongo_indexes

import (
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const _DOCUMENT_NUMBER = "documentNumber"

func GetUserIndexes() []IndexConfig {
	return []IndexConfig{
		{
			CollectionName: main_configs_mongo_collections.USERS_COLLECTION_NAME,
			Index: mongo.IndexModel{
				Keys: bson.D{{_DOCUMENT_NUMBER, 1}},
			},
		},
	}
}
