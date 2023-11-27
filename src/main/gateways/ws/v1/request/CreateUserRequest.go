package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"fmt"
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

func (this *CreateUserRequest) Validate() main_domains_exceptions.ApplicationException {

	val := reflect.ValueOf(this).Elem()
	name := val.Type().Field(1).Name
	log.Println(name)

	var messages []string

	if this.DocumentNumber == "" {
		messages = append(messages, fmt.Sprintf("%s: Field must not be empty", name))
	}

	name2 := val.Type().Field(2).Name
	if this.Name == "" {
		messages = append(messages, fmt.Sprintf("%s: Field must not be empty", name2))
	}

	if len(messages) > 0 {
		return main_domains_exceptions.NewBadRequestException(messages)
	}
	return nil
}
