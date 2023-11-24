package main_usecases_beans

import (
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_beans "baseapplicationgo/main/usecases/beans/Factories"
	"sync"
)

var once sync.Once

var usecases UsecaseBeans
var usecasesBeans *UsecaseBeans

type UsecaseBeans struct {
	CreateNewUser main_usecases.CreateNewUser
}

func GetControllerBeans() *UsecaseBeans {
	once.Do(func() {
		if usecasesBeans == nil {
			usecases = SubscriptUsecaseBeans()
			usecasesBeans = &usecases
		}
	})
	return usecasesBeans
}

func SubscriptUsecaseBeans() UsecaseBeans {
	return UsecaseBeans{
		CreateNewUser: main_usecases_beans.CreateNewUserBean(),
	}
}
