package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return err
	}
	return nil
}

func ValidateUri(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return err
	}
	return nil
}

func Detect404(c *gin.Context, result *gorm.DB) error {
	if result.Error == nil {
		return nil
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(404, gin.H{"message": result.Error.Error()})
		return result.Error
	}
	panic(result.Error)
}

