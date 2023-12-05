package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_usecases "baseapplicationgo/main/usecases"
)

type FindUsersByFilterBean struct {
}

func NewFindUsersByFilterBean() *FindUsersByFilterBean {
	return &FindUsersByFilterBean{}
}

func (this *FindUsersByFilterBean) Get() *main_usecases.FindUsersByFilter {
	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(*userRepository)
	return main_usecases.NewFindUsersByFilter(userDatabaseGateway)
}
