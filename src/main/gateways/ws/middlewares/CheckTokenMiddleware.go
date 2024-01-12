package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"net/http"
)

const _MSG_KEY_UNAUTHORIZED_ERROR = "exceptions.unauthorized.error"

type CheckTokenBeforeRequestMiddleware struct {
	authHeaderKey string
	stringUtils   main_utils.StringUtils
	messageUtils  main_utils_messages.ApplicationMessages
}

func NewCheckTokenBeforeRequestMiddleware() *CheckTokenBeforeRequestMiddleware {
	return &CheckTokenBeforeRequestMiddleware{
		authHeaderKey: "Authorization",
		stringUtils:   *main_utils.NewStringUtils(),
		messageUtils:  *main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CheckTokenBeforeRequestMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		languageHeader := r.Header.Get(this.authHeaderKey)
		if this.stringUtils.IsEmpty(languageHeader) {
			errApp := main_domains_exceptions.NewUnauthorizedExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_KEY_UNAUTHORIZED_ERROR))
			main_utils.ERROR_APP(w, errApp)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
