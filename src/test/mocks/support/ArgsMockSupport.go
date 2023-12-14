package test_mocks_support

type ArgsMockSupport struct {
	inputs  map[int]any
	outputs map[int]any
}

func NewArgsMockSupport() *ArgsMockSupport {
	return &ArgsMockSupport{
		inputs:  make(map[int]any),
		outputs: make(map[int]any)}
}

func NewArgsMockSupportAllArgs(
	inputs map[int]any,
	outputs map[int]any) *ArgsMockSupport {
	return &ArgsMockSupport{
		inputs:  inputs,
		outputs: outputs}
}

func (this *ArgsMockSupport) GetInputs() map[int]any {
	return this.inputs
}

func (this *ArgsMockSupport) GetOutputs() map[int]any {
	return this.outputs
}

func (this *ArgsMockSupport) AddInput(inputIndex int, inputValue any) *ArgsMockSupport {
	this.inputs[inputIndex] = inputValue
	return this
}

func (this *ArgsMockSupport) AddOutput(outputIndex int, outputValue any) *ArgsMockSupport {
	this.outputs[outputIndex] = outputValue
	return this
}
