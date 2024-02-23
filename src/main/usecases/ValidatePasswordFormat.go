package main_usecases

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"unicode"
)

type ValidatePasswordFormat struct {
	_MSG_PASSWORD_FORMAT_NOT_ACCEPTABLE string
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

func NewValidatePasswordFormat(
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *ValidatePasswordFormat {
	return &ValidatePasswordFormat{
		_MSG_PASSWORD_FORMAT_NOT_ACCEPTABLE: "providers.password.not.acceptable",
		logsMonitoringGateway:               logsMonitoringGateway,
		spanGateway:                         spanGateway,
		messageUtils:                        messageUtils,
	}
}

func (this ValidatePasswordFormat) Execute(
	ctx context.Context,
	password string,
) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "VerifyPasswordFormat-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "VerifyPasswordFormat")

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	result := hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial

	if result == false {
		return main_domains_exceptions.NewBadRequestExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_PASSWORD_FORMAT_NOT_ACCEPTABLE))
	}

	return nil
}
