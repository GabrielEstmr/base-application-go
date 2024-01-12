package main_gateways_ws_middlewares

import (
	main_utils "baseapplicationgo/main/utils"
	"log"
	"net/http"
)

type AcceptJsonContentTypeHeaderMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
}

func NewAcceptJsonContentTypeHeaderMiddleware() *AcceptJsonContentTypeHeaderMiddleware {
	return &AcceptJsonContentTypeHeaderMiddleware{
		acceptLanguageHeaderKey: "Accept-Language",
		stringUtils:             *main_utils.NewStringUtils(),
	}
}

func (this *AcceptJsonContentTypeHeaderMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("RUNS BEFORE")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
