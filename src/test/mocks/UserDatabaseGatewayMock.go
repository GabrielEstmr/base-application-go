package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
)

type UserDatabaseGatewayMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewUserDatabaseGatewayMock(methodsMock map[string]test_mocks_support.ArgsMockSupport) *UserDatabaseGatewayMock {
	return &UserDatabaseGatewayMock{cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock)}
}

func (this *UserDatabaseGatewayMock) Save(ctx context.Context, user main_domains.User) (main_domains.User, error) {
	argsMockSupport := this.cases.GetMethodMock()["Save"]
	anyInput := argsMockSupport.GetInputs()[1]
	userResp := argsMockSupport.GetOutputs()[1].(main_domains.User)
	funcError := buildError(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(user, anyInput) {
		return userResp, funcError
	}
	return user, nil
}

func buildError(outPut any) error {
	if outPut == nil {
		return nil
	}
	funcError := outPut.(error)
	return funcError
}

func (this *UserDatabaseGatewayMock) FindById(ctx context.Context, id string) (main_domains.User, error) {
	argsMockSupport := this.cases.GetMethodMock()["FindById"]
	anyInput := argsMockSupport.GetInputs()[1]
	userResp := argsMockSupport.GetOutputs()[1].(main_domains.User)
	funcError := buildError(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(id, anyInput) {
		return userResp, funcError
	}
	return main_domains.User{}, nil
}

func (this *UserDatabaseGatewayMock) FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error) {
	argsMockSupport := this.cases.GetMethodMock()["FindByDocumentNumber"]
	anyInput := argsMockSupport.GetInputs()[1]
	userResp := argsMockSupport.GetOutputs()[1].(main_domains.User)
	funcError := buildError(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(documentNumber, anyInput) {
		return userResp, funcError
	}
	return main_domains.User{}, nil
}

func (this *UserDatabaseGatewayMock) FindByFilter(
	ctx context.Context, filter main_domains.FindUserFilter,
	pageable main_domains.Pageable) (main_domains.Page, error) {
	argsMockSupport := this.cases.GetMethodMock()["FindByFilter"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	anyInput2 := argsMockSupport.GetInputs()[2]
	userResp := argsMockSupport.GetOutputs()[1].(main_domains.Page)
	funcError := buildError(argsMockSupport.GetOutputs()[2])
	if reflect.DeepEqual(filter, anyInput1) && reflect.DeepEqual(pageable, anyInput2) {
		return userResp, funcError
	}
	return main_domains.Page{}, nil
}
