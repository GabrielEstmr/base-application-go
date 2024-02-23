package main_gateways_ws_v1

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_ws_commonsresources "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"net/http"
)

type SessionController struct {
	endSession                            main_usecases.EndSession
	refreshSession                        main_usecases.RefreshSession
	createInternalProviderUserSession     main_usecases.CreateInternalProviderUserSession
	createExternalAuthProviderUserSession main_usecases_lockers.AtomicLockedCreateExternalAuthProviderUserSession
	messageUtils                          main_utils_messages.ApplicationMessages
	logsMonitoringGateway                 main_gateways.LogsMonitoringGateway
	spanGateway                           main_gateways.SpanGateway
}

func NewSessionController(
	endSession main_usecases.EndSession,
	refreshSession main_usecases.RefreshSession,
	createInternalProviderUserSession main_usecases.CreateInternalProviderUserSession,
	createExternalAuthProviderUserSession main_usecases_lockers.AtomicLockedCreateExternalAuthProviderUserSession,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *SessionController {
	return &SessionController{
		endSession:                            endSession,
		refreshSession:                        refreshSession,
		createInternalProviderUserSession:     createInternalProviderUserSession,
		createExternalAuthProviderUserSession: createExternalAuthProviderUserSession,
		messageUtils:                          *main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway:                 logsMonitoringGateway,
		spanGateway:                           spanGateway,
	}
}

func (this *SessionController) CreateExternalProviderSession(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "SessionController-CreateExternalProviderSession")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Creating a new external-provider session")

	paramsRequest, err1 := new(main_gateways_ws_v1_request.ExternalProviderSessionRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	errV := paramsRequest.Validate()
	if errV != nil {
		this.logsMonitoringGateway.ERROR(span, errV.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errV
	}

	params, errT := paramsRequest.ToDomain()
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, errT.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errT
	}

	sessionCredentials, errC := this.createExternalAuthProviderUserSession.Execute(span.GetCtx(), params)
	if errC != nil {
		this.logsMonitoringGateway.ERROR(span, errC.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errC
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusOK,
		main_gateways_ws_v1_response.NewSessionCredentialsResponse(sessionCredentials)), nil

}

func (this *SessionController) CreateInternalProviderSession(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "SessionController-CreateInternalProviderSession")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Creating a new external-provider session")

	paramsRequest, err1 := new(main_gateways_ws_v1_request.InternalProviderSessionRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	errV := paramsRequest.Validate()
	if errV != nil {
		this.logsMonitoringGateway.ERROR(span, errV.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errV
	}

	sessionCredentials, errC := this.createInternalProviderUserSession.Execute(
		span.GetCtx(),
		paramsRequest.GetFirstUsername(),
		paramsRequest.GetFirstPassword())
	if errC != nil {
		this.logsMonitoringGateway.ERROR(span, errC.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errC
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusOK,
		main_gateways_ws_v1_response.NewSessionCredentialsResponse(sessionCredentials)), nil

}

func (this *SessionController) RefreshSession(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "SessionController-RefreshSession")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Refreshing session")

	paramsRequest, err1 := new(main_gateways_ws_v1_request.RefreshSessionRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	errV := paramsRequest.Validate()
	if errV != nil {
		this.logsMonitoringGateway.ERROR(span, errV.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errV
	}

	sessionCredentials, errC := this.refreshSession.Execute(
		span.GetCtx(),
		paramsRequest.GetFirstRefreshToken())
	if errC != nil {
		this.logsMonitoringGateway.ERROR(span, errC.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errC
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusOK,
		main_gateways_ws_v1_response.NewSessionCredentialsResponse(sessionCredentials)), nil

}

func (this *SessionController) EndSession(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "SessionController-RefreshSession")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Refreshing session")

	paramsRequest, err1 := new(main_gateways_ws_v1_request.EndSessionRequest).QueryParamsToObject(w, r)
	if err1 != nil {
		this.logsMonitoringGateway.ERROR(span, err1.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err1
	}

	errV := paramsRequest.Validate()
	if errV != nil {
		this.logsMonitoringGateway.ERROR(span, errV.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errV
	}

	errC := this.endSession.Execute(
		span.GetCtx(),
		paramsRequest.GetFirstRefreshToken())
	if errC != nil {
		this.logsMonitoringGateway.ERROR(span, errC.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errC
	}

	return *main_gateways_ws_commonsresources.
		NewControllerResponse(http.StatusNoContent, nil), nil
}
