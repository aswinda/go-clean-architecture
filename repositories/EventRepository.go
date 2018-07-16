package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/notifyme/application/api/interfaces"
	"github.com/aswinda/notifyme/application/api/models"
)

type EventRepositoryWithCircuitBreaker struct {
	EventRepository interfaces.IEventRepository
}

type EventRepository struct {
	interfaces.IDbHandler
}

func (repository *EventRepositoryWithCircuitBreaker) GetEventDetail(eventId int) (models.EventModel, error) {
	output := make(chan models.EventModel, 1)
	hystrix.ConfigureCommand("get_event", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_event", func() error {
		event, _ := repository.EventRepository.GetEventDetail(eventId)

		output <- event
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.EventModel{}, err
	}
}

func (repository *EventRepository) GetEventDetail(eventId int) (models.EventModel, error) {
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM events WHERE id = '%d'", eventId))

	if err != nil {
		return models.EventModel{}, err
	}

	var event models.EventModel

	row.Next()
	row.Scan(&event.Id, &event.Name, &event.Age)

	return event, nil
}
