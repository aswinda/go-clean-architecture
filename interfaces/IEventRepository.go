package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type IEventRepository interface {
	GetEventDetail(eventId int) (models.EventModel, error)
}
