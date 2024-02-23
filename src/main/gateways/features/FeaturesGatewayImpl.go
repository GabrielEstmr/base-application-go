package main_gateways_features

import (
	main_configs_ff "baseapplicationgo/main/configs/ff"
	main_configs_ff_lib "baseapplicationgo/main/configs/ff/lib"
	main_domains_features "baseapplicationgo/main/domains/features"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
)

type FeaturesGatewayImpl struct {
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	ffConfig              *main_configs_ff_lib.FfConfig
}

func NewFeaturesGatewayImpl() *FeaturesGatewayImpl {
	return &FeaturesGatewayImpl{
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		ffConfig:              main_configs_ff.GetFfConfigDataBean(),
	}
}

func (this *FeaturesGatewayImpl) IsEnabled(key string) (bool, error) {
	return this.ffConfig.GetFeaturesMethods().IsEnabled(key)
}

func (this *FeaturesGatewayImpl) IsDisabled(key string) (bool, error) {
	return this.ffConfig.GetFeaturesMethods().IsDisabled(key)
}

func (this *FeaturesGatewayImpl) Enable(key string) (main_domains_features.FeaturesData, error) {
	span := this.spanGateway.Get(context.Background(), "FeaturesGatewayImpl-Enable")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("Enabling fature %s", key))
	feature, err := this.ffConfig.GetFeaturesMethods().Enable(key)
	return *main_domains_features.NewFeaturesData(feature), err
}

func (this *FeaturesGatewayImpl) Disable(key string) (main_domains_features.FeaturesData, error) {
	span := this.spanGateway.Get(context.Background(), "FeaturesGatewayImpl-Enable")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("Desabling fature %s", key))
	feature, err := this.ffConfig.GetFeaturesMethods().Disable(key)
	return *main_domains_features.NewFeaturesData(feature), err
}
