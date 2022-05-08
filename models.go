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
	Motions []Motion `json:"motions"`
}

type SimpleMotion struct {
	BaseModel
	Name        string `binding:"required" json:"name"`
	Description string `json:"description"`
	SessionID   int    `json:"sessionID"`
	Status      int8   `json:"status"`
}

type Motion struct {
	SimpleMotion
	For     []User `json:"for"     gorm:"many2many:motion_for;"`
	Against []User `json:"against" gorm:"many2many:motion_against;"`
	Abstain []User `json:"abstain" gorm:"many2many:motion_abstain;"`
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
