package main_domains_exceptions

import "net/http"

type UnauthorizedException struct {
	code     int
	messages []string
}

const STATUS_UNAUTHORIZED = http.StatusUnauthorized

func NewUnauthorizedException(messages []string) *UnauthorizedException {
	return &UnauthorizedException{
		code:     STATUS_UNAUTHORIZED,
		messages: messages,
	}
}

func NewUnauthorizedExceptionSglMsg(message string) *UnauthorizedException {
	return &UnauthorizedException{
		code:     STATUS_UNAUTHORIZED,
		messages: []string{message},
	}
}

func (this UnauthorizedException) GetCode() int {
	return this.code
}

func (this UnauthorizedException) GetMessages() []string {
	return this.messages
}

func (this UnauthorizedException) Error() string {
	var message string
	for _, value := range this.messages {
		message = message + value
	}
	return message
}
