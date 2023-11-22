package main_configurations_yml

import (
	"fmt"
	"os"
	"strings"
)

const IDX_START_ENV_SEPARATOR = "${"
const IDX_END_ENV_SEPARATOR = "}"

func ReplaceEnvNameToItsValue(value string) string {
	if hasEnvToSubstitute(value) {
		substituteEnvToValue(value)
	}
	return value
}

func hasEnvToSubstitute(value string) bool {
	idxStartEnv := strings.Index(value, IDX_START_ENV_SEPARATOR)
	idxEndEnv := strings.Index(value, IDX_END_ENV_SEPARATOR)
	return idxStartEnv != -1 && idxEndEnv != -1
}

func substituteEnvToValue(value string) string {
	before, afterTemp := cutOrPanic(value, IDX_START_ENV_SEPARATOR)
	envName, after := cutOrPanic(afterTemp, IDX_END_ENV_SEPARATOR)
	envValue := getEnvOrPanic(envName)
	newValue := buildWithEnvValue(before, envValue, after)
	return newValue
}

func buildWithEnvValue(before string, envValue string, after string) string {
	return before + envValue + after
}

func cutOrPanic(value string, sep string) (string, string) {
	before, after, foundPrefix := strings.Cut(value, sep)
	if !foundPrefix {
		panic(fmt.Sprintf(MSG_ERROR_MANDATORY_SEPARATOR, value))
	}
	return before, after
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic(fmt.Sprintf(MSG_ERROR_MANDATORY_ENV + env))
	}
	return res
}
