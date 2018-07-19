package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITicketRepository interface {
	StoreTicket(body models.TicketModel) (models.TicketModel, error)
	TicketDetail(ticketId int) (models.TicketModel, error)
}
