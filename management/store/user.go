package store

import (
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
