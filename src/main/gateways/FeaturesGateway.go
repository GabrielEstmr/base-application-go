package main_gateways

import (
	main_domains_features "baseapplicationgo/main/domains/features"
)

type FeaturesGateway interface {
	IsEnabled(key string) (bool, error)
	IsDisabled(key string) (bool, error)
	Enable(key string) (main_domains_features.FeaturesData, error)
	Disable(key string) (main_domains_features.FeaturesData, error)
}
