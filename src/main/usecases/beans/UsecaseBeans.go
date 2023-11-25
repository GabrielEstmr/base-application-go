package main_usecases_beans

import (
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_beans_factories "baseapplicationgo/main/usecases/beans/Factories"
	"sync"
)

var once sync.Once

var usecaseBeans *UsecaseBeans

type UsecaseBeans struct {
	CreateNewUser main_usecases.CreateNewUser
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
		CreateNewUser: *main_usecases_beans_factories.CreateNewUserBean(),
	}
}
