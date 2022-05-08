package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func validateJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return err
	}
	return nil
}

func validateUri(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return err
	}
	return nil
}

func detect404(c *gin.Context, result *gorm.DB) error {
	if result.Error == nil {
		return nil
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(404, gin.H{"message": result.Error.Error()})
		return result.Error
	}
	panic(result.Error)
}

func findInUsers(users []User, user User) int {
	left := 0
	right := len(users)
	for left < right {
		mid := left + (right-left)>>1
		if users[mid].ID < user.ID {
			left = mid + 1
		} else if users[mid].ID > user.ID {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
