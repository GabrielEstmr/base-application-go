package test_mocks

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
)

type CreateEmailBodyFactoryMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewCreateEmailBodyFactoryMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *CreateEmailBodyFactoryMock {
	return &CreateEmailBodyFactoryMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this CreateEmailBodyFactoryMock) Get(ctx context.Context,
	emailType main_domains_enums.EmailTemplateType) main_usecases_interfaces.CreateEmailBody {
	argsMockSupport := this.cases.GetMethodMock()["Get"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	funcError := argsMockSupport.GetOutputs()[1].(main_usecases_interfaces.CreateEmailBody)
	if reflect.DeepEqual(emailType, anyInput1) {
		return funcError
	}
	return nil
}
