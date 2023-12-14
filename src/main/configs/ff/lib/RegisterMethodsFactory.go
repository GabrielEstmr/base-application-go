package main_configs_ff_lib

import "errors"

type RegisterMethodsFactory struct {
	ffConfigData *FfConfigData
}

func NewRegisterMethodsFactory(ffConfigData *FfConfigData) *RegisterMethodsFactory {
	return &RegisterMethodsFactory{ffConfigData}
}

func (this *RegisterMethodsFactory) Get() (RegisterMethods, error) {
	if this.ffConfigData.GetClientType() == MONGO {
		return NewRegisterMethodsMongoImpl(this.ffConfigData), nil
	}
	return nil, errors.New("could not instantiate a valid FeaturesData")
}
