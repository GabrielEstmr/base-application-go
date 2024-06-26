package main_utils

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"encoding/json"
	"net/http"
)

const _QUERY_BUILDER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"
const _QUERY_BUILDER_MSG_ARCH_ISSUE = "exceptions.architecture.application.issue"

type QueryParams struct {
	obj interface{}
}

func (this *QueryParams) GetObj() interface{} {
	return this.obj
}

func NewQueryParams(obj interface{}) *QueryParams {
	return &QueryParams{obj: obj}
}

func QueryParamsToObject(
	any *QueryParams,
	w http.ResponseWriter,
	r *http.Request,
) (
	*QueryParams,
	main_domains_exceptions.ApplicationException,
) {
	if err := r.ParseForm(); err != nil {
		return nil, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			main_utils_messages.NewApplicationMessages().GetDefaultLocale(_QUERY_BUILDER_MSG_ARCH_ISSUE))
	}
	data, err := json.Marshal(r.Form)
	if err != nil {
		return nil, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			main_utils_messages.NewApplicationMessages().GetDefaultLocale(
				_QUERY_BUILDER_MSG_ARCH_ISSUE))
	}
	if err2 := json.Unmarshal(data, any.GetObj()); err2 != nil {
		errLog := main_domains_exceptions.NewBadRequestExceptionSglMsg(
			main_utils_messages.NewApplicationMessages().GetDefaultLocale(
				_QUERY_BUILDER_MSG_MALFORMED_REQUEST_BODY))
		return nil, errLog
	}
	return any, nil
}
