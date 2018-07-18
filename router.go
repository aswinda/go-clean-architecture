package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	eventController := ServiceContainer().InjectEventController()

	r := chi.NewRouter()
	r.Get("/event/get_info", eventController.GetEventDetailAction)
	r.Post("/event/create", eventController.CreateEventAction)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
