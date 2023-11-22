package main_configurations_profile

import (
	main_configurations_env "baseapplicationgo/main/configurations/env"
	main_utils "baseapplicationgo/main/utils"
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
	profile := main_configurations_env.GetBeanPropertyByName(main_configurations_env.MP_INDICATOR_APPLICATION_PROFILE)
	appProfile, err := FindApplicationProfileByDescription(profile)
	main_utils.FailOnError(err, MSG_ERROR_TO_GET_PROFILE)
	return &appProfile
}
