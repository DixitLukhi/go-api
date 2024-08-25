package main

import (
	"fmt"
	"management/api"
	"management/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})

	// server := controller.Server{}
	// server.NewServer(store.Postgress{})
	fmt.Println("main server", api)
}
