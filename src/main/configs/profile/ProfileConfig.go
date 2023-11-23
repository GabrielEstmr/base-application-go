package main_configs_profile

import (
	configsEnv "baseapplicationgo/main/configs/env"
	utils "baseapplicationgo/main/utils"
	"sync"
)

const MSG_ERROR_TO_GET_PROFILE = "Error to get application profile"

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
	profile := configsEnv.GetBeanPropertyByName(configsEnv.MP_INDICATOR_APPLICATION_PROFILE)
	appProfile, err := FindApplicationProfileByDescription(profile)
	utils.FailOnError(err, MSG_ERROR_TO_GET_PROFILE)
	return &appProfile
}
