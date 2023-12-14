package main_configs_ff_lib

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type FfConfigData struct {
	db              *mongo.Database
	clientType      string
	featuresColName string
}

func NewMongoFfConfigData(
	client *mongo.Database,
	clientType string,
	featuresColName string) *FfConfigData {
	return &FfConfigData{
		db:              client,
		clientType:      clientType,
		featuresColName: featuresColName,
	}
}

func (this *FfConfigData) GetDb() *mongo.Database {
	return this.db
}

func (this *FfConfigData) GetClientType() string {
	return this.clientType
}

func (this *FfConfigData) GetFeaturesColName() string {
	return this.featuresColName
}
