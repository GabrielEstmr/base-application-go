package main_gateways

import main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"

type FeaturesGateway interface {
	IsEnabled(key string) (bool, error)
	IsDisabled(key string) (bool, error)
	Enable(key string) (main_configs_ff_lib_resources.FeaturesData, error)
	Disable(key string) (main_configs_ff_lib_resources.FeaturesData, error)
}
