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

var once sync.Once
var mongoDatabase *mongo.Database = nil
var mongoDatabaseBean mongo.Database

func GetMongoDbBean() *mongo.Database {
	once.Do(func() {
		if mongoDatabase == nil {
			mongoDatabaseBean = getDatabaseConnection()
			mongoDatabase = &mongoDatabaseBean
		}
	})
	return mongoDatabase
}

// TODO: check if its really necessary
func CloseConnection() {
	log.Println(_MSG_MONGO_BEAN_CLOSING_CONNECTION)
	err := mongoDatabase.Client().Disconnect(context.TODO())
	if err != nil {
		return
	}
}

// TODO: check func
// WHen to close connection
// IF Connection failed, how to solve
func getDatabaseConnection() mongo.Database {
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
	return *client.Database(configsYml.GetYmlValueByName(MONGO_DATABASE_NAME))
}
