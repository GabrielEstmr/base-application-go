package main_utils

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Code     string   `json:"code"`
	Messages []string `json:"message"`
}

func newResponseError(code string, messages []string) *ResponseError {
	return &ResponseError{
		Code:     code,
		Messages: messages,
	}
}

func newResponseErrorSglMsg(code string, message string) *ResponseError {
	return &ResponseError{
		Code:     code,
		Messages: []string{message},
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		json.NewEncoder(w).Encode(newResponseErrorSglMsg(strconv.Itoa(statusCode), "InternalServerError"))
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, error error) {
	r := newResponseErrorSglMsg(strconv.Itoa(statusCode), error.Error())
	JSON(w, statusCode, r)
}

func ERROR_APP(w http.ResponseWriter, appException main_domains_exceptions.ApplicationException) {
	r := newResponseError(strconv.Itoa(appException.GetCode()), appException.GetMessages())
	JSON(w, appException.GetCode(), r)
}
