package services

import (
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type TransactionService struct {
	interfaces.ITransactionRepository
}

func (service *TransactionService) GetTransactionDetail(transactionId int) (models.TransactionModel, error) {
	result, err := service.TransactionDetail(transactionId)

	return result, err
}

func (service *TransactionService) CreateTransaction(body models.TransactionModel) (models.TransactionModel, error) {
	result, err := service.StoreTransaction(body)

	return result, err
}
