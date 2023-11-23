package mainConfigsEnv

func GetBeanPropertyByName(envName EnvironmentProperty) string {
	GetEnvConfigBean()
	properties := *EnvValues
	property := properties[envName.GetDescription()]
	return property
}
