package main_domains_enums

type AccountStatus string

const (
	ACCOUNT_CREATED  AccountStatus = "ACCOUNT_CREATED"
	ACCOUNT_ENABLED  AccountStatus = "ACCOUNT_ENABLED"
	ACCOUNT_DISABLED AccountStatus = "ACCOUNT_DISABLED"
)

var accountStatusEnum = map[AccountStatus]AccountStatus{
	ACCOUNT_CREATED:  ACCOUNT_CREATED,
	ACCOUNT_ENABLED:  ACCOUNT_ENABLED,
	ACCOUNT_DISABLED: ACCOUNT_DISABLED,
}

var accountStatusEnumFromNames = map[string]AccountStatus{
	"ACCOUNT_CREATED":  ACCOUNT_CREATED,
	"ACCOUNT_ENABLED":  ACCOUNT_ENABLED,
	"ACCOUNT_DISABLED": ACCOUNT_DISABLED,
}

func (this AccountStatus) Exists() bool {
	_, exists := accountStatusEnum[this]
	return exists
}

func (this AccountStatus) FromValue(value string) AccountStatus {
	valueMap, exists := accountStatusEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this AccountStatus) Name() string {
	valueMap, exists := accountStatusEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
