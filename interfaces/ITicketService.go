package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITicketService interface {
	CreateTicket(body models.TicketModel) (models.TicketModel, error)
	GetTicketDetail(ticketId int) (models.TicketModel, error)
}
