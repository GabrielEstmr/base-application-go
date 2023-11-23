package main_gateways_ws_beans

import (
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	"sync"
)

var once sync.Once

var Beans *ControllerBeans = nil

type ControllerBeans struct {
	AccountControllerBean *gatewaysWsV1.AccountController
}

func GetControllerBeans() *ControllerBeans {
	once.Do(func() {
		if Beans == nil {
			Beans = NewIndicatorController()
		}
	})
	return Beans
}

func NewIndicatorController() *ControllerBeans {
	return &ControllerBeans{
		AccountControllerBean: accountControllerBeanFunc(),
	}
}
