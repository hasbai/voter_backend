package session

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.GET("/session", GetTheLatestSession)
	router.GET("/sessions", ListSessions)
	router.GET("/sessions/:id", GetSession)
	router.PUT("/sessions", AddSession)
}
