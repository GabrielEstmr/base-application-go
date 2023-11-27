package main_domains_exceptions

type ResourceNotFoundException struct {
	code    int
	message string
}

func (this *ResourceNotFoundException) New(code int, message string) *ResourceNotFoundException {
	this.code = code
	this.message = message
	return this
}

func (this *ResourceNotFoundException) GetCode() int {
	return this.code
}

func (this *ResourceNotFoundException) Error() string {
	return this.message
}
