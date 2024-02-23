package main_configs_yml

func GetYmlValueByName(propertyName YmlToken) string {
	properties := *GetYmlConfigBean()
	property := properties[propertyName.GetToken()]
	return property.Value
}
