package main_usecases

import (
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"log/slog"
)

const _MSG_FIND_USER_BY_FILTER_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUsersByFilter struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	apLog               *slog.Logger
}

func NewFindUsersByFilter(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *FindUsersByFilter {
	return &FindUsersByFilter{
		userDatabaseGateway,
		main_configs_logs.GetLogConfigBean(),
	}
}

func (this *FindUsersByFilter) Execute(
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	page, err := this.userDatabaseGateway.FindByFilter(filter, pageable)
	if err != nil {
		return main_domains.Page{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(
					_MSG_FIND_USER_BY_FILTER_ARCH_ISSUE))
	}
	return page, nil
}
