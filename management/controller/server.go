package controller

import (
	"fmt"
	"management/model"
	"management/store"
	"management/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	PostgresDB store.StoreOperations
}

func (s *Server) NewServer(pgstore store.Postgress) {
	util.SetLogger()
	util.Logger.Infof("Logger setup done\n")

	s.PostgresDB = &pgstore
	err := s.PostgresDB.NewStore()
	if err != nil {
		util.Logger.Errorf("Error while creating store : %v", err)
		util.Log(model.LogLevelError, model.Controller, model.NewServer, "Error while creating store", err)
	} else {
		util.Logger.Infof("Connected to store\n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected to store", nil)
	}

	fmt.Println("server : ", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}
