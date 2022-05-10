package user

import (
	"github.com/gin-gonic/gin"
	"voter_backend/db"
	"voter_backend/utils"
)

func GetUser(c *gin.Context) (User, error) {
	username := c.GetHeader("Authorization")
	user := User{}
	err := utils.Detect404(c, db.DB.Where("name = ?", username).First(&user))
	return user, err
}
