package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.GET("/users", listUsers)
	router.PUT("/users/:name", addUser)
}
