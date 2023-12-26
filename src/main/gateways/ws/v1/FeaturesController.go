package main_gateways_ws_v1

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_features "baseapplicationgo/main/gateways/features"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"fmt"
	"net/http"
	"strings"
)

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

func (this *FeaturesController) EnableFeatureByKey(w http.ResponseWriter, r *http.Request) {

	span := this.spanGateway.Get(r.Context(), "FeaturesController-EnableFeatureByKey")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Enable feature by key")

	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_ENABLE)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Enable feature with key: %s", key))
	feature, err := this.featuresGateway.Enable(key)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		main_utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if feature.IsEmpty() {
		main_utils.ERROR(w, http.StatusNotFound, err)
		return
	}

	main_utils.JSON(w, http.StatusOK, fmt.Sprintf("Feature with key %s has been enabled", key))
}

func (this *FeaturesController) DisableFeatureByKey(w http.ResponseWriter, r *http.Request) {

	span := this.spanGateway.Get(r.Context(), "FeaturesController-DisableFeatureByKey")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "Disable feature by key")

	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, _FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_DISABLE)
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("Disable feature with key: %s", key))
	feature, err := this.featuresGateway.Disable(key)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		main_utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if feature.IsEmpty() {
		main_utils.ERROR(w, http.StatusNotFound, err)
		return
	}

	main_utils.JSON(w, http.StatusOK, fmt.Sprintf("Disable with key %s has been disable", key))
}
