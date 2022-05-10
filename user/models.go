package user

import (
	"voter_backend/db"
)

type User struct {
	db.BaseModel
	Name string `json:"name" gorm:"unique" binding:"required"`
}
