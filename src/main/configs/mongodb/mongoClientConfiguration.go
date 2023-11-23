package main_configs_mongo

import (
	configsYml "baseapplicationgo/main/configs/yml"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

const _MSG_MONGO_BEAN_INITIALIZING = "Initializing mongo database bean."
const _MSG_MONGO_BEAN_PINGED = "Successfully connected and pinged."
const _MSG_MONGO_BEAN_CLOSING_CONNECTION = "Closing connection of mongo database bean."
const _MSG_ERROR_TO_CONNECT_TO_DATABASE = "Application has been failed to connect to mongo database. URI: %s"
const _MSG_ERROR_TO_PING_DATABASE = "Application has been failed to ping mongo database. URI: %s"

const MONGO_URI_NAME = "MongoDB.URI"
const MONGO_DATABASE_NAME = "MongoDB.DatabaseName"

var MongoDatabase *mongo.Database = nil
var once sync.Once

func GetDatabaseBean() *mongo.Database {
	once.Do(func() {

		if MongoDatabase == nil {
			MongoDatabase = getDatabaseConnection()
		}

	})
	return MongoDatabase
}

func CloseConnection() {
	log.Println(_MSG_MONGO_BEAN_CLOSING_CONNECTION)
	err := MongoDatabase.Client().Disconnect(context.TODO())
	if err != nil {
		return
	}
}

func getDatabaseConnection() *mongo.Database {
	log.Println(_MSG_MONGO_BEAN_INITIALIZING)
	databaseUri := configsYml.GetYmlValueByName(MONGO_URI_NAME)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf(_MSG_ERROR_TO_CONNECT_TO_DATABASE, databaseUri)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf(_MSG_ERROR_TO_PING_DATABASE, databaseUri)
		panic(err)
	}
	log.Println(_MSG_MONGO_BEAN_PINGED)
	return client.Database(MONGO_DATABASE_NAME)
}
