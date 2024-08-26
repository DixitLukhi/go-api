package controller

import (
	"management/model"
	"management/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server Server) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "Creating new user", nil)

	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "Error while creating new user", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgresDB.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "Error while inserting new user", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
