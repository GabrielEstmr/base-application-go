package mainConfigsProfile

import (
	mainConfigsEnv "baseapplicationgo/main/configurations/env"
	mainUtils "baseapplicationgo/main/utils"
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
	profile := mainConfigsEnv.GetBeanPropertyByName(mainConfigsEnv.MP_INDICATOR_APPLICATION_PROFILE)
	appProfile, err := FindApplicationProfileByDescription(profile)
	mainUtils.FailOnError(err, MSG_ERROR_TO_GET_PROFILE)
	return &appProfile
}
