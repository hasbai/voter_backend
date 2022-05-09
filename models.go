package main

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"primarykey" json:"id"`
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

type Motion struct {
	BaseModel
	Name        string   `binding:"required" json:"name"`
	Description string   `json:"description"`
	SessionID   int      `json:"sessionID"`
	Status      int8     `json:"status"`
	UserID      int      `json:"userID"`
	For         intArray `json:"for" `
	Against     intArray `json:"against"`
	Abstain     intArray `json:"abstain" `
}

type User struct {
	BaseModel
	Name string `json:"name" gorm:"unique" binding:"required"`
}

type intArray []int

func (p intArray) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *intArray) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
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
