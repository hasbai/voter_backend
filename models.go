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

type Session struct {
	BaseModel
	Name    string `json:"name" gorm:"unique" binding:"required"`
	Motions []Motion
}

type Motion struct {
	BaseModel
	Name        string `json:"name" gorm:"unique" binding:"required"`
	Description string `json:"description"`
	Records     []Record
	SessionID   int
}

type Record struct {
	BaseModel
	Vote     int8 `json:"vote" binding:"required"`
	User     User
	UserID   int
	Motion   Motion
	MotionID int
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
	err = db.AutoMigrate(&User{}, &Record{}, &Motion{}, &Session{})
	if err != nil {
		panic(err)
	}
}
