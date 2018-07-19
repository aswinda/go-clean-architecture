package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ILocationService interface {
	CreateLocation(body models.LocationModel) (models.LocationModel, error)
}
