package controller

import (
	"fmt"
	"management/store"
)

type Server struct {
	PostgresDB store.StoreOperations
}

func (s *Server) NewServer(pgstore store.Postgress) {
	s.PostgresDB = &pgstore
	s.PostgresDB.NewStore()
	fmt.Println("server : ", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)
}
