package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"fmt"
	"log"
	"log/slog"
)

const _MSG_LEY_DOC_ALREADY_EXISTS = "create.user.user.with.given.document.already.exists"

type CreateNewUser struct {
	userDatabaseGateway      main_gateways.UserDatabaseGateway
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway
	apLog                    *slog.Logger
}

func NewCreateNewUser(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway,
) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway,
		userDatabaseCacheGateway,
		main_configs_logs.GetLogConfigBean(),
	}
}

func (this *CreateNewUser) Execute(user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException) {

	this.apLog.Info(fmt.Sprintf("Creating new User with documentNumber: %s", user.DocumentNumber))
	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(user.DocumentNumber)
	if !userAlreadyPersisted.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewConflictExceptionSglMsg(
			main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(_MSG_LEY_DOC_ALREADY_EXISTS))
	}

	idPersistedUser, err := this.userDatabaseGateway.Save(user)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("Failed to Save Document")
	}
	user.Id = idPersistedUser

	_, errCache := this.userDatabaseCacheGateway.Save(user)
	if errCache != nil {
		log.Println(errCache)
	}

	return user, nil
}
