package main_configs_messages

import (
	main_configs_messages_resources "baseapplicationgo/main/configs/messages/resources"
	main_utils "baseapplicationgo/main/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const _MSG_INITIALIZING_MSG_BEANS = "Initializing Message configuration beans"
const _MSG_MSG_BEANS_INITIATED = "Message configuration beans successfully initiated"
const _MSG_ERROR_READ_MSG = "Error to read message file."

const _MSG_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources/messages"
const _MSG_FILE_DEFAULT_BASE_NAME = "/messages-%s.properties"

const _DEFAULT_KEY_SEPARATOR = "="

var once sync.Once
var msgConfigsBean *main_configs_messages_resources.ApplicationMessages

func GetMessagesConfigBean() *main_configs_messages_resources.ApplicationMessages {
	once.Do(func() {
		if msgConfigsBean == nil {
			msgConfigsBean = getMessagesConfig()
		}
	})
	return msgConfigsBean
}

func getMessagesConfig() *main_configs_messages_resources.ApplicationMessages {
	log.Println(_MSG_INITIALIZING_MSG_BEANS)
	var config = make(map[string]string)
	for _, langEnum := range main_configs_messages_resources.GetLanguageProfileValues() {
		lang := langEnum.GetLanguageProfileDescription()
		msgFilePath := _MSG_BASE_DIRECTORY_MAIN_REFERENCE + getFileName(lang)
		file, err := os.Open(msgFilePath)
		main_utils.FailOnError(err, _MSG_ERROR_READ_MSG)
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
	log.Println(_MSG_MSG_BEANS_INITIATED)
	return main_configs_messages_resources.NewApplicationMessages(config)
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
