package test_mocks_support

type MethodArgsMockSupport struct {
	methodMock map[string]ArgsMockSupport
}

func NewMethodArgsMockSupport(
	methodName map[string]ArgsMockSupport) *MethodArgsMockSupport {
	return &MethodArgsMockSupport{
		methodMock: methodName}
}

func (this MethodArgsMockSupport) GetMethodMock() map[string]ArgsMockSupport {
	return this.methodMock
}
