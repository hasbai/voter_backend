package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "voter_backend/docs"
)

// index
// @Produce application/json
// @Success 200 {object} MessageModel
// @Router / [get]
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

// @title Voter
// @version 0.1.0
// @description voter backend

// @contact.name Maintainer Shi Yue
// @contact.email jsclndnz@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /
func main() {
	initDB()
	router := gin.Default()
	router = registerRouter(router)
	if err := router.Run("localhost:8000"); err != nil {
		panic(err)
	}
}

func registerRouter(router *gin.Engine) *gin.Engine {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/docs/index.html")
	})

	router.GET("/", index)
	router.GET("/users", listUsers)
	router.PUT("/users/:name", addUser)

	router.GET("/session", getLastSession)
	router.GET("/sessions", listSessions)
	router.GET("/sessions/:id", getSession)
	router.PUT("/sessions", addSession)

	router.GET("/motion", getLastMotion)
	router.GET("/motions/:id", getMotion)
	router.POST("/motions", addMotion)

	return router
}
