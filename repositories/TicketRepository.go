package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type TicketRepositoryWithCircuitBreaker struct {
	TicketRepository interfaces.ITicketRepository
}

type TicketRepository struct {
	interfaces.IDbHandler
}

func (repository *TicketRepositoryWithCircuitBreaker) StoreTicket(body models.TicketModel) (models.TicketModel, error) {
	output := make(chan models.TicketModel, 1)
	hystrix.ConfigureCommand("store_ticket", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("store_ticket", func() error {
		ticket, _ := repository.TicketRepository.StoreTicket(body)

		output <- ticket
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.TicketModel{}, err
	}
}

func (repository *TicketRepository) StoreTicket(body models.TicketModel) (models.TicketModel, error) {

	queryString := fmt.Sprintf("insert into tickets (event_id, ticket_type_id, quota, price, status) values ('%d', '%d', '%d', '%d', '%d')", body.EventID, body.TicketTypeID, body.Quota, body.Price, 1)
	id, err := repository.Execute(queryString)

	if err != nil {
		return models.TicketModel{}, err
	}

	queryString = fmt.Sprintf("SELECT * FROM tickets WHERE id = '%d'", id)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.TicketModel{}, err
	}

	var ticket models.TicketModel

	row.Next()
	row.Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.TicketTypeID,
		&ticket.Quota,
		&ticket.Price,
		&ticket.Status,
		&ticket.CreatedAt,
		&ticket.UpdatedAt)

	return ticket, nil
}

func (repository *TicketRepositoryWithCircuitBreaker) TicketDetail(ticketId int) (models.TicketModel, error) {
	output := make(chan models.TicketModel, 1)
	hystrix.ConfigureCommand("get_ticket", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_ticket", func() error {
		ticket, _ := repository.TicketRepository.TicketDetail(ticketId)

		output <- ticket
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.TicketModel{}, err
	}
}

func (repository *TicketRepository) TicketDetail(ticketId int) (models.TicketModel, error) {
	queryString := fmt.Sprintf("SELECT * FROM tickets WHERE id = '%d'", ticketId)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.TicketModel{}, err
	}

	var ticket models.TicketModel

	row.Next()
	row.Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.TicketTypeID,
		&ticket.Quota,
		&ticket.Price,
		&ticket.Status,
		&ticket.CreatedAt,
		&ticket.UpdatedAt)

	return ticket, nil
}
