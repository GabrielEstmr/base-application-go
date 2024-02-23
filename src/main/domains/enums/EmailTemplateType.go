package main_domains_enums

type EmailTemplateType string

const (
	EMAIL_TYPE_WELCOME_EMAIL     EmailTemplateType = "EMAIL_TYPE_WELCOME_EMAIL"
	EMAIL_TYPE_NOTIFICATION_USER EmailTemplateType = "EMAIL_TYPE_NOTIFICATION_USER"
	EMAIL_TYPE_VERIFICATION_USER EmailTemplateType = "EMAIL_TYPE_VERIFICATION_USER"
)

var emailTemplateTypeEnum = map[EmailTemplateType]EmailTemplateType{
	EMAIL_TYPE_WELCOME_EMAIL:     EMAIL_TYPE_WELCOME_EMAIL,
	EMAIL_TYPE_NOTIFICATION_USER: EMAIL_TYPE_NOTIFICATION_USER,
	EMAIL_TYPE_VERIFICATION_USER: EMAIL_TYPE_VERIFICATION_USER,
}

var emailTemplateTypeEnumFromNames = map[string]EmailTemplateType{
	"EMAIL_TYPE_WELCOME_EMAIL":     EMAIL_TYPE_WELCOME_EMAIL,
	"EMAIL_TYPE_NOTIFICATION_USER": EMAIL_TYPE_NOTIFICATION_USER,
	"EMAIL_TYPE_VERIFICATION_USER": EMAIL_TYPE_VERIFICATION_USER,
}

func (this EmailTemplateType) Exists() bool {
	_, exists := emailTemplateTypeEnum[this]
	return exists
}

func (this EmailTemplateType) FromValue(value string) EmailTemplateType {
	valueMap, exists := emailTemplateTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this EmailTemplateType) Name() string {
	valueMap, exists := emailTemplateTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
