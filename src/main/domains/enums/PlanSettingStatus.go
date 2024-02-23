package main_domains_enums

type PlanSettingStatus string

const (
	PLAN_SETTING_ENABLED  PlanSettingStatus = "PLAN_SETTING_ENABLED"
	PLAN_SETTING_DISABLED PlanSettingStatus = "PLAN_SETTING_DISABLED"
)

var planSettingStatusEnum = map[PlanSettingStatus]PlanSettingStatus{
	PLAN_SETTING_ENABLED:  PLAN_SETTING_ENABLED,
	PLAN_SETTING_DISABLED: PLAN_SETTING_DISABLED,
}

var planSettingStatusEnumFromNames = map[string]PlanSettingStatus{
	"PLAN_SETTING_ENABLED":  PLAN_SETTING_ENABLED,
	"PLAN_SETTING_DISABLED": PLAN_SETTING_DISABLED,
}

func (this PlanSettingStatus) Exists() bool {
	_, exists := planSettingStatusEnum[this]
	return exists
}

func (this PlanSettingStatus) FromValue(value string) PlanSettingStatus {
	valueMap, exists := planSettingStatusEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this PlanSettingStatus) Name() string {
	valueMap, exists := planSettingStatusEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
