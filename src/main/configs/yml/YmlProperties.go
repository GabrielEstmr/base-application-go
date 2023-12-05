package main_configs_yml

import (
	main_error "baseapplicationgo/main/configs/error"
	main_configs_profile "baseapplicationgo/main/configs/profile"
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
var ymlConfigsBean *map[string]property

type property struct {
	Value string
}

func GetYmlConfigBean() *map[string]property {
	once.Do(func() {
		if ymlConfigsBean == nil {
			ymlConfigsBean = getYmlConfig()
		}
	})
	return ymlConfigsBean
}

func getYmlConfig() *map[string]property {

	slog.Info(_MSG_INITIALIZING_YML_BEANS)
	profile := main_configs_profile.GetProfileBean().GetLowerCaseName()
	ymlPath := _YML_BASE_DIRECTORY_MAIN_REFERENCE + fmt.Sprintf(
		_YML_FILE_DEFAULT_BASE_NAME, profile)

	yFile, err := os.ReadFile(ymlPath)
	main_error.FailOnError(err, _MSG_ERROR_READ_YML)

	data := make(map[string]property)
	err2 := yaml.Unmarshal(yFile, &data)
	main_error.FailOnError(err2, _MSG_ERROR_PARSE_YML)

	for key, _ := range data {
		for {
			newValue, hasIdx := ReplaceEnvNameToValue(data[key].Value)
			data[key] = property{newValue}
			if !hasIdx {
				break
			}
		}

	}
	slog.Info(_MSG_YML_BEANS_INITIATED)
	return &data
}
