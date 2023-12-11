package main_configs_mongo

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"log/slog"
	"sync"
)

const _MSG_MONGO_BEAN_INITIALIZING = "Initializing mongo database client."
const _MSG_MONGO_BEAN_PINGED = "Successfully connected and pinged."
const _MSG_MONGO_BEAN_CLOSING_CONNECTION = "Closing connection of mongo database bean."
const _MSG_ERROR_TO_CONNECT_TO_DATABASE = "Application has been failed to connect to mongo database. URI: %s"
const _MSG_ERROR_TO_PING_DATABASE = "Application has been failed to ping mongo database. URI: %s"
const _MSG_ERROR_TO_CLOSE_MONGO_CONNECTION = "Application has been failed to close mongo connection"

const MONGO_URI_NAME = "MongoDB.URI"
const MONGO_DATABASE_NAME = "MongoDB.DatabaseName"

var once sync.Once
var mongoDatabase *mongo.Database = nil
var mongoDatabaseBean mongo.Database

func GetMongoDBClient() *mongo.Database {
	once.Do(func() {
		if mongoDatabase == nil {
			mongoDatabaseBean = getDatabaseConnection()
			mongoDatabase = &mongoDatabaseBean
		}
	})
	return mongoDatabase
}

func CloseConnection() {
	slog.Info(_MSG_MONGO_BEAN_CLOSING_CONNECTION)
	err := mongoDatabase.Client().Disconnect(context.TODO())
	if err != nil {
		main_configs_error.FailOnError(err, _MSG_ERROR_TO_CLOSE_MONGO_CONNECTION)
	}
}

// IF Connection failed, how to solve
func getDatabaseConnection() mongo.Database {
	slog.Info(_MSG_MONGO_BEAN_INITIALIZING)
	databaseUri := main_configs_yml.GetYmlValueByName(MONGO_URI_NAME)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf(_MSG_ERROR_TO_CONNECT_TO_DATABASE, databaseUri)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf(_MSG_ERROR_TO_PING_DATABASE, databaseUri)
		panic(err)
	}
	slog.Info(_MSG_MONGO_BEAN_PINGED)
	return *client.Database(main_configs_yml.GetYmlValueByName(MONGO_DATABASE_NAME))
}
