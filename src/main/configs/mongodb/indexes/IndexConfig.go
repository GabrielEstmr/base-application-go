package main_configs_mongo_indexes

import "go.mongodb.org/mongo-driver/mongo"

type IndexConfig struct {
	CollectionName string
	Index          mongo.IndexModel
}
