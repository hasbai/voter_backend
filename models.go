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

func (session *Session) AfterFind(tx *gorm.DB) (err error) {
	if session.Motions == nil {
		session.Motions = []Motion{}
	}
	return
}
func (session *Session) AfterCreate(tx *gorm.DB) (err error) {
	return session.AfterFind(tx)
}

type Motion struct {
	BaseModel
	Name        string   `binding:"required" json:"name"`
	Description string   `json:"description"`
	SessionID   int      `json:"sessionID"`
	Status      int8     `json:"status"`
	UserID      int      `json:"userID"`
	For         intArray `json:"for"     `
	Against     intArray `json:"against" `
	Abstain     intArray `json:"abstain"`
}

func (motion *Motion) AfterFind(tx *gorm.DB) (err error) {
	if motion.For == nil {
		motion.For = []int{}
	}
	if motion.Against == nil {
		motion.Against = []int{}
	}
	if motion.Abstain == nil {
		motion.Abstain = []int{}
	}
	return
}
func (motion *Motion) AfterCreate(tx *gorm.DB) (err error) {
	return motion.AfterFind(tx)
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

//func (p intArray) Value() (driver.Value, error) {
//	s := fmt.Sprint(p)
//	return s[1 : len(s)-1], nil
//}

//func (p *intArray) Scan(value interface{}) error {
//	data := value.(string)
//	array := *p
//	if len(data) == 0 {
//		array = []int{}
//		return nil
//	}
//	stringArray := strings.Split(data, " ")
//	array = make([]int, len(stringArray))
//	var err error
//	for i := 0; i < len(stringArray); i++ {
//		array[i], err = strconv.Atoi(stringArray[i])
//	}
//	return err
//}

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
