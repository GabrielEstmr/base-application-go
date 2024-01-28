package test_mocks

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	test_support_utils "baseapplicationgo/test/support/utils"
	"context"
	"reflect"
)

type EmailGatewayMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewEmailGatewayMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *EmailGatewayMock {
	return &EmailGatewayMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this EmailGatewayMock) SendMail(
	ctx context.Context,
	to []string,
	body []byte) main_domains_exceptions.ApplicationException {
	argsMockSupport := this.cases.GetMethodMock()["SendMail"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	anyInput2 := argsMockSupport.GetInputs()[2]
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(
		argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(to, anyInput1) && reflect.DeepEqual(body, anyInput2) {
		return funcError
	}
	return nil
}
