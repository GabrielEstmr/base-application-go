package test_mocks

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"baseapplicationgo/main/domains/lock"
	test_mocks_support "baseapplicationgo/test/mocks/support"
	"context"
	"reflect"
	"time"
)

type LockGatewayMock struct {
	cases test_mocks_support.MethodArgsMockSupport
}

func NewLockGatewayMock(
	methodsMock map[string]test_mocks_support.ArgsMockSupport,
) *LockGatewayMock {
	return &LockGatewayMock{
		cases: *test_mocks_support.NewMethodArgsMockSupport(methodsMock),
	}
}

func (this LockGatewayMock) Get(
	ctx context.Context,
	key string,
	ttl time.Duration) *lock.SingleLock {
	argsMockSupport := this.cases.GetMethodMock()["SendMail"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	anyInput2 := argsMockSupport.GetInputs()[2]
	outPut1 := argsMockSupport.GetOutputs()[1]
	if reflect.DeepEqual(key, anyInput1) && reflect.DeepEqual(ttl, anyInput2) {
		return outPut1.(*lock.SingleLock)
	}
	return nil
}

func (this LockGatewayMock) GetWithScope(
	ctx context.Context,
	scope main_domains_enums.LockScope,
	key string,
	ttl time.Duration) *lock.SingleLock {
	argsMockSupport := this.cases.GetMethodMock()["SendMail"]
	anyInput1 := argsMockSupport.GetInputs()[1]
	anyInput2 := argsMockSupport.GetInputs()[2].(main_domains_enums.LockScope)
	anyInput3 := argsMockSupport.GetInputs()[3]
	outPut1 := argsMockSupport.GetOutputs()[1]
	if reflect.DeepEqual(key, anyInput1) && reflect.DeepEqual(scope, anyInput2) && reflect.DeepEqual(ttl, anyInput3) {
		return outPut1.(*lock.SingleLock)
	}
	return nil
}
