package controllers

import (
	"fmt"
	"net/http"

	"github.com/aswinda/loket-backend-test/interfaces"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type EventController struct {
	interfaces.IEventService
}

func (controller *EventController) GetEventDetailAction(response http.ResponseWriter, request *http.Request) {
	// eventIdParam := chi.URLParam(request, "event_id")
	// eventId, err := strconv.Atoi(eventIdParam)

	queryValues := request.URL.Query()
	eventId := queryValues.Get("event_id")

	// ctx := request.Context()
	// key := ctx.Value("event_id").(string)

	// response.Write([]byte(fmt.Sprintf("hi %v", key)))
	// validator

	validate = validator.New()
	err := validate.Var(eventId, "required,numeric")

	if err != nil {
		// fmt.Println(err) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		for _, err := range err.(validator.ValidationErrors) {

			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			// fmt.Println(err.StructField())     // by passing alt name to ReportError like below
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

	// detail, err := controller.GetEventDetail(eventId)
	// if err != nil {
	// 	json.NewEncoder(response).Encode("Something went wrong!!")
	// }

	// // standarize response
	// json.NewEncoder(response).Encode(detail)
}
