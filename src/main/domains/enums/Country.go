package main_domains_enums

type Country string

const (
	BR Country = "BR"
	US Country = "US"
	FR Country = "FR"
)

var countryEnum = map[Country]Country{
	BR: BR,
	US: US,
	FR: FR,
}

var countryEnumFromNames = map[string]Country{
	"BR": BR,
	"US": US,
	"FR": FR,
}

func (this Country) Exists() bool {
	_, exists := countryEnum[this]
	return exists
}

func (this Country) FromValue(value string) Country {
	valueMap, exists := countryEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this Country) Name() string {
	valueMap, exists := countryEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
