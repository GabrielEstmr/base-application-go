package main_gateways_ws_middlewares

import (
	"net/http"
)

const _LANGUAGE_MIDDLEWARE_STRING_EMPTY = ""

type ContentTypeLanguageMiddleware struct {
}

func NewContentTypeLanguageMiddleware() *ContentTypeLanguageMiddleware {
	return &ContentTypeLanguageMiddleware{}
}

func (this *ContentTypeLanguageMiddleware) LogMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		languageHeader := r.Header.Get("User-Agent")
		if languageHeader == _LANGUAGE_MIDDLEWARE_STRING_EMPTY {
			r.Header.Set("Content-Type-Language", "en-US")
		}
	})
}
