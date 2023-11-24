package main_gateways_ws_beans

import (
	main_gateways_ws_beans_factories "baseapplicationgo/main/gateways/ws/beans/Factories"
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	"sync"
)

var once sync.Once

var controllers ControllerBeans
var controllerBeans *ControllerBeans

type ControllerBeans struct {
	UserControllerV1Bean gatewaysWsV1.UserController
}

func GetControllerBeans() *ControllerBeans {
	once.Do(func() {
		if controllerBeans == nil {
			controllers = SubscriptControllerBeans()
			controllerBeans = &controllers
		}
	})
	return controllerBeans
}

func SubscriptControllerBeans() ControllerBeans {
	return ControllerBeans{
		UserControllerV1Bean: main_gateways_ws_beans_factories.UserControllerBeanFactory(),
	}
}
