package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      Name      `gorm:"embedded" json:"name"`
	Address   Address   `gorm:"embedded" json:"address"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy string    `json:"deleted_by,omitempty"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Address struct {
	Lane    string `json:"lane"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}
