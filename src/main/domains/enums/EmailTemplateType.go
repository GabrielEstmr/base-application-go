package main_domains_enums

type EmailTemplateType string

const (
	EMAIL_TYPE_WELCOME_EMAIL     EmailTemplateType = "WELCOME_EMAIL"
	EMAIL_TYPE_NOTIFICATION_USER EmailTemplateType = "NOTIFICATION_USER"
)

var emailTemplateTypeEnum = map[EmailTemplateType]EmailTemplateType{
	EMAIL_TYPE_WELCOME_EMAIL:     EMAIL_TYPE_WELCOME_EMAIL,
	EMAIL_TYPE_NOTIFICATION_USER: EMAIL_TYPE_NOTIFICATION_USER,
}

func ExistsEmailTemplateType(value EmailTemplateType) bool {
	_, exists := emailTemplateTypeEnum[value]
	return exists
}

func GetEmailTemplateTypeDescription(value EmailTemplateType) string {
	valueMap, exists := emailTemplateTypeEnum[value]
	if exists {
		return string(valueMap)
	}
	return ""
}

func GetEmailTemplateTypeFromDescription(description string) EmailTemplateType {
	valueMap, exists := emailTemplateTypeEnum[EmailTemplateType(description)]
	if exists {
		return valueMap
	}
	return ""
}
