package store

import (
	"fmt"
	"management/model"
	"management/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (store Postgress) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Creating new user", nil)

	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "Error while creating new user", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "New user created", nil)
	return nil
}

func (store Postgress) GetUsers() ([]model.User, error) {

	users := []model.User{}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, "Fetching all users from db", nil)

	if err := store.DB.Find(&users).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUsers, "Error while fetching all users from db", err)
		return users, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, "All users from db", users)
	return users, nil
}

func (store Postgress) GetUser(userID uuid.UUID) (model.User, error) {

	util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "Fetching user from db", nil)

	var user model.User

	if err := store.DB.First(&user, "ïd = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "User record not foud", err)
		} else {
			util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "Error while fetching user from db", err)
		}
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, "User from db", user)
	return user, nil
}

func (store Postgress) SignUp(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.SignUp, "Creating new user with signup api", nil)

	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SignUp, "Error while creating new user  with signup api", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.SignUp, "New user created  with signup api", nil)
	return nil
}

func (store Postgress) SignIn(userSignIn *model.UserSignIn) (*model.User, error) {
	var user model.User
	util.Log(model.LogLevelInfo, model.StorePackage, model.SignIn, "Reading user data from sign in", nil)

	response := store.DB.Where("email = ? AND password = ?", userSignIn.Email, userSignIn.Password).First(&user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.SignIn, "Error while creating new user  with signup api", response.Error)
		return &user, fmt.Errorf("error while fetching user record for id, err = %v", response.Error)
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.SignIn, "User data", user)
	return &user, nil
}
