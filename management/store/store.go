package store

import (
	"fmt"
	"management/model"
	"management/util"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=dixitlukhi password=dixit93281 dbname=manage port=5432 sslmode=disable"

	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "Creating new store", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "Error in creating store", err)
		return err
	} else {
		store.DB = db
	}

	err = db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "Error in running automigrate", err)
		return err
	}
	fmt.Println("db : ", db)
	return nil
}

type StoreOperations interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUser(userID uuid.UUID) (model.User, error)
}
