package main_configs_ff

import (
	main_configs_cache "baseapplicationgo/main/configs/cache"
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_ff_lib "baseapplicationgo/main/configs/ff/lib"
	"baseapplicationgo/main/configs/ff/lib/factories"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log"
	"sync"
)

const _MSG_INITIALIZING_FEATURES_BEANS = "Initializing features configuration bean"
const _MSG_FEATURES_BEANS_INITIATED = "Application features bean successfully initiated"
const _MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY = "Error to get features register provider"
const _MSG_ERROR_INSTANTIATE_FEATURES_METHODS_FACTORY = "Error to get features method provider"
const _MSG_ERROR_REGISTER_FEATURES = "Error to register features"

const _FF_FEATURES_DB_NAME = "ff-features"
const _APP_NAME_YML = "Apm.server.name"

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

	log.Println(_MSG_INITIALIZING_FEATURES_BEANS)

	configData := main_configs_ff_lib.NewFfConfigDataBean(
		main_configs_mongo.GetMongoDBClient(),
		main_configs_ff_lib.MONGO,
		true,
		main_configs_cache.GetRedisClusterBean(),
		main_configs_yml.GetYmlValueByName(_APP_NAME_YML),
		main_configs_ff_lib.REDIS,
		_FF_FEATURES_DB_NAME,
	)

	registerImpl, err := main_configs_ff_lib_factories.NewRegisterMethodsFactory(configData).Get()
	main_configs_error.FailOnError(err, _MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY)

	errRegister := registerImpl.RegisterFeatures(FEATURES)
	main_configs_error.FailOnError(errRegister, _MSG_ERROR_REGISTER_FEATURES)

	featureMethodsImpl, errFm := main_configs_ff_lib_factories.NewFeaturesMethodsFactory(configData).Get()
	main_configs_error.FailOnError(errFm, _MSG_ERROR_INSTANTIATE_FEATURES_METHODS_FACTORY)

	log.Println(_MSG_FEATURES_BEANS_INITIATED)
	return main_configs_ff_lib.NewFfConfig(configData, featureMethodsImpl, registerImpl)
}
