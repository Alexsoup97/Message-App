package main

import (
	"github.com/Alexsoup97/message-app/db"
	"github.com/Alexsoup97/message-app/internal/routes"
	"github.com/Alexsoup97/message-app/internal/service"
	"github.com/go-chi/chi/v5"
)

type IService interface{}

type router interface{}

type EndpointContainer struct {
	service IService
	router  chi.Router
	prefix  string
}

func setupEndpoints(db *db.Storage, router chi.Router) {

	endpointContainers := make([]EndpointContainer, 0)

	//User Service
	service := service.CreateUserService(db)
	route := routes.CreateUserRouter(service)
	endpointContainers = append(endpointContainers, EndpointContainer{
		service: service,
		router:  route,
		prefix:  "/user",
	})

	route = routes.CreateMessageRouter(db)
	endpointContainers = append(endpointContainers, EndpointContainer{
		service: service,
		router:  route,
		prefix:  "/messages",
	})

	//Mounting routers
	for _, endpoint := range endpointContainers {
		router.Mount(endpoint.prefix, endpoint.router)
	}
}
