package mainConfigsEnv

type EnvironmentProperty string

const mpIndicatorApplicationProfile = "MP_INDICATOR_APPLICATION_PROFILE"

const (
	MP_INDICATOR_APPLICATION_PROFILE EnvironmentProperty = mpIndicatorApplicationProfile
)

func (s EnvironmentProperty) GetDescription() string {
	switch s {
	case MP_INDICATOR_APPLICATION_PROFILE:
		return mpIndicatorApplicationProfile
	}
	return "unknown"
}
