package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type IEventRepository interface {
	EventDetail(eventId int) (models.EventModel, error)
	StoreEvent(body models.EventModel) (models.EventModel, error)
}
