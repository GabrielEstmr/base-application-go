package main_usecases_beans

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_usecases "baseapplicationgo/main/usecases"
	"sync"
)

var once sync.Once
var CreateNewUserBean *main_usecases.CreateNewUser
var CreateNewUser main_usecases.CreateNewUser

func CreateNewUserBeanFactory() *main_usecases.CreateNewUser {
	once.Do(func() {
		if CreateNewUserBean == nil {
			CreateNewUser = createNewUserBean()
			CreateNewUserBean = &CreateNewUser
		}
	})
	return CreateNewUserBean
}

func createNewUserBean() main_usecases.CreateNewUser {
	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(userRepository)
	CreateNewUser = *main_usecases.NewCreateNewUser(&userDatabaseGateway)
	return CreateNewUser
}
