package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITransactionRepository interface {
	TransactionDetail(transactionId int) (models.TransactionModel, error)
	StoreTransaction(body models.TransactionModel) (models.TransactionModel, error)
}