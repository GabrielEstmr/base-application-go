package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	test_support_utils "baseapplicationgo/test/support/utils"
	"context"
	"reflect"
)

type CreateEmailBodyMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewCreateEmailBodyMock(methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *CreateEmailBodyMock {
	return &CreateEmailBodyMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this CreateEmailBodyMock) Execute(ctx context.Context,
	emailParams main_domains.EmailParams) ([]byte, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["Execute"]
	anyInput := argsMockSupport.GetInputs()[1]
	bytesOutput := argsMockSupport.GetOutputs()[1].([]byte)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(emailParams, anyInput) {
		return bytesOutput, funcError
	}
	return nil, nil
}
