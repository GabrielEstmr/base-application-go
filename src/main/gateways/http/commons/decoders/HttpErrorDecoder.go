package main_gateways_http_commons_decoders

import (
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_http_commons_response "baseapplicationgo/main/gateways/http/commons/response"
)

type HttpErrorDecoder interface {
	DecodeErrors(
		span main_domains_apm.SpanLogInfo,
		response main_gateways_http_commons_response.HttpResponse,
		err error,
	) main_domains_exceptions.ApplicationException
}
