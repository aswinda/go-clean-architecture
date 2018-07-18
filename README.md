# Loket Backend Test

To run this project you must install docker and docker-compose

run docker-compose
```
docker-compose -f docker-compose.local.yml up
```

open other terminal and install all golang dependency
```
go get github.com/go-chi/chi
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/afex/hystrix-go/hystrix
go get gopkg.in/go-playground/validator.v9
go get gopkg.in/matryer/respond.v1
```