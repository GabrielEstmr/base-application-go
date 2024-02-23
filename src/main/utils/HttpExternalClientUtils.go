package main_utils

import (
	gateways_authprovider_resources "baseapplicationgo/main/gateways/http/commons/response"
	"io"
	"net/http"
)

type HttpExternalClientUtils struct {
}

func NewHttpExternalClientUtils() *HttpExternalClientUtils {
	return &HttpExternalClientUtils{}
}

func (this *HttpExternalClientUtils) BuildHttpResponse(resp *http.Response) (gateways_authprovider_resources.HttpResponse, error) {
	bodyText, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return *new(gateways_authprovider_resources.HttpResponse), errRead
	}
	return *gateways_authprovider_resources.NewHttpResponse(resp, bodyText), nil
}

func (this *HttpExternalClientUtils) BuildHttpResponseWithBody(resp *http.Response, body string) (gateways_authprovider_resources.HttpResponse, error) {
	return *gateways_authprovider_resources.NewHttpResponse(resp, []byte(body)), nil
}

func (this *HttpExternalClientUtils) BuildHttpResponseWithBodyBytes(resp *http.Response, body []byte) (gateways_authprovider_resources.HttpResponse, error) {
	return *gateways_authprovider_resources.NewHttpResponse(resp, body), nil
}
