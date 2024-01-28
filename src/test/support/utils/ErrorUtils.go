package test_support_utils

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"reflect"
)

type ErrorUtils struct {
}

func NewErrorUtils() *ErrorUtils {
	return &ErrorUtils{}
}

func (this ErrorUtils) BuildAppErrorFromOutPut(outPut any) main_domains_exceptions.ApplicationException {
	if outPut == nil {
		return nil
	}

	switch reflect.TypeOf(outPut).String() {
	case reflect.TypeOf(*new(main_domains_exceptions.InternalServerErrorException)).String():
		return outPut.(main_domains_exceptions.InternalServerErrorException)
	case reflect.TypeOf(*new(main_domains_exceptions.ResourceNotFoundException)).String():
		return outPut.(main_domains_exceptions.ResourceNotFoundException)
	case reflect.TypeOf(*new(main_domains_exceptions.BadRequestException)).String():
		return outPut.(main_domains_exceptions.BadRequestException)
	case reflect.TypeOf(*new(main_domains_exceptions.ConflictException)).String():
		return outPut.(main_domains_exceptions.ConflictException)
	case reflect.TypeOf(*new(main_domains_exceptions.UnauthorizedException)).String():
		return outPut.(main_domains_exceptions.UnauthorizedException)
	default:
		return outPut.(main_domains_exceptions.InternalServerErrorException)
	}
}
