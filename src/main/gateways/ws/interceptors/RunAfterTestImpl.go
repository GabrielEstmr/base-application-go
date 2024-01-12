package main_gateways_ws_interceptors

import (
	main_utils "baseapplicationgo/main/utils"
	"log"
	"net/http"
)

type RunAfterTestImpl struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
}

func NewRunAfterTestImpl() *RunAfterTestImpl {
	return &RunAfterTestImpl{
		acceptLanguageHeaderKey: "Accept-Language",
		stringUtils:             *main_utils.NewStringUtils(),
	}
}

func (this *RunAfterTestImpl) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Println("runsafterTest TEST FINAL")
	}
	return http.HandlerFunc(fn)
}

//
//func runsafterTest(h http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		h.ServeHTTP(w, r)
//		log.Println("runsafterTest TEST FINAL")
//		//w.Write([]byte("run after, "))
//	}
//	return http.HandlerFunc(fn)
//}
//
//func runsbeforeTest(h http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		log.Println("runsbeforeTest TEST FINAL")
//		//w.Write([]byte("run before, "))
//		h.ServeHTTP(w, r)
//	}
//	return http.HandlerFunc(fn)
//}
