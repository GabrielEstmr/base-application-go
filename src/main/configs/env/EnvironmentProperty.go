package main_configs_env

type EnvironmentProperty string

const (
	APPLICATION_PROFILE EnvironmentProperty = "APPLICATION_PROFILE"
	//RM_USER_KEYCLOAK_CLIENT_ID EnvironmentProperty = "RM_USER_KEYCLOAK_CLIENT_ID"
	//RM_USER_KEYCLOAK_SECRET    EnvironmentProperty = "RM_USER_KEYCLOAK_SECRET"
	//RM_USER_KEYCLOAK_REALM     EnvironmentProperty = "RM_USER_KEYCLOAK_REALM"
)

var environmentPropertyEnum = map[EnvironmentProperty]EnvironmentProperty{
	APPLICATION_PROFILE: APPLICATION_PROFILE,
	//RM_USER_KEYCLOAK_CLIENT_ID: RM_USER_KEYCLOAK_CLIENT_ID,
	//RM_USER_KEYCLOAK_SECRET:    RM_USER_KEYCLOAK_SECRET,
	//RM_USER_KEYCLOAK_REALM:     RM_USER_KEYCLOAK_REALM,
}

var environmentPropertyEnumFromNames = map[string]EnvironmentProperty{
	"APPLICATION_PROFILE": APPLICATION_PROFILE,
	//"RM_USER_KEYCLOAK_CLIENT_ID": RM_USER_KEYCLOAK_CLIENT_ID,
	//"RM_USER_KEYCLOAK_SECRET":    RM_USER_KEYCLOAK_SECRET,
	//"RM_USER_KEYCLOAK_REALM":     RM_USER_KEYCLOAK_REALM,
}

func (this EnvironmentProperty) Exists() bool {
	_, exists := environmentPropertyEnum[this]
	return exists
}

func (this EnvironmentProperty) FromValue(value string) EnvironmentProperty {
	valueMap, exists := environmentPropertyEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this EnvironmentProperty) Name() string {
	valueMap, exists := environmentPropertyEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

func (this EnvironmentProperty) Values() []EnvironmentProperty {
	values := make([]EnvironmentProperty, 0)
	for _, v := range environmentPropertyEnum {
		values = append(values, v)
	}
	return values
}
