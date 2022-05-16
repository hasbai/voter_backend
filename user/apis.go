package user

import (
	"github.com/gin-gonic/gin"
	"voter_backend/db"
	"voter_backend/utils"
)

// addUser
// @Summary Add A User
// @Description
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param json body AddUser true "json"
// @Success 200 {object} User
// @Success 201 {object} User
// @Router /users [post]
func addUser(c *gin.Context) {
	var user User
	var body AddUser
	if err := utils.ValidateJSON(c, &body); err != nil {
		return
	}
	result := db.DB.FirstOrCreate(&user, User{Name: body.Name})
	if body.Email != "" {
		user.Email = body.Email
		db.DB.Save(&user)
	}
	var code int
	if result.RowsAffected > 0 {
		code = 201
	} else {
		code = 200
	}
	c.JSON(code, user)
}

// listUsers
// @Summary List Users
// @Description
// @Tags User
// @Accept application/json
// @Produce application/json
// @Success 200 {array} User
// @Router /users [get]
func listUsers(c *gin.Context) {
	var users []User
	db.DB.Find(&users)
	c.JSON(200, users)
}
