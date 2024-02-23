package main_domains_enums

type UserEmailVerificationScope string

const (
	EMAIL_VERIFICATION_SCOPE_ENABLE_USER     UserEmailVerificationScope = "ENABLE_USER"
	EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD UserEmailVerificationScope = "CHANGE_PASSWORD"
)

var userEmailVerificationScopeEnum = map[UserEmailVerificationScope]UserEmailVerificationScope{
	EMAIL_VERIFICATION_SCOPE_ENABLE_USER:     EMAIL_VERIFICATION_SCOPE_ENABLE_USER,
	EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD: EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD,
}

var userEmailVerificationScopeEnumFromNames = map[string]UserEmailVerificationScope{
	"ENABLE_USER":     EMAIL_VERIFICATION_SCOPE_ENABLE_USER,
	"CHANGE_PASSWORD": EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD,
}

func (this UserEmailVerificationScope) Exists() bool {
	_, exists := userEmailVerificationScopeEnum[this]
	return exists
}

func (this UserEmailVerificationScope) FromValue(value string) UserEmailVerificationScope {
	valueMap, exists := userEmailVerificationScopeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this UserEmailVerificationScope) Name() string {
	valueMap, exists := userEmailVerificationScopeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
