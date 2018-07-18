package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type IEventService interface {
	GetEventDetail(eventId int) (models.EventModel, error)
}
