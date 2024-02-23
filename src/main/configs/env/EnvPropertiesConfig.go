package main_configs_env

import (
	main_error "baseapplicationgo/main/configs/error"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

const _MSG_ENV_BEAN = "invalid env property value"
const _MSG_ERROR_READ_ENV_FILE = "error to read .env file"

const YML_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources"

var once sync.Once
var envValues *map[string]string

func GetEnvConfigBean() *map[string]string {
	once.Do(func() {
		if envValues == nil {
			envValues = getEnvConfig()
		}
	})
	return envValues
}

func getEnvConfig() *map[string]string {
	err := godotenv.Load(YML_BASE_DIRECTORY_MAIN_REFERENCE + "/.env")
	main_error.FailOnError(err, _MSG_ERROR_READ_ENV_FILE)

	data := make(map[string]string)
	for _, value := range new(EnvironmentProperty).Values() {
		envValue := os.Getenv(value.Name())
		if envValue == "" {
			log.Panicf("%s: %s", _MSG_ENV_BEAN, errors.New(_MSG_ENV_BEAN))
		}
		data[value.Name()] = envValue
	}
	return &data
}
