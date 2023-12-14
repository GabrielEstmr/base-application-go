package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"log/slog"
)

const _MSG_FIND_USER_BY_ID_DOC_NOT_FOUND = "find.user.user.not.found"
const _MSG_FIND_USER_BY_ID_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUserById struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	apLog               *slog.Logger
	messageUtils        main_utils_messages.ApplicationMessages
}

func NewFindUserById(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *FindUserById {
	return &FindUserById{
		userDatabaseGateway,
		main_configs_logs.GetLogConfigBean(),
		*main_utils_messages.NewApplicationMessages(),
	}
}

func (this *FindUserById) Execute(id string) (main_domains.User, main_domains_exceptions.ApplicationException) {
	user, err := this.userDatabaseGateway.FindById(id)
	if err != nil {
		return main_domains.User{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_FIND_USER_BY_ID_ARCH_ISSUE)
	}
	if user.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(_MSG_FIND_USER_BY_ID_DOC_NOT_FOUND))
	}
	return user, nil
}
