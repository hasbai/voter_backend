package user

import (
	"github.com/gin-gonic/gin"
	"voter_backend/db"
)

type UpdateUser struct {
	Vote int8 `json:"vote"`
}

// addUser
// @Summary Add A User
// @Description
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param name path string true "username"
// @Success 200 {object} User
// @Router /users/{name} [put]
func addUser(c *gin.Context) {
	name := c.Param("name")
	var user User
	db.DB.FirstOrCreate(&user, User{Name: name})
	c.JSON(200, user)
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
