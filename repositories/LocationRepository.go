package repositories

import (
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/aswinda/loket-backend-test/interfaces"
	"github.com/aswinda/loket-backend-test/models"
)

type LocationRepositoryWithCircuitBreaker struct {
	LocationRepository interfaces.ILocationRepository
}

type LocationRepository struct {
	interfaces.IDbHandler
}

func (repository *LocationRepositoryWithCircuitBreaker) StoreLocation(body models.LocationModel) (models.LocationModel, error) {
	output := make(chan models.LocationModel, 1)
	hystrix.ConfigureCommand("store_location", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("store_location", func() error {
		location, _ := repository.LocationRepository.StoreLocation(body)

		output <- location
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.LocationModel{}, err
	}
}

func (repository *LocationRepository) StoreLocation(body models.LocationModel) (models.LocationModel, error) {

	queryString := fmt.Sprintf("insert into locations ( name, status) values ('%s', '%d')", body.Name, 1)
	id, err := repository.Execute(queryString)

	if err != nil {
		return models.LocationModel{}, err
	}

	queryString = fmt.Sprintf("SELECT * FROM locations WHERE id = '%d'", id)
	row, err := repository.Query(queryString)

	if err != nil {
		return models.LocationModel{}, err
	}

	var location models.LocationModel

	row.Next()
	row.Scan(
		&location.Id,
		&location.Name,
		&location.Status,
		&location.Created_at,
		&location.Updated_at)

	return location, nil
}
