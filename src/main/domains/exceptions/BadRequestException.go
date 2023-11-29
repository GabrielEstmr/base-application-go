package main_domains_exceptions

import (
	"net/http"
)

type BadRequestException struct {
	code     int
	messages []string
}

func NewBadRequestException(messages []string) *BadRequestException {
	return &BadRequestException{
		code:     http.StatusBadRequest,
		messages: messages,
	}
}

func NewBadRequestExceptionSglMsg(message string) *BadRequestException {
	return &BadRequestException{
		code:     http.StatusConflict,
		messages: []string{message},
	}
}

func (this *BadRequestException) GetCode() int {
	return this.code
}

func (this *BadRequestException) GetMessages() []string {
	return this.messages
}

func (this *BadRequestException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
