package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/matryer/respond.v1"
)

type TransactionController struct {
	interfaces.ITransactionService
}

func (controller *TransactionController) GetTransactionDetailAction(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	transactionIdS := queryValues.Get("transaction_id")
	transactionId, _ := strconv.Atoi(transactionIdS)

	validate = validator.New()
	err := validate.Var(transactionId, "required,numeric")

	if err != nil {
		errs := err.(validator.ValidationErrors)
		str := fmt.Sprintf("%s", errs)
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, str, map[string]interface{}{}))
		return
	}

	detail, err := controller.GetTransactionDetail(transactionId)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, fmt.Sprintf("%s", err), map[string]interface{}{}))
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(detail)
	json.Unmarshal(inrec, &inInterface)

	respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusOK, "", inInterface))
}

func (controller *TransactionController) CreateTransactionAction(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var req models.TransactionModel
	for {
		if err := dec.Decode(&req); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	validate = validator.New()

	err := validate.Struct(req)

	if err != nil {

		errs := err.(validator.ValidationErrors)
		str := fmt.Sprintf("%s", errs)

		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, str, map[string]interface{}{}))
		return
	}

	result, err := controller.CreateTransaction(req)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, fmt.Sprintf("%s", err), map[string]interface{}{}))
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(result)
	json.Unmarshal(inrec, &inInterface)

	respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusOK, "", inInterface))
}
