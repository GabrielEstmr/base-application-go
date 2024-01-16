package main_gateways_rabbitmq_resources

import "time"

type Event struct {
	Message   any       `json:"message"`
	EventDate time.Time `json:"eventDate"`
}

func NewEvent(
	message any,
) *Event {
	return &Event{
		Message:   message,
		EventDate: time.Now()}
}
