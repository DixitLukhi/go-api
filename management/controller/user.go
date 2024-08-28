package controller

import (
	"management/model"
	"management/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func (server Server) SignUp(c *gin.Context) {
	var user model.User

	util.Log(model.LogLevelInfo, model.ControllerPackage, model.SignUp, "Unmarshaling user data", nil)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	err := server.PostgresDB.SignUp(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.SignUp, "Error in saving user record", user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to sign up user"})
		return
	}

	token := jwt.NewWithClais(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   user.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(),
	})

	tokenString, err := token.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (server Server) SignIn(c *gin.Context) {
	var user model.UserSignIn
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.SignIn,
			"error while unmarshaling payload", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user Data from payload"})
		return
	}

	userResp, err := server.PostgresDB.SignIn(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.SignIn, "error in getting user data from pgress for emailId", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user Data for given user"})
		return
	}
	if userResp.Email != user.Email || userResp.Password != user.Password {
		util.Log(model.LogLevelInfo, model.ControllerPackage, model.SignIn, "user data not matched , database response", userResp)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user data"})
		return
	}

	// Create a new token
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		model.Email:    user.Email,
		model.Password: user.Password,
		model.UserID:   userResp.ID,
		model.Expire:   time.Now().Add(model.TokenExpiration).Unix(), // Token expiration time
		// Additional data can be added here
	})

	// Sign the newtoken with the secret key
	tokenString, err := newtoken.SignedString(model.SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
