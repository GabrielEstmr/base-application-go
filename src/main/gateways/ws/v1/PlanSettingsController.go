package main_gateways_ws_v1

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_ws_commonsresources "baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_v1_request "baseapplicationgo/main/gateways/ws/v1/request"
	main_gateways_ws_v1_response "baseapplicationgo/main/gateways/ws/v1/response"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const _PLAN_SETTING_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"
const _PLAN_SETTING_CONTROLLER_MSG_KEY_ARCH_ISSUE = "exceptions.architecture.application.issue"

const _PLAN_SETTING_CONTROLLER_PATH_PREFIX = "/api/v1/plan-settings/"

type PlanSettingController struct {
	createNewPlanSetting     main_usecases.CreateNewPlanSetting
	findPlanSettingById      main_usecases.FindPlanSettingById
	findPlanSettingsByFilter main_usecases.FindPlanSettingsByFilter
	messageUtils             main_utils_messages.ApplicationMessages
	logsMonitoringGateway    main_gateways.LogsMonitoringGateway
	spanGateway              main_gateways.SpanGateway
}

func NewPlanSettingController(
	createNewPlanSetting main_usecases.CreateNewPlanSetting,
	findPlanSettingById main_usecases.FindPlanSettingById,
	findPlanSettingsByFilter main_usecases.FindPlanSettingsByFilter,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *PlanSettingController {
	return &PlanSettingController{
		createNewPlanSetting,
		findPlanSettingById,
		findPlanSettingsByFilter,
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
		spanGateway,
	}
}

func (this *PlanSettingController) CreatePlanSetting(_ http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "PlanSettingsController-CreatePlanSetting")
	defer span.End()

	this.logsMonitoringGateway.INFO(span, "Creating new plan setting")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_PLAN_SETTING_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	var planSettingRequest main_gateways_ws_v1_request.CreatePlanSetting
	if err = json.Unmarshal(requestBody, &planSettingRequest); err != nil {
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_PLAN_SETTING_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}

	bodyErr := planSettingRequest.Validate()
	bodyMetadataErr := planSettingRequest.Metadata.Validate()
	if bodyErr != nil || bodyMetadataErr != nil {
		this.logsMonitoringGateway.ERROR(span, bodyErr.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				bodyErr.Error())
	}

	planSetting := planSettingRequest.ToDomain()

	persistedPlanSetting, errApp := this.createNewPlanSetting.Execute(span.GetCtx(), planSetting)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusCreated,
		main_gateways_ws_v1_response.NewPlanSettingResponse(persistedPlanSetting)), nil
}

func (this *PlanSettingController) FindSettingById(
	w http.ResponseWriter,
	r *http.Request,
) (
	main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "PlanSettingController-CreatePlanSetting")
	defer span.End()

	id := strings.TrimPrefix(r.URL.Path, _PLAN_SETTING_CONTROLLER_PATH_PREFIX)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Finding PlanSetting by id: %s", id))
	planSetting, errApp := this.findPlanSettingById.Execute(span.GetCtx(), id)
	if errApp != nil {
		this.logsMonitoringGateway.ERROR(span, errApp.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errApp
	}

	return *main_gateways_ws_commonsresources.NewControllerResponse(
		http.StatusOK, main_gateways_ws_v1_response.NewPlanSettingResponse(planSetting)), nil
}

func (this *PlanSettingController) FindPlanSetting(
	w http.ResponseWriter,
	r *http.Request,
) (main_gateways_ws_commonsresources.ControllerResponse,
	main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(r.Context(), "PlanSettingController-FindPlanSetting")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Finding PlanSetting by query")

	filter, errParams := new(main_gateways_ws_v1_request.FindPlanSettingFilterRequest).QueryParamsToObject(w, r)
	if errParams != nil {
		this.logsMonitoringGateway.ERROR(span, errParams.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errParams
	}

	pageableReq, errP := new(main_gateways_ws_commonsresources.PageableRequest).QueryParamsToObject(w, r)
	if errP != nil {
		this.logsMonitoringGateway.ERROR(span, errP.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), errP
	}

	pageable, errT := pageableReq.ToDomain()
	if errT != nil {
		this.logsMonitoringGateway.ERROR(span, errT.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				_PLAN_SETTING_CONTROLLER_MSG_KEY_ARCH_ISSUE))
	}

	page, err := this.findPlanSettingsByFilter.Execute(span.GetCtx(), filter.ToDomain(), pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commonsresources.ControllerResponse), err
	}

	content := page.GetContent()
	respContent := make([]main_gateways_ws_v1_response.PlanSettingResponse, 0)
	for _, value := range content {
		planSetting := value.(main_domains.PlanSetting)
		respContent = append(respContent,
			*main_gateways_ws_v1_response.NewPlanSettingResponse(planSetting))
	}
	return *main_gateways_ws_commonsresources.NewControllerResponse(http.StatusOK,
		main_gateways_ws_commonsresources.NewPageResponse(
			respContent,
			page.GetPage(),
			page.GetSize(),
			page.GetTotalElements(),
			page.GetTotalPages())), nil
}
