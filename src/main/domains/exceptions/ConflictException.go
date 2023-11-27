package main_domains_exceptions

type ConflictException struct {
	Code    int
	Message string
}

func (this *ConflictException) GetCode() int {
	return this.Code
}

func (this *ConflictException) Error() string {
	return this.Message
}
