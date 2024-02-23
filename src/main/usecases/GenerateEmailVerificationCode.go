package main_usecases

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"math/rand"
	"strconv"
)

type GenerateEmailVerificationCode struct {
	_MIN_VALUE_GENERATE   int
	_MAX_VALUE_GENERATE   int
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewGenerateEmailVerificationCode(
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *GenerateEmailVerificationCode {
	return &GenerateEmailVerificationCode{
		_MIN_VALUE_GENERATE:   100000000,
		_MAX_VALUE_GENERATE:   999999999,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageUtils,
	}
}

func (this *GenerateEmailVerificationCode) Execute() string {
	return this.generateRandomEightDigitsNumber()
}

func (this *GenerateEmailVerificationCode) generateRandomEightDigitsNumber() string {
	value := rand.Intn(this._MAX_VALUE_GENERATE-this._MIN_VALUE_GENERATE) + this._MIN_VALUE_GENERATE
	return strconv.Itoa(value)
}
