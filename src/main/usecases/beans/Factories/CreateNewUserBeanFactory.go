package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_usecases "baseapplicationgo/main/usecases"
)

func CreateNewUserBean() *main_usecases.CreateNewUser {
	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(userRepository)
	return main_usecases.NewCreateNewUser(userDatabaseGateway)
}
