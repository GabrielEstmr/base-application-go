package main_configs_yml

func GetYmlValueByName(propertyName string) string {
	properties := *GetYmlConfigBean()
	property := properties[propertyName]
	return property.Value
}
