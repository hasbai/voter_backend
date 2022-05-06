package main

import "github.com/gin-gonic/gin"

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()
	router = registerRouter(router)
	//goland:noinspection GoUnhandledErrorResult
	router.Run("localhost:8000")
}
