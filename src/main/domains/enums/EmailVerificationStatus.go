package main_domains_enums

type EmailVerificationStatus string

const (
	EMAIL_VERIFICATION_CREATED  EmailVerificationStatus = "EMAIL_VERIFICATION_CREATED"
	EMAIL_VERIFICATION_SENT     EmailVerificationStatus = "EMAIL_VERIFICATION_SENT"
	EMAIL_VERIFICATION_DISABLED EmailVerificationStatus = "EMAIL_VERIFICATION_DISABLED"
	EMAIL_VERIFICATION_USED     EmailVerificationStatus = "EMAIL_VERIFICATION_USED"
)

var emailVerificationStatusEnum = map[EmailVerificationStatus]EmailVerificationStatus{
	EMAIL_VERIFICATION_CREATED:  EMAIL_VERIFICATION_CREATED,
	EMAIL_VERIFICATION_SENT:     EMAIL_VERIFICATION_SENT,
	EMAIL_VERIFICATION_DISABLED: EMAIL_VERIFICATION_DISABLED,
	EMAIL_VERIFICATION_USED:     EMAIL_VERIFICATION_USED,
}

var accountStatusEnumFromNamesEnumFromNames = map[string]EmailVerificationStatus{
	"EMAIL_VERIFICATION_CREATED":  EMAIL_VERIFICATION_CREATED,
	"EMAIL_VERIFICATION_SENT":     EMAIL_VERIFICATION_SENT,
	"EMAIL_VERIFICATION_DISABLED": EMAIL_VERIFICATION_DISABLED,
	"EMAIL_VERIFICATION_USED":     EMAIL_VERIFICATION_USED,
}

func (this EmailVerificationStatus) Exists() bool {
	_, exists := emailVerificationStatusEnum[this]
	return exists
}

func (this EmailVerificationStatus) FromValue(value string) EmailVerificationStatus {
	valueMap, exists := accountStatusEnumFromNamesEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this EmailVerificationStatus) Name() string {
	valueMap, exists := emailVerificationStatusEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
