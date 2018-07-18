package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITransactionService interface {
	GetTransactionDetail(transactionId int) (models.TransactionModel, error)
	CreateTransaction(body models.TransactionModel) (models.TransactionModel, error)
}
