package main_utils

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	ScopeError string `json:"scopeError"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	if err := setErrorIfParseFails(w, data); err != nil {
		return
	}
	w.WriteHeader(statusCode)
}

func ERROR(w http.ResponseWriter, statusCode int, scopeError error) {
	JSON(w, statusCode, responseError{ScopeError: scopeError.Error()})
}

func setErrorIfParseFails(w http.ResponseWriter, data interface{}) error {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ERROR(w, http.StatusInternalServerError, err)
		return err
	}
	return nil
}
