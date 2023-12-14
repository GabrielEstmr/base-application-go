package main_configs_ff_lib

type FeaturesMethods interface {
	IsEnabled(key string) (bool, error)
	IsDisabled(key string) (bool, error)
}
