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

func (server Server) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, "Fetching all users", nil)

	users, err := server.PostgresDB.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "Error while fetching all users", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, users)
}

func (server Server) GetUser(ctx *gin.Context) {

	id := ctx.Param("id")

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "Fetching user by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "Fetching user", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := server.PostgresDB.GetUser(uuid)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "Error while fetching user", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
