package main_configs_apm_metrics

import (
	"context"
	"go.opentelemetry.io/otel/metric"
	"sync"
)

var globalMeter *metric.Meter
var once sync.Once

func GetGlobalMeterBean() *metric.Meter {
	once.Do(func() {

		if globalMeter == nil {
			globalMeter = getGlobalMeter()
		}

	})
	return globalMeter
}

func getGlobalMeter() *metric.Meter {
	ctx := context.Background()
	metricProviderBean := GetMetricProviderBean(&ctx)
	meter := metricProviderBean.Meter("Name App")
	return &meter
}
