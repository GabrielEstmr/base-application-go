package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	test_support_utils "baseapplicationgo/test/support/utils"
	"golang.org/x/net/context"
	"reflect"
)

type CreateEmailBodySendAndPersistAsSentMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewCreateEmailBodySendAndPersistAsSentMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *CreateEmailBodySendAndPersistAsSentMock {
	return &CreateEmailBodySendAndPersistAsSentMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this CreateEmailBodySendAndPersistAsSentMock) Execute(
	ctx context.Context,
	email main_domains.Email,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["Execute"]
	anyInput := argsMockSupport.GetInputs()[1]
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Email)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(email, anyInput) {
		return emailResp, funcError
	}
	return main_domains.Email{}, nil
}
