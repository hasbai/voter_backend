package main

import (
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	name := c.Param("name")
	var user User
	db.FirstOrCreate(&user, User{Name: name})
	c.JSON(200, user)
}
