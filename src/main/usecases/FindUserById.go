package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"fmt"
	"log/slog"
)

type FindUserById struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	apLog               *slog.Logger
}

func NewFindUserById(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *FindUserById {
	return &FindUserById{
		userDatabaseGateway,
		main_configs_logs.GetLogConfigBean(),
	}
}

func (this *FindUserById) Execute(id string) (main_domains.User, main_domains_exceptions.ApplicationException) {

	this.apLog.Debug(fmt.Sprintf("Finding User by id: %s", id))

	user, err := this.userDatabaseGateway.FindById(id)
	if err != nil {
		return main_domains.User{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg("Failed to Find Document")
	}
	if user.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
			main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(_MSG_LEY_DOC_ALREADY_EXISTS))
	}
	return user, nil
}
