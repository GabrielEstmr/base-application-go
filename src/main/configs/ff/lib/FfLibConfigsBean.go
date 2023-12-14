package main_configs_ff_lib

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var once sync.Once
var ffLibConfigsBean *FfConfigData

func GetFfConfigDataBean() *FfConfigData {
	return ffLibConfigsBean
}

func NewFfConfigDataBean(client *mongo.Database,
	clientType string,
	featuresDbName string) *FfConfigData {
	once.Do(func() {
		if ffLibConfigsBean == nil {
			ffLibConfigsBean = getFfConfigData(client, clientType, featuresDbName)
		}
	})
	return ffLibConfigsBean
}

func getFfConfigData(client *mongo.Database,
	clientType string,
	featuresDbName string) *FfConfigData {
	return NewMongoFfConfigData(client, clientType, featuresDbName)
}
