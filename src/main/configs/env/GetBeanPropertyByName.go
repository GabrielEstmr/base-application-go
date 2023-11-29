package main_configs_env

func GetBeanPropertyByName(envName EnvironmentProperty) string {
	properties := *GetEnvConfigBean()
	property := properties[envName.GetDescription()]
	return property
}
