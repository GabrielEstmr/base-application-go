package main_gateways_http_commons_decoders

import (
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_http_commons_response "baseapplicationgo/main/gateways/http/commons/response"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	"fmt"
	"net/http"
)

type DefaultHttpErrorDecoder struct {
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewDefaultHttpErrorDecoder() *DefaultHttpErrorDecoder {
	return &DefaultHttpErrorDecoder{
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl()}
}

func (this *DefaultHttpErrorDecoder) DecodeErrors(
	span main_domains_apm.SpanLogInfo,
	response main_gateways_http_commons_response.HttpResponse,
	err error,
) main_domains_exceptions.ApplicationException {
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("internal server error")
	}

	if response.GetClosedResponse().StatusCode == http.StatusBadRequest {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(response.GetClosedResponse().Status))
		return main_domains_exceptions.NewBadRequestExceptionSglMsg("bad request")
	}

	if response.GetClosedResponse().StatusCode == http.StatusNotFound {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(response.GetClosedResponse().Status))
		return main_domains_exceptions.NewResourceNotFoundExceptionSglMsg("not found")
	}

	if response.GetClosedResponse().StatusCode == http.StatusConflict {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(response.GetClosedResponse().Status))
		return main_domains_exceptions.NewConflictExceptionSglMsg("conflict")
	}

	if response.GetClosedResponse().StatusCode == http.StatusUnauthorized {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(response.GetClosedResponse().Status))
		return main_domains_exceptions.NewUnauthorizedExceptionSglMsg("unauthorized")
	}

	if response.GetClosedResponse().StatusCode == http.StatusForbidden {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(response.GetClosedResponse().Status))
		return main_domains_exceptions.NewForbiddenExceptionSglMsg("forbidden")
	}

	return nil
}
