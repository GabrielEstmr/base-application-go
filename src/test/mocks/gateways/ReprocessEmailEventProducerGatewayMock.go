package test_mocks_gateways

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	test_support_utils "baseapplicationgo/test/support/utils"
	"context"
	"reflect"
)

type ReprocessEmailEventProducerGateway struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewReprocessEmailEventProducerGateway(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *ReprocessEmailEventProducerGateway {
	return &ReprocessEmailEventProducerGateway{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this *ReprocessEmailEventProducerGateway) Send(ctx context.Context, id string) main_domains_exceptions.ApplicationException {
	argsMockSupport := this.cases.GetMethodMock()["Get"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	funcError := test_support_utils.NewErrorUtils().BuildAppErrorFromOutPut(argsMockSupport.GetOutputs()[1])
	if reflect.DeepEqual(id, anyInput1) {
		return funcError
	}
	return nil
}
