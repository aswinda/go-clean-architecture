package services

import (
	"errors"
	"fmt"

	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type TransactionService struct {
	interfaces.ITransactionRepository
	interfaces.ITicketRepository
}

func (service *TransactionService) GetTransactionDetail(transactionId int) (models.TransactionModel, error) {
	result, err := service.TransactionDetail(transactionId)

	return result, err
}

func (service *TransactionService) CreateTransaction(body models.TransactionModel) (models.TransactionModel, error) {
	fmt.Println(body)
	// result, err := service.StoreTransaction(body)
	// for _, element := range [3]string{"a", "b", "c"} {
	//     fmt.Print(element)
	// }
	return models.TransactionModel{}, nil
}

func (service *TransactionService) PurchaseTransaction(body models.TransactionPurchase) (models.TransactionModel, error) {

	prices := make([]int, len(body.TicketID))
	totalPrice := 0
	totalAmount := 0
	for i := 0; i < len(body.TicketID); i++ {
		// get ticket quota
		result, _ := service.TicketDetail(body.TicketID[i])

		// check event id
		if body.EventID != result.EventID {
			return models.TransactionModel{}, errors.New("Ticket not the same event id")
		}

		// check quota
		if result.Quota < body.Amount[i] {
			return models.TransactionModel{}, errors.New(fmt.Sprintf("Ticket ID '%d' : runs out", body.TicketID[i]))
		}

		// pricing
		prices[i] = result.Price
		totalAmount += body.Amount[i]
		totalPrice += result.Price * body.Amount[i]
	}

	//  store transaction
	transactionModel := models.TransactionModel{
		CustomerID:  body.CustomerID,
		EventID:     body.EventID,
		TotalAmount: totalAmount,
		TotalPrice:  totalPrice}

	result, err := service.StoreTransaction(transactionModel)

	if err != nil {
		return models.TransactionModel{}, err
	}

	// store transaction detail
	for i := 0; i < len(body.TicketID); i++ {
		transactionDetailModel := models.TransactionDetailModel{
			TicketID:      body.TicketID[i],
			TransactionID: result.ID,
			Amount:        body.Amount[i],
			TotalPrice:    prices[i]}

		_, err := service.StoreTransactionDetail(transactionDetailModel)
		if err != nil {
			return models.TransactionModel{}, err
		}
	}

	return result, nil
}
