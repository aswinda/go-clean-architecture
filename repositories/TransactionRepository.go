package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type TransactionRepositoryWithCircuitBreaker struct {
	TransactionRepository interfaces.ITransactionRepository
}

type TransactionRepository struct {
	interfaces.IDbHandler
}

func (repository *TransactionRepositoryWithCircuitBreaker) TransactionDetail(transactionId int) (models.TransactionModel, error) {
	output := make(chan models.TransactionModel, 1)
	hystrix.ConfigureCommand("get_transaction", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_transaction", func() error {
		transaction, _ := repository.TransactionRepository.TransactionDetail(transactionId)

		output <- transaction
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.TransactionModel{}, err
	}
}

func (repository *TransactionRepository) TransactionDetail(transactionId int) (models.TransactionModel, error) {
	queryString := fmt.Sprintf("SELECT * FROM transactions WHERE id = '%d'", transactionId)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.TransactionModel{}, err
	}

	var transaction models.TransactionModel

	row.Next()
	row.Scan(
		&transaction.ID,
		&transaction.CustomerID,
		&transaction.EventID,
		&transaction.TotalPrice,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt)

	return transaction, nil
}

func (repository *TransactionRepositoryWithCircuitBreaker) StoreTransaction(body models.TransactionModel) (models.TransactionModel, error) {
	output := make(chan models.TransactionModel, 1)
	hystrix.ConfigureCommand("store_transaction", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("store_transaction", func() error {
		transaction, _ := repository.TransactionRepository.StoreTransaction(body)

		output <- transaction
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.TransactionModel{}, err
	}
}

func (repository *TransactionRepository) StoreTransaction(body models.TransactionModel) (models.TransactionModel, error) {

	queryString := fmt.Sprintf("insert into transactions ( customer_id, event_id, total_price, status) values ('%d', '%d', '%d', '%d')", body.CustomerID, body.EventID, body.TotalPrice, 1)
	id, err := repository.Execute(queryString)

	if err != nil {
		return models.TransactionModel{}, err
	}

	queryString = fmt.Sprintf("SELECT * FROM transactions WHERE id = '%d'", id)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.TransactionModel{}, err
	}

	var transaction models.TransactionModel

	row.Next()
	row.Scan(
		&transaction.ID,
		&transaction.CustomerID,
		&transaction.EventID,
		&transaction.TotalPrice,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt)

	return transaction, nil
}
