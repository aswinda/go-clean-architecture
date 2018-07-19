package services

import (
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type TicketService struct {
	interfaces.ITicketRepository
}

func (service *TicketService) CreateTicket(body models.TicketModel) (models.TicketModel, error) {
	result, err := service.StoreTicket(body)

	return result, err
}

func (service *TicketService) GetTicketDetail(ticketId int) (models.TicketModel, error) {
	result, err := service.TicketDetail(ticketId)

	return result, err
}
