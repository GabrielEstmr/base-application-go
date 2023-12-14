package main_gateways

type FeaturesGateway interface {
	IsEnabled(key string) (bool, error)
	IsDisabled(key string) (bool, error)
}
