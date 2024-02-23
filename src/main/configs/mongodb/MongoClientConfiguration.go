package main_configs_mongo

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log/slog"
	"sync"
)

const _MSG_MONGO_BEAN_INITIALIZING = "Initializing mongo client."
const _MSG_MONGO_BEAN_INITIALIZED = "Mongo client initialized"
const _MSG_ERROR_TO_CONNECT_TO_DATABASE = "Application has been failed to connect to mongo database. URI: %s"

var once sync.Once
var mongoClient *mongo.Client = nil

func GetMongoDBClientBean() *mongo.Client {
	once.Do(func() {
		if mongoClient == nil {
			mongoClient = getClient()
		}
	})
	return mongoClient
}

func CloseConnection() {
	slog.Info(_MSG_MONGO_BEAN_CLOSING_CONNECTION)
	err := mongoClient.Disconnect(context.TODO())
	if err != nil {
		main_configs_error.FailOnError(err, _MSG_ERROR_TO_CLOSE_MONGO_CONNECTION)
	}
}

// IF Connection failed, how to solve
func getClient() *mongo.Client {
	slog.Info(_MSG_MONGO_BEAN_INITIALIZING)
	databaseUri := main_configs_yml.GetYmlValueByName(main_configs_yml.MongoDBURI)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf(_MSG_ERROR_TO_CONNECT_TO_DATABASE, databaseUri)
	}
	slog.Info(_MSG_MONGO_BEAN_INITIALIZED)
	return client
}
