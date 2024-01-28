package main_configs_apm_logs_impl

import (
	main_configs_apm_logs "baseapplicationgo/main/configs/apm/logs"
	main_gateways_logs_request2 "baseapplicationgo/main/configs/apm/logs/impl/request"
	main_configs_apm_logs_resources "baseapplicationgo/main/configs/apm/logs/resources"
	"bytes"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

const _MSG_LOKI_IMPL_ERROR_POST_LOKI_LOGS = "Error to send message to loki integration"

type LogsMethodsImpl struct {
	logConfig main_configs_apm_logs_resources.LogProviderConfig
}

func NewLogsGatewayImpl() *LogsMethodsImpl {
	return &LogsMethodsImpl{
		logConfig: *main_configs_apm_logs.GetLogProviderBean(),
	}
}

func (this *LogsMethodsImpl) DEBUG(
	span trace.Span,
	msg string,
	args ...any,
) {
	go this.postLog(METHOD_DEBUG, span, msg, args)
}

func (this *LogsMethodsImpl) WARN(
	span trace.Span,
	msg string,
	args ...any,
) {
	go this.postLog(METHOD_WARN, span, msg, args)
}

func (this *LogsMethodsImpl) INFO(
	span trace.Span,
	msg string,
	args ...any,
) {
	go this.postLog(METHOD_INFO, span, msg, args)
}

func (this *LogsMethodsImpl) ERROR(
	span trace.Span,
	msg string,
	args ...any,
) {
	go this.postLog(METHOD_ERROR, span, msg, args)
}

func (this *LogsMethodsImpl) postLog(
	methodLog string,
	span trace.Span,
	msg string,
	args ...any,
) {
	client := http.Client{
		Timeout: 2000 * time.Millisecond,
	}
	baseUrl := "http://" + this.logConfig.GetHost() + "/loki/api/v1/push"

	stringJson, err := main_gateways_logs_request2.NewLogMessage(
		span.SpanContext().TraceID().String(),
		span.SpanContext().SpanID().String(),
		methodLog, msg, "", 1).TO_STRING_JSON()
	if err != nil {
		fmt.Println(err.Error())
	}

	request := main_gateways_logs_request2.NewCreateLogRequest(methodLog, this.logConfig.GetAppName(), stringJson)
	body, _ := json.Marshal(request)
	payload := bytes.NewBuffer(body)
	_, errPost := client.Post(baseUrl, "application/json", payload)
	if errPost != nil {
		fmt.Println(_MSG_LOKI_IMPL_ERROR_POST_LOKI_LOGS, errPost.Error())
	}

}
