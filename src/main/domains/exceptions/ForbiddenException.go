package main_domains_exceptions

import "net/http"

type ForbiddenException struct {
	code     int
	messages []string
}

const STATUS_FORBIDDEN = http.StatusForbidden

func NewForbiddenException(messages []string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     STATUS_FORBIDDEN,
		messages: messages,
	}
}

func NewForbiddenExceptionSglMsg(message string) *InternalServerErrorException {
	return &InternalServerErrorException{
		code:     STATUS_FORBIDDEN,
		messages: []string{message},
	}
}

func (this ForbiddenException) GetCode() int {
	return this.code
}

func (this ForbiddenException) GetMessages() []string {
	return this.messages
}

func (this ForbiddenException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
