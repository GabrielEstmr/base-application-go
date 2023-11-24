package main_gateways_ws_beans

import (
	main_gateways_ws_beans_factories "baseapplicationgo/main/gateways/ws/beans/Factories"
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	"sync"
)

var once sync.Once

var Beans *ControllerBeans = nil

type ControllerBeans struct {
	UserControllerV1Bean *gatewaysWsV1.UserController
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
		UserControllerV1Bean: main_gateways_ws_beans_factories.UserControllerBeanFunc(),
	}
}
