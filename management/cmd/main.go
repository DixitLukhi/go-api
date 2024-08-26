package main

import (
	"management/api"
	"management/controller"

	"github.com/gin-gonic/gin"
)

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
