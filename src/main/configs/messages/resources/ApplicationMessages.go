package main_configs_messages_resources

type ApplicationMessages struct {
	properties map[string]string
}

func NewApplicationMessages(properties map[string]string) *ApplicationMessages {
	return &ApplicationMessages{properties}
}

func (this *ApplicationMessages) GetProperties() map[string]string {
	return this.properties
}

func (this *ApplicationMessages) SetProperties(properties map[string]string) {
	this.properties = properties
}

func (this *ApplicationMessages) GetDefaultLocale(key string) string {
	return this.GetMessageByLocale(key, DEFAULT)
}

func (this *ApplicationMessages) GetMessageByLocale(key string,
	langProfile LanguageProfiles) string {
	profileDescription := langProfile.LanguageProfileName()
	langKey := key + "-" + profileDescription
	return this.properties[langKey]
}
