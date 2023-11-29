package main_domains_exceptions

import "net/http"

type ResourceNotFoundException struct {
	code     int
	messages []string
}

func NewResourceNotFoundException(messages []string) *ResourceNotFoundException {
	return &ResourceNotFoundException{
		code:     http.StatusNotFound,
		messages: messages,
	}
}

func NewResourceNotFoundExceptionSglMsg(message string) *ResourceNotFoundException {
	return &ResourceNotFoundException{
		code:     http.StatusConflict,
		messages: []string{message},
	}
}

func (this *ResourceNotFoundException) GetCode() int {
	return this.code
}

func (this *ResourceNotFoundException) GetMessages() []string {
	return this.messages
}

func (this *ResourceNotFoundException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
