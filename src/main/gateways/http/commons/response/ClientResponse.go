package main_gateways_http_commons_response

import "net/http"

type HttpResponse struct {
	closedResponse *http.Response
	responseBody   []byte
}

func NewHttpResponse(
	closedResponse *http.Response,
	responseBody []byte,
) *HttpResponse {
	return &HttpResponse{
		closedResponse: closedResponse,
		responseBody:   responseBody,
	}
}

func (this HttpResponse) GetClosedResponse() *http.Response {
	return this.closedResponse
}

func (this HttpResponse) GetResponseBody() []byte {
	return this.responseBody
}
