package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "voter_backend/docs"
	"voter_backend/motion"
	"voter_backend/session"
	"voter_backend/user"
	"voter_backend/ws"
)

// index
// @Produce application/json
// @Success 200 {object} utils.MessageModel
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	migrateDB()
	go ws.Manager.Start()

	app := gin.Default()
	registerRouter(app)
	user.RegisterRouter(app)
	motion.RegisterRouter(app)
	session.RegisterRouter(app)
	ws.RegisterRouter(app)

	if err := app.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}

func registerRouter(router *gin.Engine) {
	router.GET("/", index)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/docs/index.html")
	})
}
