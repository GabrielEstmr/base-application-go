package main_usecases

import (
	domains "baseapplicationgo/main/domains"
	gateways "baseapplicationgo/main/gateways"
	main_utils "baseapplicationgo/main/utils"
	"time"
)

// *
// serve para atribuir valor REAL de um ponteiro
// serve para falar que uma var é um ponteiro

// &
// serve para pegar referencia em memoria de uma var

type CreateNewUser struct {
	userDatabaseGateway gateways.UserDatabaseGateway
}

//func NewCreateNewUser(accountDatabaseGateway *gateways.AccountDatabaseGateway) *CreateNewUser {
//	return &CreateNewUser{*accountDatabaseGateway}
//}

func NewCreateNewUser(userDatabaseGateway *gateways.UserDatabaseGateway) *CreateNewUser {
	return &CreateNewUser{*userDatabaseGateway}
}

func (this *CreateNewUser) Execute(name string, documentNumber string, birthday time.Time) (string, error) {

	user := domains.User{
		Name:           name,
		DocumentNumber: documentNumber,
		Birthday:       birthday,
	}

	save, err := this.userDatabaseGateway.Save(user)
	main_utils.FailOnError(err, "oiasokf")

	return save, nil
}
