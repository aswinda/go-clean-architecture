package services

import (
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type LocationService struct {
	interfaces.ILocationRepository
}

func (service *LocationService) CreateLocation(body models.LocationModel) (models.LocationModel, error) {
	result, err := service.StoreLocation(body)

	return result, err
}
