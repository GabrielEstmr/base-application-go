package main_configs_mongo

import (
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"log/slog"
)

const _MSG_MONGO_DATABASE_BEAN_INITIALIZING = "Initializing mongo database."
const _MSG_MONGO_BEAN_PINGED = "Successfully connected and pinged."
const _MSG_MONGO_BEAN_CLOSING_CONNECTION = "Closing connection of mongo database bean."
const _MSG_ERROR_TO_PING_DATABASE = "Application has been failed to ping mongo database. URI: %s"
const _MSG_ERROR_TO_CLOSE_MONGO_CONNECTION = "Application has been failed to close mongo connection"

//var mongoDatabase *mongo.Database = nil

func GetMongoDBDatabaseBean() *mongo.Database {
	slog.Info(_MSG_MONGO_DATABASE_BEAN_INITIALIZING)
	databaseUri := main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBURI)
	client := GetMongoDBClientBean()
	// Ping the primary
	if errP := client.Ping(context.TODO(), readpref.Primary()); errP != nil {
		log.Fatalf(_MSG_ERROR_TO_PING_DATABASE, databaseUri)
	}
	slog.Info(_MSG_MONGO_BEAN_PINGED)
	return client.Database(main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBDatabaseName))
}

//
//func GetDatabaseConnection() *mongo.Database {
//	slog.Info(_MSG_MONGO_DATABASE_BEAN_INITIALIZING)
//	databaseUri := main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBURI)
//	client := GetMongoDBClientBean()
//	// Ping the primary
//	if errP := client.Ping(context.TODO(), readpref.Primary()); errP != nil {
//		log.Fatalf(_MSG_ERROR_TO_PING_DATABASE, databaseUri)
//	}
//	slog.Info(_MSG_MONGO_BEAN_PINGED)
//	return client.Database(main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBDatabaseName))
//}
