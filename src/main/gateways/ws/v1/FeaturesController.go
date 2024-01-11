package main_gateways_ws_v1

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_features "baseapplicationgo/main/gateways/features"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"fmt"
	"net/http"
	"strings"
)

const _FEATURES_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"
const _FEATURES_CONTROLLER_MSG_KEY_RESOURCE_NOT_FOUND = "exceptions.not.found.resource.error"
const _FEATURES_CONTROLLER_MSG_KEY_ARCH_ISSUE = "exceptions.architecture.application.issue"

const _FEATURES_CONTROLLER_PATH_PREFIX = "/api/v1/features/"
const _FEATURES_CONTROLLER_PATH_SUFFIX_ENABLE = "/enable"
const _FEATURES_CONTROLLER_PATH_SUFFIX_DISABLE = "/disable"

type FeaturesController struct {
	featuresGateway       main_gateways.FeaturesGateway
	messageUtils          main_utils_messages.ApplicationMessages
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewFeaturesController(
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *FeaturesController {
	return &FeaturesController{
		main_gateways_features.NewFeaturesGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
		logsMonitoringGateway,
		spanGateway,
	}
}

func (this *FeaturesController) EnableFeatureByKey(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commons.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "FeaturesController-EnableFeatureByKey")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Enable feature by key")

	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_ENABLE)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Enable feature with key: %s", key))
	feature, err := this.featuresGateway.Enable(key)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewBadRequestExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_FEATURES_CONTROLLER_MSG_KEY_MALFORMED_REQUEST_BODY))
	}
	if feature.IsEmpty() {
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_FEATURES_CONTROLLER_MSG_KEY_RESOURCE_NOT_FOUND))
	}

	return *main_gateways_ws_commons.NewControllerResponse(
		http.StatusOK, fmt.Sprintf("Feature with key %s has been enabled", key)), nil
}

func (this *FeaturesController) DisableFeatureByKey(w http.ResponseWriter, r *http.Request) (
	main_gateways_ws_commons.ControllerResponse,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(r.Context(), "FeaturesController-DisableFeatureByKey")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Disable feature by key")

	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_DISABLE)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Disable feature with key: %s", key))
	feature, err := this.featuresGateway.Disable(key)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_FEATURES_CONTROLLER_MSG_KEY_ARCH_ISSUE))
	}
	if feature.IsEmpty() {
		return *new(main_gateways_ws_commons.ControllerResponse),
			main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_FEATURES_CONTROLLER_MSG_KEY_RESOURCE_NOT_FOUND))
	}

	return *main_gateways_ws_commons.NewControllerResponse(
		http.StatusOK, fmt.Sprintf("Disable with key %s has been disable", key)), nil
}
