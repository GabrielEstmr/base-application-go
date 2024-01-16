package main_gateways_ws_v1

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_features "baseapplicationgo/main/gateways/features"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"io"
	"net/http"
)

const _RABBITMQ_CONTROLLER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

type RabbitMqController struct {
	createTransactionAmqpEvent *main_usecases.CreateTransactionAmqpEvent
	featuresGateway            main_gateways.FeaturesGateway
	messageUtils               main_utils_messages.ApplicationMessages
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
	spanGateway                main_gateways.SpanGateway
}

func NewRabbitMqController(
	createTransactionAmqpEvent *main_usecases.CreateTransactionAmqpEvent,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *RabbitMqController {
	return &RabbitMqController{
		createTransactionAmqpEvent,
		main_gateways_features.NewFeaturesGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
		spanGateway,
	}
}

func (this *RabbitMqController) CreateRabbitMqTransactionEvent(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commons.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	ctx := r.Context()
	span := this.spanGateway.Get(ctx, "RabbitMqController-CreateRabbitMqTransactionEvent")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Creating a new transaction event")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_RABBITMQ_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
	}

	var transactionRequest main_gateways_ws_v1_request.CreateTransactionRequest
	if err = json.Unmarshal(requestBody, &transactionRequest); err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_RABBITMQ_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
	}

	bodyErr := transactionRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commons.ControllerResponse), bodyErr
	}
	transaction := transactionRequest.ToDomain()

	errApp := this.createTransactionAmqpEvent.Execute(span.GetCtx(), transaction)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commons.ControllerResponse), errApp
	}

	return *main_gateways_ws_commons.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewTransactionResponse(transaction)), nil
}
