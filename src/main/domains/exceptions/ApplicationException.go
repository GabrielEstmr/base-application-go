package main_domains_exceptions

type ApplicationException interface {
	GetCode() int
	Error() string
}
