package main_utils_messages

import (
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains_enums "baseapplicationgo/main/domains/enums"
)

type ApplicationMessages struct {
	properties map[string]string
}

func NewApplicationMessagesAllArgs(properties map[string]string) *ApplicationMessages {
	return &ApplicationMessages{properties: properties}
}

func NewApplicationMessages() *ApplicationMessages {
	return &ApplicationMessages{*main_configs_messages.GetMessagesConfigBean()}
}

func (this *ApplicationMessages) GetProperties() map[string]string {
	return this.properties
}

func (this *ApplicationMessages) SetProperties(properties map[string]string) {
	this.properties = properties
}

func (this *ApplicationMessages) GetDefaultLocale(key string) string {
	return this.GetMessageByLocale(key, main_domains_enums.LANGUAGE_PROFILE_DEFAULT)
}

func (this *ApplicationMessages) GetMessageByLocale(key string, langProfile main_domains_enums.LanguageProfiles) string {
	langKey := key + "-" + langProfile.Name()
	return this.properties[langKey]
}
