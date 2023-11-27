package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	gateways "baseapplicationgo/main/gateways"
)

type CreateNewUser struct {
	userDatabaseGateway gateways.UserDatabaseGateway
}

func NewCreateNewUser(userDatabaseGateway gateways.UserDatabaseGateway) *CreateNewUser {
	return &CreateNewUser{userDatabaseGateway}
}

func (this *CreateNewUser) Execute(user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException) {

	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(user.DocumentNumber)
	if !userAlreadyPersisted.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewConflictExceptionSglMsg("DocumentNumber Already Exists")
	}

	idPersistedUser, err := this.userDatabaseGateway.Save(user)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("Failed to Save Document")
	}
	user.Id = idPersistedUser

	return user, nil
}
