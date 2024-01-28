package main_gateways_ws_beans

import (
	main_gateways_ws_beans_factories "baseapplicationgo/main/gateways/ws/beans/Factories"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	"sync"
)

var once sync.Once

var controllerBeans *ControllerBeans

type ControllerBeans struct {
	UserControllerV1Bean        *main_gateways_ws_v1.UserController
	FeatureControllerV1Bean     *main_gateways_ws_v1.FeaturesController
	RabbitMqControllerV1Bean    *main_gateways_ws_v1.RabbitMqController
	TransactionControllerV1Bean *main_gateways_ws_v1.TransactionController
	EmailControllerV1Bean       *main_gateways_ws_v1.EmailController
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
		UserControllerV1Bean:        main_gateways_ws_beans_factories.NewUserControllerBean().Get(),
		FeatureControllerV1Bean:     main_gateways_ws_beans_factories.NewFeatureControllerBean().Get(),
		RabbitMqControllerV1Bean:    main_gateways_ws_beans_factories.NewRabbitMqControllerBean().Get(),
		TransactionControllerV1Bean: main_gateways_ws_beans_factories.NewTransactionControllerBean().Get(),
		EmailControllerV1Bean:       main_gateways_ws_beans_factories.NewEmailControllerBean().Get(),
	}
}
