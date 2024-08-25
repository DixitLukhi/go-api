package main

import (
	"fmt"
	"management/controller"
)

func main() {
	server := controller.Server{}
	server.NewServer()
	fmt.Println("main server", server)
}
