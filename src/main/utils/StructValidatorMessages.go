package main_utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type StructValidatorMessages struct {
	messages []string
}

func (s *StructValidatorMessages) GetMessages() []string {
	return s.messages
}

func NewStructValidatorMessages(any interface{}) *StructValidatorMessages {
	return &StructValidatorMessages{
		messages: buildMessages(any),
	}
}

func buildMessages(any interface{}) []string {
	var messages []string
	validate := validator.New()
	err := validate.Struct(any)
	validationErrors := buildValidationErrors(err)
	if err != nil {
		for _, element := range validationErrors {
			messages = append(messages, element.Error())
		}
	}
	return messages
}

func buildValidationErrors(err error) validator.ValidationErrors {
	if err == nil {
		return nil
	}
	var validatorErrs validator.ValidationErrors
	errors.As(err, &validatorErrs)
	return validatorErrs
}
