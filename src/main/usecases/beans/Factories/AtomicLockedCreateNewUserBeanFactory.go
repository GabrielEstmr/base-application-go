package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type AtomicLockedCreateNewUser struct {
}

func NewAtomicLockedCreateNewUser() *AtomicLockedCreateNewUser {
	return &AtomicLockedCreateNewUser{}
}

func (this *AtomicLockedCreateNewUser) Get() *main_usecases_lockers.AtomicLockedCreateNewUser {

	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases_lockers.NewAtomicLockedCreateNewUser(
		distributedLockGateway,
		NewCreateNewUserBean().Get(),
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
