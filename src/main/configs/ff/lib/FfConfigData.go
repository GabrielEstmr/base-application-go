package main_configs_ff_lib

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type FfConfigData struct {
	db              *mongo.Database
	clientType      string
	featuresColName string
	registerMethods RegisterMethods
	featuresMethods FeaturesMethods
}

func NewMongoFfConfigData(
	client *mongo.Database,
	clientType string,
	featuresColName string) *FfConfigData {
	return &FfConfigData{
		db:              client,
		clientType:      clientType,
		featuresColName: featuresColName,
		registerMethods: NewRegisterMethodsMongoImpl(),
		featuresMethods: NewFeaturesMongoMethodsImpl(),
	}
}

func (this *FfConfigData) RegisterFeatures() *mongo.Database {
	return this.db
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

func (this *FfConfigData) GetRegisterMethods() RegisterMethods {
	return this.registerMethods
}

func (this *FfConfigData) GetFeaturesMethods() FeaturesMethods {
	return this.featuresMethods
}
