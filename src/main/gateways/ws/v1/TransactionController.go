package main_gateways_ws_v1

import (
	main_gateways "baseapplicationgo/main/gateways"
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

func (this *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	span := this.spanGateway.Get(r.Context(), "TransactionController-CreateTransaction")
	defer span.End()

	this.logsMonitoringGateway.INFO(span, "Creating a new transaction")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		errLog := errors.New(
			this.messageUtils.GetDefaultLocale(
				_TRANSACTION_CONTROLLER_MSG_MALFORMED_REQUEST_BODY))
		main_utils.ERROR(w, http.StatusBadRequest, errLog)
		return
	}

	var createTransactionRequest main_gateways_ws_v1_request.CreateTransactionRequest
	if err = json.Unmarshal(requestBody, &createTransactionRequest); err != nil {
		main_utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bodyErr := createTransactionRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		main_utils.ERROR_APP(w, bodyErr)
		return
	}
	transaction := createTransactionRequest.ToDomain()

	persistedTransaction, errApp := this.createNewTransaction.Execute(r.Context(), transaction)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		main_utils.ERROR_APP(w, errApp)
		return
	}

	main_utils.JSON(
		w,
		http.StatusCreated,
		main_gateways_ws_v1_response.NewTransactionResponse(persistedTransaction))
}
