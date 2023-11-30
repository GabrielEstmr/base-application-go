package main_configs_yml

import (
	main_configs_profile "baseapplicationgo/main/configs/profile"
	utils "baseapplicationgo/main/utils"
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"sync"
)

const _MSG_INITIALIZING_YML_BEANS = "Initializing Yml configuration beans"
const _MSG_YML_BEANS_INITIATED = "Yml configuration beans successfully initiated"
const _MSG_ERROR_READ_YML = "Error to read yml file."
const _MSG_ERROR_PARSE_YML = "Error to parse yml file."

const _YML_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources"
const _YML_FILE_DEFAULT_BASE_NAME = "/application-properties-%s.yaml"

var once sync.Once
var ymlConfigs *map[string]property
var ymlConfigsBean map[string]property

type property struct {
	Value string
}

func GetYmlConfigBean() *map[string]property {
	once.Do(func() {
		if ymlConfigs == nil {
			ymlConfigsBean = getYmlConfig()
			ymlConfigs = &ymlConfigsBean
		}

	})
	return ymlConfigs
}

func getYmlConfig() map[string]property {

	slog.Info(_MSG_INITIALIZING_YML_BEANS)
	profile := main_configs_profile.GetProfileBean().GetLowerCaseName()
	ymlPath := _YML_BASE_DIRECTORY_MAIN_REFERENCE + fmt.Sprintf(
		_YML_FILE_DEFAULT_BASE_NAME, profile)

	yFile, err := os.ReadFile(ymlPath)
	utils.FailOnError(err, _MSG_ERROR_READ_YML)

	data := make(map[string]property)
	err2 := yaml.Unmarshal(yFile, &data)
	utils.FailOnError(err2, _MSG_ERROR_PARSE_YML)

	for key, _ := range data {
		newValue := ReplaceEnvNameToValue(data[key].Value)
		data[key] = property{newValue}

	}
	slog.Info(_MSG_YML_BEANS_INITIATED)
	return data
}
