# Loket Backend Test

To run this project you must install docker and docker-compose

get the project 
```
go get github.com/aswinda/loket-backend-test
```

run docker-compose
```
docker-compose -f docker-compose.local.yml up
```

import the database in folder schema
mysql config on file config.json

open other terminal and install all golang dependency
```
go get github.com/go-chi/chi
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/afex/hystrix-go/hystrix
go get gopkg.in/go-playground/validator.v9
go get gopkg.in/matryer/respond.v1
```

Then run the project 
```
go build && ./loket-backend-test
```

# Development
you can test the api with postman, postman file included in this project 

Trello is used for project management
https://trello.com/b/6HCM5X84/loket-backend-test

This project use Uncle Bob Clean Architecture you can read it here
https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html
Inspired by : 
http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

# Feature
- Dependency Injection
- Circuit Breaker

for the detailed explanation you can read link above


Feel free to contact me at aswinda.pp@gmail.com

