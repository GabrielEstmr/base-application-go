package main_configs_mongo

import (
	main_configs_mongo_indexes "baseapplicationgo/main/configs/mongodb/indexes"
	"context"
	"log"
)

const _MSG_INDEX_CREATED = "Index Created. Collection: %s, IndexName: %s"

func CreateIndexes(indexes []main_configs_mongo_indexes.IndexConfig) {
	for _, value := range indexes {
		name, err := GetMongoDbBean().Collection(value.CollectionName).Indexes().CreateOne(context.TODO(), value.Index)
		if err != nil {
			panic(err)
		}
		log.Printf(_MSG_INDEX_CREATED, value.CollectionName, name)
	}
}
