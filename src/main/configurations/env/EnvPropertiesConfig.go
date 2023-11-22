package main_configurations_env

import (
	main_utils "baseapplicationgo/main/utils"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

const MSG_ENV_BEAN = "Invalid env property value."
const MSG_ERROR_READ_ENV_FILE = "Error to read .env file."

const YML_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources"

var once sync.Once
var EnvValues *map[string]string

func GetEnvConfigBean() *map[string]string {
	once.Do(func() {
		if EnvValues == nil {
			EnvValues = getEnvConfig()
		}
	})
	return EnvValues
}

func getEnvConfig() *map[string]string {

	envNames := []string{
		MP_INDICATOR_APPLICATION_PROFILE.GetDescription(),
	}

	err := godotenv.Load(YML_BASE_DIRECTORY_MAIN_REFERENCE + "/.env")
	main_utils.FailOnError(err, MSG_ERROR_READ_ENV_FILE)

	data := make(map[string]string)
	for _, value := range envNames {
		envValue := os.Getenv(value)
		if envValue == "" {
			log.Panicf("%s: %s", MSG_ENV_BEAN, errors.New(MSG_ENV_BEAN))
		}
		data[value] = envValue
	}
	return &data
}
