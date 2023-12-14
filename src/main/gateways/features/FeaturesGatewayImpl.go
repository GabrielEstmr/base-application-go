package main_gateways_features

import (
	main_configs_ff "baseapplicationgo/main/configs/ff"
	main_configs_ff_lib "baseapplicationgo/main/configs/ff/lib"
)

type FeaturesGatewayImpl struct {
	ffConfig *main_configs_ff_lib.FfConfig
}

func NewFeaturesGatewayImpl() *FeaturesGatewayImpl {
	return &FeaturesGatewayImpl{
		ffConfig: main_configs_ff.GetFfConfigDataBean(),
	}
}

func (this *FeaturesGatewayImpl) IsEnabled(key string) (bool, error) {
	return this.ffConfig.GetFeaturesMethods().IsEnabled(key)
}

func (this *FeaturesGatewayImpl) IsDisabled(key string) (bool, error) {
	return this.ffConfig.GetFeaturesMethods().IsDisabled(key)
}
