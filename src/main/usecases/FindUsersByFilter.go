package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"log/slog"
)

const _MSG_FIND_USER_BY_FILTER_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUsersByFilter struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	apLog               *slog.Logger
	messageUtils        main_utils_messages.ApplicationMessages
}

func NewFindUsersByFilter(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *FindUsersByFilter {
	return &FindUsersByFilter{
		userDatabaseGateway,
		main_configs_logs.GetLogConfigBean(),
		*main_utils_messages.NewApplicationMessages(),
	}
}

func (this *FindUsersByFilter) Execute(
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	page, err := this.userDatabaseGateway.FindByFilter(filter, pageable)
	if err != nil {
		return main_domains.Page{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_FIND_USER_BY_FILTER_ARCH_ISSUE))
	}
	return page, nil
}
