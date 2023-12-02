package main_domains_exceptions

import "net/http"

type InternalServerErrorException struct {
	code     int
	messages []string
}

func NewInternalServerErrorException(messages []string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     http.StatusInternalServerError,
		messages: messages,
	}
}

func NewInternalServerErrorExceptionSglMsg(message string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     http.StatusInternalServerError,
		messages: []string{message},
	}
}

func (this *InternalServerErrorException) GetCode() int {
	return this.code
}

func (this *InternalServerErrorException) GetMessages() []string {
	return this.messages
}

func (this *InternalServerErrorException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
