package main_gateways_rabbitmq_resources

import (
	"encoding/json"
	"github.com/google/uuid"
)

type EventUUID struct {
	Event   `json:"event"`
	EventId string `json:"eventId"`
}

func NewEventUUID(
	message any,
) (*EventUUID, error) {
	eventId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &EventUUID{
		Event:   *NewEvent(message),
		EventId: eventId.String(),
	}, nil
}

func (this *EventUUID) FromJSON(data []byte) (EventUUID, error) {
	var event EventUUID
	if errU := json.Unmarshal(data, &event); errU != nil {
		return *new(EventUUID), errU
	}
	return event, nil
}

func (this *EventUUID) GetMessageJSON(data []byte) ([]byte, error) {
	event, errJ := this.Event.FromJSON(data)
	if errJ != nil {
		return nil, errJ
	}
	msgBytes, errJson := json.Marshal(event.Message)
	if errJson != nil {
		return nil, errJson
	}
	return msgBytes, nil
}
