package main_configs_yml

func GetYmlValueByName(propertyName string) string {
	properties := *YmlConfigs
	property := properties[propertyName]
	return property.Value
}
