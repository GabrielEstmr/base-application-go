package main_domains_enums

type LockScope string

const (
	LOCK_SCOPE_USER_MODIFICATION                    LockScope = "USER_MODIFICATION"
	LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION LockScope = "USER_VERIFICATION_EMAIL_MODIFICATION"
)

var lockScopeEnum = map[LockScope]LockScope{
	LOCK_SCOPE_USER_MODIFICATION:                    LOCK_SCOPE_USER_MODIFICATION,
	LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION: LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION,
}

var lockScopeEnumFromNames = map[string]LockScope{
	"USER_MODIFICATION":                    LOCK_SCOPE_USER_MODIFICATION,
	"USER_VERIFICATION_EMAIL_MODIFICATION": LOCK_SCOPE_USER_VERIFICATION_EMAIL_MODIFICATION,
}

func (this LockScope) Exists() bool {
	_, exists := lockScopeEnum[this]
	return exists
}

func (this LockScope) FromValue(value string) LockScope {
	valueMap, exists := lockScopeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this LockScope) Name() string {
	valueMap, exists := lockScopeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
