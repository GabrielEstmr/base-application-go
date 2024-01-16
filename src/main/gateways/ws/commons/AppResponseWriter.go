package main_gateways_ws_commons

import "net/http"

// Example of extends on Golang
type AppResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewAppResponseWriter(w http.ResponseWriter) *AppResponseWriter {
	return &AppResponseWriter{w, http.StatusOK}
}

func (this *AppResponseWriter) WriteHeader(code int) {
	this.statusCode = code
	this.ResponseWriter.WriteHeader(code)
}

func (this *AppResponseWriter) IsError() bool {
	return this.statusCode >= 400
}

func (this *AppResponseWriter) GetStatusCode() int {
	return this.statusCode
}
