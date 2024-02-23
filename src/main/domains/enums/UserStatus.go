package main_domains_enums

type UserStatus string

const (
	USER_CREATED  UserStatus = "USER_CREATED"
	USER_ENABLED  UserStatus = "USER_ENABLED"
	USER_DISABLED UserStatus = "USER_DISABLED"
)

var userStatusEnum = map[UserStatus]UserStatus{
	USER_CREATED:  USER_CREATED,
	USER_ENABLED:  USER_ENABLED,
	USER_DISABLED: USER_DISABLED,
}

var userStatusEnumFromNames = map[string]UserStatus{
	"USER_CREATED":  USER_CREATED,
	"USER_ENABLED":  USER_ENABLED,
	"USER_DISABLED": USER_DISABLED,
}

func (this UserStatus) Exists() bool {
	_, exists := userStatusEnum[this]
	return exists
}

func (this UserStatus) FromValue(value string) UserStatus {
	valueMap, exists := userStatusEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this UserStatus) Name() string {
	valueMap, exists := userStatusEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
