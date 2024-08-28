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
		group.POST("/signup", api.SignUp)
		group.POST("/signin", api.SignIn)
		group.GET("/all", api.GetUsers)
		group.GET("/:id", api.GetUser)
	}
}

// Handler to create user
//
//		@router       /user/create [post]
//	 @summary Create a user
//		@Tags         users
//		@Produce      json
//
// @produce json
//
//	@Param        user body model.User true "User object"
//	@Success      201  {object}   model.User
func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "Creating new user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get all users
//
//		@router       /user/all [get]
//	 @summary Get all users
//		@Tags         users
//		@Produce      json
//		@Param        page query int false "Page number (default: 1)"
//		@Param        limit query int false "Number of results per page (default: 10)"
//		@Success      200  {array}   model.User
func (api ApiRoutes) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "Fetching all users", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get a user by ID
//
//		@router       /user/{id} [get]
//	 @summary Get a user by ID
//		@Tags         users
//		@Produce      json
//		@Param        id path string true "User ID"
//		@Success      200  {object}   model.User
func (api ApiRoutes) GetUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "Fetch user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to SignUp a user
// @router /user/signup [post]
// @summary SignUp a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @Success 200 {string} string "Successful SignUp"
// @failure 400 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api ApiRoutes) SignUp(c *gin.Context) {
	api.Server.SignUp(c)
}

// Handler to signIn a user by email and password
// @router /user/signin [post]
// @summary SighIn user
// @tags users
// @produce json
// @param user body model.UserSignIn true "User object"
// @Success 200 {string} string "Successful SignIn"
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api ApiRoutes) SignIn(c *gin.Context) {
	api.Server.SignIn(c)
}
