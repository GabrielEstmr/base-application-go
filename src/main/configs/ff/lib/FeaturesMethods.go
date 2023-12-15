package main_configs_ff_lib

import main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"

type FeaturesMethods interface {
	IsEnabled(key string) (bool, error)
	IsDisabled(key string) (bool, error)
	Enable(key string) (main_configs_ff_lib_resources.FeaturesData, error)
	Disable(key string) (main_configs_ff_lib_resources.FeaturesData, error)
}
