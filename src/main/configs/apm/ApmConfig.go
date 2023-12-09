package main_configs_apm

import (
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	main_configs_apm_tracer "baseapplicationgo/main/configs/apm/tracer"
	main_configs_profile "baseapplicationgo/main/configs/profile"
	"context"
)

func InitiateApmConfig(mainCtx *context.Context) {
	profile := main_configs_profile.GetProfileBean().GetLowerCaseName()
	if profile != "local" {
		main_configs_apm_tracer.GetTracerProviderBean(mainCtx)
		main_configs_apm_metrics.GetMetricProviderBean(mainCtx)
	}
}
