package main_utils

import (
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"encoding/json"
	"net/http"
)

const _QUERY_BUILDER_MSG_MALFORMED_REQUEST_BODY = "controllers.param.missing.or.malformed"

type QueryParams struct {
	obj interface{}
}

func (q *QueryParams) GetObj() interface{} {
	return q.obj
}

func NewQueryParams(obj interface{}) *QueryParams {
	return &QueryParams{obj: obj}
}

func QueryParamsToObject(any *QueryParams, w http.ResponseWriter, r *http.Request) (*QueryParams, error) {
	if err := r.ParseForm(); err != nil {
		return any, err
	}
	data, err := json.Marshal(r.Form)
	if err != nil {
		return any, err
	}
	if err = json.Unmarshal(data, any.GetObj()); err != nil {
		errLog := main_domains_exceptions.NewBadRequestExceptionSglMsg(
			main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(
				_QUERY_BUILDER_MSG_MALFORMED_REQUEST_BODY))
		ERROR_APP(w, errLog)
	}
	return any, nil
}
