package main_configs_router

import (
	"github.com/gorilla/mux"
	"sync"
)

var once sync.Once
var routerBean *mux.Router = nil
var router mux.Router

func GetRouterBean() *mux.Router {
	once.Do(func() {
		if routerBean == nil {
			router = getMuxRouterRouter()
			routerBean = &router
		}

	})
	return routerBean
}

func getMuxRouterRouter() mux.Router {
	return *mux.NewRouter()
}
