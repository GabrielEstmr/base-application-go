package main_configs_messages

import (
	main_error "baseapplicationgo/main/configs/error"
	main_configs_messages_resources "baseapplicationgo/main/configs/messages/resources"
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const _MSG_INITIALIZING_MSG_BEANS = "Initializing Message configuration beans"
const _MSG_MSG_BEANS_INITIATED = "Message configuration beans successfully initiated"
const _MSG_ERROR_READ_MSG = "Error to read message file."

const _MSG_BASE_DIRECTORY_MAIN_REFERENCE = "zresources/messages"
const _MSG_FILE_DEFAULT_BASE_NAME = "/messages-%s.properties"

const _DEFAULT_KEY_SEPARATOR = "="

var once sync.Once
var msgConfigsBean *map[string]string

func GetMessagesConfigBean() *map[string]string {
	once.Do(func() {
		if msgConfigsBean == nil {
			msgConfigsBean = getMessagesConfig()
		}
	})
	return msgConfigsBean
}

func getMessagesConfig() *map[string]string {
	slog.Info(_MSG_INITIALIZING_MSG_BEANS)
	var config = make(map[string]string)
	for _, langEnum := range main_configs_messages_resources.GetLanguageProfileValues() {
		lang := langEnum.GetLanguageProfileDescription()

		_, b, _, _ := runtime.Caller(0)
		basePath := strings.Replace(
			filepath.Dir(b), "main/configs/messages", "", 1) + _MSG_BASE_DIRECTORY_MAIN_REFERENCE

		msgFilePath := basePath + getFileName(lang)
		file, err := os.Open(msgFilePath)
		main_error.FailOnError(err, _MSG_ERROR_READ_MSG)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if sepIdx := getSeparatorIndex(line); sepIdx >= 0 {
				if key := getPropertyKey(langEnum.LanguageProfileName(), line, sepIdx); len(key) > 0 {
					config[key] = getValueIfNotEmpty(line, sepIdx)
				}
			}
		}
	}
	slog.Info(_MSG_MSG_BEANS_INITIATED)
	return &config
}

func getSeparatorIndex(line string) int {
	return strings.Index(line, _DEFAULT_KEY_SEPARATOR)
}

func getPropertyKey(lang, line string, separatorIndex int) string {
	return strings.TrimSpace(line[:separatorIndex]) + "-" + lang
}

func getValueIfNotEmpty(line string, separatorIndex int) string {
	value := ""
	if len(line) > separatorIndex {
		value = strings.TrimSpace(line[separatorIndex+1:])
	}
	return value
}

func getFileName(lang string) string {
	return fmt.Sprintf(_MSG_FILE_DEFAULT_BASE_NAME, lang)
}
