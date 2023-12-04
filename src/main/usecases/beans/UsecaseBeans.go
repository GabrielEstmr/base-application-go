package main_usecases_beans

import (
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_beans_factories "baseapplicationgo/main/usecases/beans/Factories"
	"sync"
)

var once sync.Once

var usecaseBeans *UsecaseBeans

type UsecaseBeans struct {
	CreateNewUser     *main_usecases.CreateNewUser
	FindUserById      *main_usecases.FindUserById
	FindUsersByFilter *main_usecases.FindUsersByFilter
}

func GetUsecaseBeans() *UsecaseBeans {
	once.Do(func() {
		if usecaseBeans == nil {
			usecaseBeans = subscriptUsecaseBeans()
		}
	})
	return usecaseBeans
}

func subscriptUsecaseBeans() *UsecaseBeans {
	return &UsecaseBeans{
		CreateNewUser:     main_usecases_beans_factories.NewCreateNewUserBean().Get(),
		FindUserById:      main_usecases_beans_factories.NewFindUserByIdBean().Get(),
		FindUsersByFilter: main_usecases_beans_factories.NewFindUsersByFilterBean().Get(),
	}
}
