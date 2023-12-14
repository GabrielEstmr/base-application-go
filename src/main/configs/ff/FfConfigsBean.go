package main_configs_ff

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_ff_lib "baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	"sync"
)

const _FF_FEATURES_NAME = "ff-features"

var once sync.Once
var ffConfigData *main_configs_ff_lib.FfConfig = nil

func GetFfConfigDataBean() *main_configs_ff_lib.FfConfig {
	once.Do(func() {
		if ffConfigData == nil {
			ffConfigData = getFfConfigData()
		}
	})
	return ffConfigData
}

func getFfConfigData() *main_configs_ff_lib.FfConfig {
	configData := main_configs_ff_lib.NewFfConfigDataBean(
		main_configs_mongo.GetMongoDBClient(),
		main_configs_ff_lib.MONGO,
		true,
		main_configs_cache.GetRedisClusterBean(),
		"PREFIX",
		main_configs_ff_lib.REDIS,
		_FF_FEATURES_NAME,
	)

	features := map[string]main_configs_ff_lib_resources.FeaturesData{
		"key1": *main_configs_ff_lib_resources.NewFeaturesData("key1",
			"key1",
			"key1",
			false),
	}

	registerImpl, err := main_configs_ff_lib.NewRegisterMethodsFactory(configData).Get()
	main_configs_error.FailOnError(err, "_MSG_ERROR_TRACER_RESOURCE")

	errRegister := registerImpl.RegisterFeatures(features)
	main_configs_error.FailOnError(errRegister, "_MSG_ERROR_TRACER_RESOURCE")

	featureMethodsImpl, errFm := main_configs_ff_lib.NewFeaturesMethodsFactory(configData).Get()
	main_configs_error.FailOnError(errFm, "_MSG_ERROR_TRACER_RESOURCE")

	return main_configs_ff_lib.NewFfConfig(configData, featureMethodsImpl, registerImpl)
}
