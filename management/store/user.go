package store

import (
	"management/model"
	"management/util"
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
