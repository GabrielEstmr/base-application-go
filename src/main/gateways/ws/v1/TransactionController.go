package main_gateways_ws_v1

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"io"
	"net/http"
)

const _TRANSACTION_CONTROLLER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

type TransactionController struct {
	createNewTransaction  *main_usecases.PersistTransaction
	messageUtils          main_utils_messages.ApplicationMessages
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewTransactionController(
	createNewTransaction *main_usecases.PersistTransaction,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *TransactionController {
	return &TransactionController{
		createNewTransaction,
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
		spanGateway,
	}
}

func (this *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commons.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "TransactionController-CreateTransaction")
	defer span.End()

	this.logsMonitoringGateway.INFO(span, "Creating a new transaction")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_TRANSACTION_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
	}

	var createTransactionRequest main_gateways_ws_v1_request.CreateTransactionRequest
	if err = json.Unmarshal(requestBody, &createTransactionRequest); err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_TRANSACTION_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
	}

	bodyErr := createTransactionRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			bodyErr
	}
	transaction := createTransactionRequest.ToDomain()

	persistedTransaction, errApp := this.createNewTransaction.Execute(r.Context(), transaction)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			errApp
	}

	return *main_gateways_ws_commons.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewTransactionResponse(persistedTransaction)), nil
}
