package main_domains_enums

type PlanType string

const (
	FREE         PlanType = "FREE"
	BASIC        PlanType = "BASIC"
	INTERMEDIATE PlanType = "INTERMEDIATE"
	PLUS         PlanType = "PLUS"
	PRO          PlanType = "PRO"
)

var planTypeEnum = map[PlanType]PlanType{
	FREE:         FREE,
	BASIC:        BASIC,
	INTERMEDIATE: INTERMEDIATE,
	PLUS:         PLUS,
	PRO:          PRO,
}

var planTypeEnumFromNames = map[string]PlanType{
	"FREE":         FREE,
	"BASIC":        BASIC,
	"INTERMEDIATE": INTERMEDIATE,
	"PLUS":         PLUS,
	"PRO":          PRO,
}

func (this PlanType) Exists() bool {
	_, exists := planTypeEnum[this]
	return exists
}

func (this PlanType) FromValue(value string) PlanType {
	valueMap, exists := planTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this PlanType) Name() string {
	valueMap, exists := planTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
