package main

import "github.com/gin-gonic/gin"

func registerRouter(router *gin.Engine) *gin.Engine {
	router.GET("/", index)
	router.PUT("/users/:name", createUser)
	return router
}
