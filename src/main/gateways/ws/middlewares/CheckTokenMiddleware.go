package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"log"
	"net/http"
)

const _MSG_KEY_UNAUTHORIZED_ERROR = "exceptions.unauthorized.error"

type CheckTokenMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
	messageUtils            main_utils_messages.ApplicationMessages
}

func NewCheckTokenMiddleware() *CheckTokenMiddleware {
	return &CheckTokenMiddleware{
		acceptLanguageHeaderKey: "Authorization",
		stringUtils:             *main_utils.NewStringUtils(),
		messageUtils:            *main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CheckTokenMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("RUN BEFORE!!!!")
		languageHeader := r.Header.Get(this.acceptLanguageHeaderKey)
		if this.stringUtils.IsEmpty(languageHeader) {
			errApp := main_domains_exceptions.NewUnauthorizedExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_KEY_UNAUTHORIZED_ERROR))
			main_utils.ERROR_APP(w, errApp)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
