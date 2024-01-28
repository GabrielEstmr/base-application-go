package test_mocks

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
)

type SendEmailGatewayFactoryMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewSendEmailGatewayFactoryMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *SendEmailGatewayFactoryMock {
	return &SendEmailGatewayFactoryMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this SendEmailGatewayFactoryMock) Get(ctx context.Context,
	emailType main_domains_enums.EmailTemplateType) main_gateways.EmailGateway {
	argsMockSupport := this.cases.GetMethodMock()["Get"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	funcError := argsMockSupport.GetOutputs()[1].(main_gateways.EmailGateway)
	if reflect.DeepEqual(emailType, anyInput1) {
		return funcError
	}
	return nil
}
