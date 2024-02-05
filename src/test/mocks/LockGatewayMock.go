package test_mocks

import (
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
