package main_configs

import (
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	main_configs_apm_tracer "baseapplicationgo/main/configs/apm/tracer"
	main_configs_mongo "baseapplicationgo/main/configs/mongodb"
	"context"
	"log"
)

const _MSG_TERMINATING_APPLICATION_BEANS = "Terminating application's configuration beans"
const _MSG_APPLICATION_BEANS_TERMINATED = "Application configuration beans successfully terminated"

func TerminateConfigBeans(mainCtx *context.Context) {
	log.Println(_MSG_TERMINATING_APPLICATION_BEANS)
	main_configs_mongo.CloseConnection()
	main_configs_apm_tracer.ShutdownTracerProvider(mainCtx)
	main_configs_apm_metrics.ShutdownMetricProvider(mainCtx)
	log.Println(_MSG_APPLICATION_BEANS_TERMINATED)
}
