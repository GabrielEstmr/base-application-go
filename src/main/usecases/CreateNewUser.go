package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	gateways "baseapplicationgo/main/gateways"
	main_utils "baseapplicationgo/main/utils"
)

type CreateNewUser struct {
	userDatabaseGateway gateways.UserDatabaseGateway
}

func NewCreateNewUser(userDatabaseGateway gateways.UserDatabaseGateway) *CreateNewUser {
	return &CreateNewUser{userDatabaseGateway}
}

func (this *CreateNewUser) Execute(user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException) {

	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(user.DocumentNumber)

	if (userAlreadyPersisted != main_domains.User{}) {
		return main_domains.User{}, &main_domains_exceptions.ConflictException{Code: 409, Message: "ijasdjijajifj"}
	}

	idPersistedUser, err := this.userDatabaseGateway.Save(user)
	main_utils.FailOnError(err, user.Id)
	user.Id = idPersistedUser

	return user, nil
}
