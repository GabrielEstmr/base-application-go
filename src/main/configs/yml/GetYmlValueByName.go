package mainConfigsYml

func GetYmlValueByName(propertyName string) string {
	properties := *YmlConfigs
	property := properties[propertyName]
	return property.Value
}
