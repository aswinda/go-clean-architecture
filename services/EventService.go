package services

import (
	"github.com/aswinda/loket-backend-test/interfaces"
)

type EventService struct {
	interfaces.IEventRepository
}

func (service *EventService) GetEventDetail(eventId int) (string, error) {
	result, err := service.GetEventDetail(eventId)
	if err != nil {
		// handle error
	}

	return result, nil
}
