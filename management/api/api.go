package api

import (
	"management/controller"
	"management/store"

	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(routes *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	api.UserRoutes(routes)
}
