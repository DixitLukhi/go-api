package api

import (
	"management/controller"
	"management/store"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})
}
