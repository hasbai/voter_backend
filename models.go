package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Name    string `json:"name" gorm:"unique"`
	Motions []Motion
}

type Motion struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Records     []Record
	SessionID   int
}

type Record struct {
	gorm.Model
	Vote     int8 `json:"vote"`
	User     User
	UserID   int
	Motion   Motion
	MotionID int
}

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("voter.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Record{}, &Motion{}, &Session{})
	if err != nil {
		panic(err)
	}
}
