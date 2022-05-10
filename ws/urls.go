package ws

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.GET("/ws", WebsocketHandler)
}
