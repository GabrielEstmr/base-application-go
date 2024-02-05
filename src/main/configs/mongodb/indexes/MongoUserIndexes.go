package main_configs_mongo_indexes

import (
	main_configs_mongo_collections "baseapplicationgo/main/configs/mongodb/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const _DOCUMENT_NUMBER = "documentNumber"
const _EMAIL_REPO_STATUS = "status"
const _EMAIL_REPO_EMAIL_PARAMS_EMAIL_TYPE = "emailParams.emailTemplateType"
const _EMAIL_REPO_EMAIL_EVENT_ID = "eventId"

func GetUserIndexes() []IndexConfig {
	return []IndexConfig{
		{
			CollectionName: main_configs_mongo_collections.USERS_COLLECTION_NAME,
			Index: mongo.IndexModel{
				Keys: bson.D{{_DOCUMENT_NUMBER, 1}},
			},
		},
		{
			CollectionName: main_configs_mongo_collections.EMAILS_COLLECTION_NAME,
			Index: mongo.IndexModel{
				Keys:    bson.D{{_EMAIL_REPO_STATUS, 1}},
				Options: options.Index().SetUnique(false).SetExpireAfterSeconds(30),
			},
		},
		{
			CollectionName: main_configs_mongo_collections.EMAILS_COLLECTION_NAME,
			Index: mongo.IndexModel{
				Keys:    bson.D{{_EMAIL_REPO_EMAIL_PARAMS_EMAIL_TYPE, 1}},
				Options: options.Index().SetUnique(false).SetExpireAfterSeconds(30),
			},
		},
		{
			CollectionName: main_configs_mongo_collections.EMAILS_COLLECTION_NAME,
			Index: mongo.IndexModel{
				Keys:    bson.D{{_EMAIL_REPO_EMAIL_EVENT_ID, 1}},
				Options: options.Index().SetUnique(true).SetExpireAfterSeconds(30),
			},
		},
	}
}
