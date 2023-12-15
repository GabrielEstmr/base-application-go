package main_configs_ff_lib_factories

import (
	"baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_mongo "baseapplicationgo/main/configs/ff/lib/mongo"
	"errors"
)

type FeaturesMethodsFactory struct {
	ffConfigData *main_configs_ff_lib.FfConfigData
}

func NewFeaturesMethodsFactory(ffConfigData *main_configs_ff_lib.FfConfigData) *FeaturesMethodsFactory {
	return &FeaturesMethodsFactory{ffConfigData}
}

func (this *FeaturesMethodsFactory) Get() (main_configs_ff_lib.FeaturesMethods, error) {
	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO && this.ffConfigData.GetHasCaching() == false {
		return main_configs_ff_lib_mongo.NewFeaturesMongoMethodsImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO && this.ffConfigData.GetHasCaching() == true && this.ffConfigData.GetCacheClientType() == main_configs_ff_lib.REDIS {
		return main_configs_ff_lib_mongo.NewFeaturesCachedMongoMethodsImpl(this.ffConfigData), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesData")
}
