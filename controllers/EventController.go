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

var validate *validator.Validate

type EventController struct {
	interfaces.IEventService
}

func ApiResponse(code int, messages string, data map[string]interface{}) map[string]interface{} {
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
		errs := err.(validator.ValidationErrors)
		str := fmt.Sprintf("%s", errs)
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, str, map[string]interface{}{}))
		return
	}

	detail, err := controller.GetEventDetail(eventId)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, fmt.Sprintf("%s", err), map[string]interface{}{}))
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(detail)
	json.Unmarshal(inrec, &inInterface)

	respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusOK, "", inInterface))
}

func (controller *EventController) CreateEventAction(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var req models.EventModel
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

	result, err := controller.CreateEvent(req)
	if err != nil {
		respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusBadRequest, fmt.Sprintf("%s", err), map[string]interface{}{}))
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(result)
	json.Unmarshal(inrec, &inInterface)

	respond.With(w, r, http.StatusBadRequest, ApiResponse(http.StatusOK, "", inInterface))
}
