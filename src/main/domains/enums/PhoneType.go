package main_domains_enums

type PhoneType string

const (
	MOBILE   PhoneType = "MOBILE"
	LANDLINE PhoneType = "LANDLINE"
)

var phoneTypeEnum = map[PhoneType]PhoneType{
	MOBILE:   MOBILE,
	LANDLINE: LANDLINE,
}

var phoneTypeEnumFromNames = map[string]PhoneType{
	"MOBILE":   MOBILE,
	"LANDLINE": LANDLINE,
}

func (this PhoneType) Exists() bool {
	_, exists := phoneTypeEnum[this]
	return exists
}

func (this PhoneType) FromValue(value string) PhoneType {
	valueMap, exists := phoneTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this PhoneType) Name() string {
	valueMap, exists := phoneTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
