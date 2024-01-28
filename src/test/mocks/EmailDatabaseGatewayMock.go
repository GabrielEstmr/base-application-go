package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	test_support_utils "baseapplicationgo/test/support/utils"
	"golang.org/x/net/context"
	"reflect"
)

type EmailDatabaseGatewayMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewEmailDatabaseGatewayMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *EmailDatabaseGatewayMock {
	return &EmailDatabaseGatewayMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this EmailDatabaseGatewayMock) Save(ctx context.Context, email main_domains.Email) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["Save"]
	anyInput := argsMockSupport.GetInputs()[1]
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Email)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(email, anyInput) {
		return emailResp, funcError
	}
	return main_domains.Email{}, nil
}
func (this EmailDatabaseGatewayMock) FindById(ctx context.Context, id string) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["FindById"]
	anyInput := argsMockSupport.GetInputs()[1]
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Email)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(id, anyInput) {
		return emailResp, funcError
	}
	return main_domains.Email{}, nil
}
func (this EmailDatabaseGatewayMock) FindByEventId(ctx context.Context, eventId string) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["FindByEventId"]
	anyInput := argsMockSupport.GetInputs()[1]
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Email)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(eventId, anyInput) {
		return emailResp, funcError
	}
	return main_domains.Email{}, nil
}
func (this EmailDatabaseGatewayMock) Update(ctx context.Context, email main_domains.Email) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["Update"]
	anyInput := argsMockSupport.GetInputs()[1].(main_domains.Email)
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Email)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(email, anyInput) {
		return emailResp, funcError
	}
	return main_domains.Email{}, nil
}

func (this EmailDatabaseGatewayMock) FindByFilter(ctx context.Context, filter main_domains.FindEmailFilter, pageable main_domains.Pageable) (
	main_domains.Page, main_domains_exceptions.ApplicationException) {
	argsMockSupport := this.cases.GetMethodMock()["FindByFilter"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	anyInput2 := argsMockSupport.GetInputs()[2]
	emailResp := argsMockSupport.GetOutputs()[1].(main_domains.Page)
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(filter, anyInput1) && reflect.DeepEqual(pageable, anyInput2) {
		return emailResp, funcError
	}
	return main_domains.Page{}, nil
}
