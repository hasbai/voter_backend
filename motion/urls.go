package motion

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.GET("/motion", getLastMotion)
	router.GET("/motions/:id", getMotion)
	router.POST("/motions", addMotion)
	router.POST("/motions/:id/:name", voteMotion)
	router.PUT("/motions/:id", resolveMotion)
}
