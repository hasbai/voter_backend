package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SimpleSession struct {
	BaseModel
	Name string `json:"name" gorm:"unique" binding:"required"`
}

type Session struct {
	SimpleSession
	Motions []Motion
}

type SimpleMotion struct {
	BaseModel
	Name        string `binding:"required" json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SessionID   int    `json:"sessionID,omitempty"`
	Status      int8   `json:"status,omitempty"`
}

type Motion struct {
	SimpleMotion
	For     []User `json:"for,omitempty"     gorm:"many2many:motion_for;"`
	Against []User `json:"against,omitempty" gorm:"many2many:motion_against;"`
	Abstain []User `json:"abstain,omitempty" gorm:"many2many:motion_abstain;"`
}

type User struct {
	BaseModel
	Name string `json:"name" gorm:"unique" binding:"required"`
}

var db *gorm.DB

func initDB() {
	var err error
	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(sqlite.Open("data/db.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Motion{}, &Session{})
	if err != nil {
		panic(err)
	}
}
