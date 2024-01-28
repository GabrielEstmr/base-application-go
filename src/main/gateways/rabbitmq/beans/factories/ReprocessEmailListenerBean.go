package main_gateways_rabbitmq_beans_factories

import (
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type ReprocessEmailListener struct {
}

func NewReprocessEmailListener() *ReprocessEmailListener {
	return &ReprocessEmailListener{}
}

func (this *ReprocessEmailListener) Get() *main_gateways_rabbitmq_listeners.ReprocessEmailListener {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	return main_gateways_rabbitmq_listeners.NewReprocessEmailListener(*usecaseBeans.ReprocessEmailEvent)
}
