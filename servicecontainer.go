package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/aswinda/loket-backend-test/controllers"
	"github.com/aswinda/loket-backend-test/infrastructures"
	"github.com/aswinda/loket-backend-test/repositories"
	"github.com/aswinda/loket-backend-test/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type IServiceContainer interface {
	InjectEventController() controllers.EventController
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func getConnectionStringMysql() string {
	mysqlUsername := viper.GetString(`database.user`)
	mysqlPassword := viper.GetString(`database.pass`)
	mysqlDefaultDb := viper.GetString(`database.name`)
	mysqlHost := viper.GetString(`database.host`)
	mysqlPort := viper.GetString(`database.port`)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDefaultDb)

	return connectionString
}

func (k *kernel) InjectEventController() controllers.EventController {

	mysqlConn, err := sql.Open("mysql", getConnectionStringMysql())
	if err != nil {
		log.Fatal(err)
	}
	err = mysqlConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = mysqlConn

	eventRepository := &repositories.EventRepository{mysqlHandler}
	circuit := &repositories.EventRepositoryWithCircuitBreaker{eventRepository}
	eventService := &services.EventService{circuit}
	eventController := controllers.EventController{eventService}

	return eventController
}

func (k *kernel) InjectLocationController() controllers.EventController {

	mysqlConn, err := sql.Open("mysql", getConnectionStringMysql())
	if err != nil {
		log.Fatal(err)
	}
	err = mysqlConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = mysqlConn

	locationRepository := &repositories.LocationRepository{mysqlHandler}
	circuit := &repositories.LocationRepositoryWithCircuitBreaker{locationRepository}
	locationService := &services.LocationService{circuit}
	locationController := controllers.LocationController{locationService}

	return locationController
}

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
