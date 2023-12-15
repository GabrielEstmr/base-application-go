package main_usecases

import (
	main_configs_ff "baseapplicationgo/main/configs/ff"
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"log"
	"log/slog"
)

const _MSG_FIND_USER_BY_ID_DOC_NOT_FOUND = "find.user.user.not.found"
const _MSG_FIND_USER_BY_ID_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUserById struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	apLog               *slog.Logger
	messageUtils        main_utils_messages.ApplicationMessages
	featuresGateway     main_gateways.FeaturesGateway
}

func NewFindUserById(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	featuresGateway main_gateways.FeaturesGateway,
) *FindUserById {
	return &FindUserById{
		userDatabaseGateway,
		main_configs_logs.GetLogConfigBean(),
		*main_utils_messages.NewApplicationMessages(),
		featuresGateway,
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

	disabled, err := this.featuresGateway.IsDisabled(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)

	if disabled == true && err == nil {
		log.Println("====================> FEATURE")
	}

	_, errF := this.featuresGateway.Enable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errF != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errF.Error())
	}

	_, errD := this.featuresGateway.Disable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errD != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errD.Error())
	}

	_, errF2 := this.featuresGateway.Enable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errF2 != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errF2.Error())
	}

	_, errD2 := this.featuresGateway.Disable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errD2 != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errD2.Error())
	}

	_, errF3 := this.featuresGateway.Enable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errF3 != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errF3.Error())
	}

	_, errD3 := this.featuresGateway.Disable(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)
	if errD3 != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errD3.Error())
	}

	enabled, errEEE := this.featuresGateway.IsEnabled(main_configs_ff.ENABLE_FIND_BY_ID_ENDPOINT)

	if enabled == true && errEEE == nil {
		log.Println("====================> FEATURE")
	}

	return user, nil
}
