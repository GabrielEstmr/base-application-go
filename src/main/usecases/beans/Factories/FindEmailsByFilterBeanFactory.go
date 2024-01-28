package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type FindEmailsByFilterBean struct {
}

func NewFindEmailsByFilterBean() *FindEmailsByFilterBean {
	return &FindEmailsByFilterBean{}
}

func (this *FindEmailsByFilterBean) Get() *main_usecases.FindEmailsByFilter {

	emailRepo := main_gateways_mongodb_repositories.NewEmailRepository()
	var emailDatabaseGateway main_gateways.EmailDatabaseGateway = main_gateways_mongodb.NewEmailDatabaseGatewayImpl(*emailRepo)

	return main_usecases.NewFindEmailsByFilter(
		emailDatabaseGateway,
		*main_utils_messages.NewApplicationMessages(),
		main_gateways_spans.NewSpanGatewayImpl(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	)
}
