package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type EventRepositoryWithCircuitBreaker struct {
	EventRepository interfaces.IEventRepository
}

type EventRepository struct {
	interfaces.IDbHandler
}

func (repository *EventRepositoryWithCircuitBreaker) EventDetail(eventId int) (models.EventModel, error) {
	output := make(chan models.EventModel, 1)
	hystrix.ConfigureCommand("get_event", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_event", func() error {
		event, _ := repository.EventRepository.EventDetail(eventId)

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

func (repository *EventRepository) EventDetail(eventId int) (models.EventModel, error) {
	queryString := fmt.Sprintf("SELECT * FROM events WHERE id = '%d'", eventId)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.EventModel{}, err
	}

	var event models.EventModel

	row.Next()
	row.Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.Start_time,
		&event.End_time,
		&event.Status,
		&event.Created_at,
		&event.Updated_at)

	return event, nil
}
