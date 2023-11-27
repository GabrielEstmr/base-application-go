package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	"log"
	"reflect"
	"time"
)

type CreateUserRequest struct {
	Name           string    `json:"name"`
	DocumentNumber string    `json:"documentNumber"`
	Birthday       time.Time `json:"birthday"`
}

func (this *CreateUserRequest) ToDomain() main_domains.User {
	return main_domains.User{
		Name:           this.Name,
		DocumentNumber: this.DocumentNumber,
		Birthday:       this.Birthday,
	}
}

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (this *ValidationError) Error() string {
	return this.Message
}

func (this *CreateUserRequest) Validate() error {

	var testeError error

	val := reflect.ValueOf(this).Elem()
	name := val.Type().Field(1).Name
	log.Println(name)

	if this.Name == "" {
		testeError = &ValidationError{
			Code:    "uashduhas",
			Message: "field must not be blank" + name,
		}
		return testeError
	}

	return nil
}
