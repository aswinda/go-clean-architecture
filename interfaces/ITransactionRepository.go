package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITransactionRepository interface {
	TransactionDetail(transactionId int) (models.TransactionModel, error)
	TransactionDetailList(transactionId int) ([]*models.TransactionDetailModel, error)
	StoreTransaction(body models.TransactionModel) (models.TransactionModel, error)
	StoreTransactionDetail(body models.TransactionDetailModel) (models.TransactionDetailModel, error)
}
