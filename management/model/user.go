package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      Name      `gorm:"embedded" binding:"required" json:"name" example:"Dixit"`
	Email     string    `gorm:"not null; unique" binding:"required" json:"email"`
	Password  string    `gorm:"not null" binding:"required" json:"password"`
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
	Pincode string `gorm:"not null; unique" json:"pincode"`
}

type UserSignIn struct {
	Email    string `gorm:"not null; unique" binding:"required" json:"email"`
	Password string `gorm:"not null" binding:"required" json:"password"`
}
