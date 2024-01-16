package main_configs_ff_lib_factories

import (
	"baseapplicationgo/main/configs/ff/lib"
	"baseapplicationgo/main/configs/ff/lib/mongo"
	"errors"
)

type RegisterMethodsFactory struct {
	ffConfigData *main_configs_ff_lib.FfConfigData
}

func NewRegisterMethodsFactory(ffConfigData *main_configs_ff_lib.FfConfigData) *RegisterMethodsFactory {
	return &RegisterMethodsFactory{ffConfigData}
}

func (this *RegisterMethodsFactory) Get() (main_configs_ff_lib.RegisterMethods, error) {
	if this.ffConfigData.GetClientType() == main_configs_ff_lib.MONGO {
		return main_configs_ff_lib_mongo.NewRegisterMethodsMongoImpl(this.ffConfigData), nil
	}
	return nil, errors.New("could not instantiate a valid FeaturesData")
}
