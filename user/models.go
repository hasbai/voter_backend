package user

import (
	"voter_backend/db"
)

type User struct {
	db.BaseModel
	Name  string `json:"name" gorm:"unique" binding:"required"`
	Email string `json:"email"`
}

type AddUser struct {
	Name  string `json:"name" gorm:"unique" binding:"required"`
	Email string `json:"email"`
}