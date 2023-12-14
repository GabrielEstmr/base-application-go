package main_configs_ff_lib

import (
	"errors"
)

type FeaturesMethodsFactory struct {
	ffConfigData *FfConfigData
}

func NewFeaturesMethodsFactory(ffConfigData *FfConfigData) *FeaturesMethodsFactory {
	return &FeaturesMethodsFactory{ffConfigData}
}

func (this *FeaturesMethodsFactory) Get() (FeaturesMethods, error) {
	if this.ffConfigData.GetClientType() == MONGO {
		return NewFeaturesMongoMethodsImpl(this.ffConfigData), nil
	}
	return nil, errors.New("could not instantiate a valid FeaturesData")
}
