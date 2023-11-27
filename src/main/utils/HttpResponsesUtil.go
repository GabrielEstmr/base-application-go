package main_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func newResponseError(code string, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		json.NewEncoder(w).Encode(newResponseError(fmt.Sprint(statusCode), "InternalServerError"))
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, error error) {
	r := newResponseError(strconv.Itoa(statusCode), error.Error())
	JSON(w, statusCode, r)
}
