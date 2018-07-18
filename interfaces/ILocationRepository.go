package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ILocationRepository interface {
	StoreLocation(body models.LocationModel) (models.LocationModel, error)
}
