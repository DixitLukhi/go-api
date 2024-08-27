package main

import (
	"management/api"
	"management/controller"

	"github.com/gin-gonic/gin"
)

// @title Management
// @version 1.0
// @description API for managing School opertions
// @host loclhost:8000
// @BasePath /
// @schemes http https
// @securtyDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token

func main() {
	api := api.ApiRoutes{}
	controller := controller.Server{}

	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":4000")
	// server := controller.Server{}
	// server.NewServer(store.Postgress{})
	// fmt.Println("main server", api)
}
