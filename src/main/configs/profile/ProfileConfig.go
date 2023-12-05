package main_configs_profile

import (
	configsEnv "baseapplicationgo/main/configs/env"
	main_error "baseapplicationgo/main/configs/error"
	"log/slog"
	"sync"
)

const _MSG_INITIALIZING_PROFILE_BEANS = "Initializing Profile configuration beans"
const _MSG_PROFILE_BEANS_INITIATED = "Profile configuration beans successfully initiated"
const _MSG_ERROR_TO_GET_PROFILE = "Error to get application profile"

var once sync.Once
var ApplicationProfileBean *ApplicationProfile

func GetProfileBean() *ApplicationProfile {
	once.Do(func() {
		if ApplicationProfileBean == nil {
			ApplicationProfileBean = getProfile()
		}

	})
	return ApplicationProfileBean
}

func getProfile() *ApplicationProfile {
	slog.Info(_MSG_INITIALIZING_PROFILE_BEANS)
	profile := configsEnv.GetBeanPropertyByName(configsEnv.MP_INDICATOR_APPLICATION_PROFILE)
	appProfile, err := FindApplicationProfileByDescription(profile)
	main_error.FailOnError(err, _MSG_ERROR_TO_GET_PROFILE)
	slog.Info(_MSG_PROFILE_BEANS_INITIATED)
	return &appProfile
}
