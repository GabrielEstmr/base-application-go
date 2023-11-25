package main_gateways_ws_beans

import (
	main_gateways_ws_beans_factories "baseapplicationgo/main/gateways/ws/beans/Factories"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	"sync"
)

var once sync.Once

var controllerBeans *ControllerBeans

type ControllerBeans struct {
	UserControllerV1Bean main_gateways_ws_v1.UserController
}

func GetControllerBeans() *ControllerBeans {
	once.Do(func() {
		if controllerBeans == nil {
			controllerBeans = subscriptControllerBeans()
		}
	})
	return controllerBeans
}

func subscriptControllerBeans() *ControllerBeans {
	return &ControllerBeans{
		UserControllerV1Bean: main_gateways_ws_beans_factories.UserControllerBeanFactory(),
	}
}
