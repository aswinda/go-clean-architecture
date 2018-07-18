package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aswinda/loket-backend-test/interfaces"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/matryer/respond.v1"
)

var validate *validator.Validate

type EventController struct {
	interfaces.IEventService
}

func ApiResponse(code int, messages []map[string]interface{}, data map[string]interface{}) map[string]interface{} {
	var status string
	if code == 200 {
		status = "ok"
	} else {
		status = "error"
	}

	result := map[string]interface{}{
		"status":   status,
		"code":     code,
		"messages": messages,
		"data":     data,
	}

	return result
}

func (controller *EventController) GetEventDetailAction(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	eventIdS := queryValues.Get("event_id")
	eventId, _ := strconv.Atoi(eventIdS)

	validate = validator.New()
	err := validate.Var(eventId, "required,numeric")

	if err != nil {
		var errInterface map[string]interface{}
		inrec, _ := json.Marshal(err)
		json.Unmarshal(inrec, &errInterface)
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, []map[string]interface{}{errInterface}, map[string]interface{}{}))
		// fmt.Println(err) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
		}
		return
	}

	detail, err := controller.GetEventDetail(eventId)
	if err != nil {
		fmt.Print("err")
	}

	fmt.Println(detail)
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(detail)
	json.Unmarshal(inrec, &inInterface)
	fmt.Println(inInterface)

	respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusOK, []map[string]interface{}{}, inInterface))
}
