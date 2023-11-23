package mainConfigsRouter

import (
	"github.com/gorilla/mux"
	"sync"
)

var Router *mux.Router = nil

var once sync.Once

func GetRouterBean() *mux.Router {
	once.Do(func() {

		if Router == nil {
			Router = getMuxRouterRouter()
		}

	})
	return Router
}

func getMuxRouterRouter() *mux.Router {
	return mux.NewRouter()
}
