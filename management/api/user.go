package api

import (
	"management/model"
	"management/util"

	"github.com/gin-gonic/gin"
)

func (api ApiRoutes) UserRoutes(routes *gin.Engine) {
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
		group.GET("/all", api.GetUsers)
		group.GET("/:id", api.GetUser)
	}
}

func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "Creating new user", nil)
	api.Server.CreateUser(ctx)
}

func (api ApiRoutes) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "Fetching all users", nil)
	api.Server.CreateUser(ctx)
}

func (api ApiRoutes) GetUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "Fetch user", nil)
	api.Server.CreateUser(ctx)
}
