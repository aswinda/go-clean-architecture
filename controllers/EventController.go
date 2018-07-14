package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aswinda/notifyme/application/api/interfaces"
	"github.com/go-chi/chi"
)

type EventController struct {
	interfaces.IEventService
}

func (controller *EventController) GetEventDetailAction(response http.ResponseWriter, request *http.Request) {
	eventIdParam := chi.URLParam(request, "event_id")
	eventId, err := strconv.Atoi(eventIdParam)

	// validator

	detail, err := controller.GetEventDetail(eventId)
	if err == nil {
		json.NewEncoder(response).Encode("Something went wrong!!")
	}

	// standarize response
	json.NewEncoder(response).Encode(detail)
}
