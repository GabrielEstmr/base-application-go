package main_configs_ff_lib

import main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"

type RegisterMethods interface {
	RegisterFeatures(features map[string]main_configs_ff_lib_resources.FeaturesData) error
}
