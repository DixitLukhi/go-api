package controller

import (
	"fmt"
	"management/store"
)

type Server struct {
	PostgresDB store.Postgress
}

func (s *Server) NewServer() {
	s.PostgresDB.NewStore()
	fmt.Println("server : ", s)
}
