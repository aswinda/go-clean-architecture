package services

import (
	"encoding/json"

	"github.com/aswinda/loket-backend-test/interfaces"
)

type EventService struct {
	interfaces.IEventRepository
}

func (service *EventService) GetEventDetail(eventId int) (string, error) {
	result, err := service.EventDetail(1)
	if err != nil {
		// handle error
	}

	resp, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(resp), nil
}
