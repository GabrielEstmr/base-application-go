package main_gateways_ws_v1

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_commonsresources "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"io"
	"net/http"
)

const _EMAIL_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"
const _EMAIL_CONTROLLER_MSG_KEY_ARCH_ISSUE = "exceptions.architecture.application.issue"

type EmailController struct {
	sendEmailEventsToReprocess *main_usecases.SendEmailEventsToReprocess
	findEmailsByFilter         *main_usecases.FindEmailsByFilter
	messageUtils               main_utils_messages.ApplicationMessages
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
	spanGateway                main_gateways.SpanGateway
}

func NewEmailController(
	sendEmailEventsToReprocess *main_usecases.SendEmailEventsToReprocess,
	findEmailsByFilter *main_usecases.FindEmailsByFilter,
) *EmailController {
	return &EmailController{
		sendEmailEventsToReprocess,
		findEmailsByFilter,
		*main_utils_messages.NewApplicationMessages(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *EmailController) ReprocessEmail(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "EmailController-ReprocessEmail")
	defer span.End()

	this.logsMonitoringGateway.INFO(span, "Reprocessing Email")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_EMAIL_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var emailRequest main_gateways_ws_v1_request.ReprocessEmailRequest
	if err = json.Unmarshal(requestBody, &emailRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_EMAIL_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := emailRequest.Validate()
	if bodyErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}

	this.sendEmailEventsToReprocess.Execute(span.GetCtx(), emailRequest.Ids)

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusAccepted,
		nil), nil
}

func (this *EmailController) FindEmail(
	w http.ResponseWriter,
	r *http.Request,
) (main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(r.Context(), "EmailController-FindEmail")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding Email by query")

	filter, err1 := new(main_gateways_ws_v1_request.FindEmailsFilterRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	pageableReq, err2 := new(main_gateways_ws_commonsresources.PageableRequest).QueryParamsToObject(w, r)
	if err2 != nil {
		this.logsMonitoringGateway.ERROR(span, err2.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err2
	}
	pageable, errT := pageableReq.ToDomain()
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, errT.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				_EMAIL_CONTROLLER_MSG_KEY_ARCH_ISSUE))
	}
	page, err := this.findEmailsByFilter.Execute(span.GetCtx(), filter.ToDomain(), pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err
	}

	content := page.GetContent()
	respContent := make([]main_gateways_ws_v1_response.EmailResponse, 0)
	for _, value := range content {
		email := value.(main_domains.Email)
		respContent = append(respContent,
			*main_gateways_ws_v1_response.NewEmailResponse(email))
	}
	return *main_gateways_ws_commonsresources.NewControllerResponse(http.StatusOK,
		main_gateways_ws_commonsresources.NewPageResponse(
			respContent,
			page.GetPage(),
			page.GetSize(),
			page.GetTotalElements(),
			page.GetTotalPages())), nil
}
