package main

import (
	"voter_backend/db"
	"voter_backend/motion"
	"voter_backend/session"
	"voter_backend/user"
)

func migrateDB() {
	db.InitDB()
	err := db.DB.AutoMigrate(&user.User{}, &motion.Motion{}, &session.Session{})
	if err != nil {
		panic(err)
	}
}
