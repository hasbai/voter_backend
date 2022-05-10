package session

import (
	"gorm.io/gorm"
	"voter_backend/db"
	"voter_backend/motion"
)

type SimpleSession struct {
	db.BaseModel
	Name string `json:"name" gorm:"unique" binding:"required"`
}

type Session struct {
	SimpleSession
	Motions []motion.Motion `json:"motions"`
}

func (session *Session) AfterFind(tx *gorm.DB) (err error) {
	if session.Motions == nil {
		session.Motions = []motion.Motion{}
	}
	return
}

func (session *Session) AfterCreate(tx *gorm.DB) (err error) {
	return session.AfterFind(tx)
}
