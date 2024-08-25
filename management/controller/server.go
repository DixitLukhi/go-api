package controller

import "management/store"

type Server struct {
	PostgresDB store.Postgress
}

func (s *Server) NewServer() {
	Server := Server{}
	Server.PostgresDB.NewStore()
}
