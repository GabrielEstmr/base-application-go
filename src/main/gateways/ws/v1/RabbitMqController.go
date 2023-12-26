package main_gateways_ws_v1

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_features "baseapplicationgo/main/gateways/features"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"errors"
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

func (this *RabbitMqController) CreateRabbitMqTransactionEvent(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	span := this.spanGateway.Get(ctx, "RabbitMqController-CreateRabbitMqTransactionEvent")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Creating a new transaction event")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		errLog := errors.New(
			this.messageUtils.GetDefaultLocale(
				_RABBITMQ_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
		main_utils.ERROR(w, http.StatusBadRequest, errLog)
		return
	}

	var transactionRequest main_gateways_ws_v1_request.CreateTransactionRequest
	if err = json.Unmarshal(requestBody, &transactionRequest); err != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bodyErr := transactionRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		main_utils.ERROR_APP(w, bodyErr)
		return
	}
	transaction := transactionRequest.ToDomain()

	errApp := this.createTransactionAmqpEvent.Execute(ctx, transaction)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(
		w,
		http.StatusCreated,
		main_gateways_ws_v1_response.NewTransactionResponse(transaction))
}
