package main_domains_enums

type Language string

const (
	PT_BR Language = "PT_BR"
	EN_US Language = "EN_US"
	FR_FR Language = "FR_FR"
)

var languageEnum = map[Language]Language{
	PT_BR: PT_BR,
	EN_US: EN_US,
	FR_FR: FR_FR,
}

var languageEnumFromNames = map[string]Language{
	"PT_BR": PT_BR,
	"EN_US": EN_US,
	"FR_FR": FR_FR,
}

func (this Language) Exists() bool {
	_, exists := languageEnum[this]
	return exists
}

func (this Language) FromValue(value string) Language {
	valueMap, exists := languageEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this Language) Name() string {
	valueMap, exists := languageEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
