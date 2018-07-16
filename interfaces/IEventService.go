package interfaces

type IEventService interface {
	GetEventDetail(eventId int) (string, error)
}
