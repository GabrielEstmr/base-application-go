package main_configs_yml

import (
	"fmt"
	"os"
	"strings"
)

const _MSG_ERROR_MANDATORY_ENV = "Mandatory env variable not found: %s"
const _MSG_ERROR_MANDATORY_SEPARATOR = "Mandatory separator not found: %s"

const _IDX_START_ENV_SEPARATOR = "${"
const _IDX_END_ENV_SEPARATOR = "}"

func ReplaceEnvNameToValue(value string) (string, bool) {
	hasIdx := hasEnvToSubstitute(value)
	if hasIdx {
		return substituteEnvToValue(value), hasIdx
	}
	return value, hasIdx
}

func hasEnvToSubstitute(value string) bool {
	idxStartEnv := strings.Index(value, _IDX_START_ENV_SEPARATOR)
	idxEndEnv := strings.Index(value, _IDX_END_ENV_SEPARATOR)
	return idxStartEnv != -1 && idxEndEnv != -1
}

func substituteEnvToValue(value string) string {
	before, afterTemp := cutOrPanic(value, _IDX_START_ENV_SEPARATOR)
	envName, after := cutOrPanic(afterTemp, _IDX_END_ENV_SEPARATOR)
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
		panic(fmt.Sprintf(_MSG_ERROR_MANDATORY_SEPARATOR, value))
	}
	return before, after
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic(fmt.Sprintf(_MSG_ERROR_MANDATORY_ENV + env))
	}
	return res
}
