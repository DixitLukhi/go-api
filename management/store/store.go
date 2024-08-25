package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=dixitlukhi password=dixit93281 dbname=manage port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	} else {
		store.DB = db
	}
	fmt.Println("db : ", db)
	return nil
}

type StoreOperations interface {
	NewStore() error
}
