package main_domains_enums

type AuthProviderType string

const (
	AUTH_PROVIDER_KEYCLOAK  AuthProviderType = "KEYCLOAK"
	AUTH_PROVIDER_GOOGLE    AuthProviderType = "GOOGLE"
	AUTH_PROVIDER_MICROSOFT AuthProviderType = "MICROSOFT"
)

var authProviderTypeEnum = map[AuthProviderType]AuthProviderType{
	AUTH_PROVIDER_KEYCLOAK:  AUTH_PROVIDER_KEYCLOAK,
	AUTH_PROVIDER_GOOGLE:    AUTH_PROVIDER_GOOGLE,
	AUTH_PROVIDER_MICROSOFT: AUTH_PROVIDER_MICROSOFT,
}

var authProviderTypeEnumFromNames = map[string]AuthProviderType{
	"KEYCLOAK":  AUTH_PROVIDER_KEYCLOAK,
	"GOOGLE":    AUTH_PROVIDER_GOOGLE,
	"MICROSOFT": AUTH_PROVIDER_MICROSOFT,
}

var authProviderTypeDescriptionEnum = map[AuthProviderType]string{
	AUTH_PROVIDER_KEYCLOAK:  "keycloak",
	AUTH_PROVIDER_GOOGLE:    "google",
	AUTH_PROVIDER_MICROSOFT: "microsoft",
}

func (this AuthProviderType) Exists() bool {
	_, exists := authProviderTypeEnum[this]
	return exists
}

func (this AuthProviderType) FromValue(value string) AuthProviderType {
	valueMap, exists := authProviderTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this AuthProviderType) Name() string {
	valueMap, exists := authProviderTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

func (this AuthProviderType) GetProviderDescription() string {
	valueMap, exists := authProviderTypeDescriptionEnum[this]
	if exists {
		return valueMap
	}
	return ""
}
