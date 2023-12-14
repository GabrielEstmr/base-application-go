package main_configs_ff

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_ff_lib "baseapplicationgo/main/configs/ff/lib"
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	"sync"
)

const _FF_FEATURES_NAME = "ff-features"

var once sync.Once
var ffConfigData *main_configs_ff_lib.FfConfigData = nil

func GetFfConfigDataBean() *main_configs_ff_lib.FfConfigData {
	once.Do(func() {
		if ffConfigData == nil {
			ffConfigData = getFfConfigData()
		}
	})
	return ffConfigData
}

func getFfConfigData() *main_configs_ff_lib.FfConfigData {
	bean := main_configs_ff_lib.NewFfConfigDataBean(
		main_configs_mongo.GetMongoDBClient(),
		main_configs_ff_lib.MONGO,
		_FF_FEATURES_NAME,
	)

	temp := map[string]main_configs_ff_lib_resources.FeaturesData{
		"key1": *main_configs_ff_lib_resources.NewFeaturesData("key1",
			"key1",
			"key1",
			false),
	}

	err := bean.GetRegisterMethods().RegisterFeatures(temp)
	main_configs_error.FailOnError(err, "_MSG_ERROR_TRACER_RESOURCE")
	return bean
}
