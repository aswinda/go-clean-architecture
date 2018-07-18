package services

import (
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type EventService struct {
	interfaces.IEventRepository
}

func (service *EventService) GetEventDetail(eventId int) (models.EventModel, error) {
	result, err := service.EventDetail(eventId)

	return result, err
}
