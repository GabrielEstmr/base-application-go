package main_gateways_rabbitmq_resources

import (
	"encoding/json"
	"time"
)

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

func (this *Event) FromJSON(data []byte) (Event, error) {
	var event Event
	if errU := json.Unmarshal(data, &event); errU != nil {
		return *new(Event), errU
	}
	return event, nil
}

func (this *Event) GetMessageJSON(data []byte) ([]byte, error) {
	event, errJ := this.FromJSON(data)
	if errJ != nil {
		return nil, errJ
	}
	msgBytes, errJson := json.Marshal(event.Message)
	if errJson != nil {
		return nil, errJson
	}
	return msgBytes, nil
}
