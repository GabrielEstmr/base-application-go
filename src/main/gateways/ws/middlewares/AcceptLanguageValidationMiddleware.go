package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"net/http"
)

const _MSG_KEY_LANGUAGE_HEADER_NOT_FOUND = "middlewares..headers.language.not.found.error"

type AcceptLanguageValidationMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
	messageUtils            main_utils_messages.ApplicationMessages
}

func NewAcceptLanguageValidationMiddleware() *AcceptLanguageValidationMiddleware {
	return &AcceptLanguageValidationMiddleware{
		acceptLanguageHeaderKey: "Accept-Language",
		stringUtils:             *main_utils.NewStringUtils(),
		messageUtils:            *main_utils_messages.NewApplicationMessages(),
	}
}

func (this *AcceptLanguageValidationMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		languageHeader := r.Header.Get(this.acceptLanguageHeaderKey)
		if !this.stringUtils.IsEmpty(languageHeader) {
			if this.stringUtils.IsEmpty(main_gateways_ws_commons.GetAllAvailableLanguages()[languageHeader]) {
				errApp := main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(_MSG_KEY_LANGUAGE_HEADER_NOT_FOUND))
				main_utils.ERROR_APP(w, errApp)
				return
			}
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
