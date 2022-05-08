package main

import "github.com/gin-gonic/gin"

type MotionAdd struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// addMotion
// @Summary Add A Motion
// @Description Add the motion to the latest session
// @Tags Motion
// @Accept application/json
// @Produce application/json
// @Router /motions [post]
// @Param json body MotionAdd true "json"
// @Success 201 {object} Motion
func addMotion(c *gin.Context) {
	var motion Motion
	if err := validateJSON(c, &motion); err != nil {
		return
	}
	var sessionID int
	db.Raw("select id from sessions order by id desc limit 1").Scan(&sessionID)
	if sessionID == 0 {
		c.JSON(400, gin.H{"message": "No sessions exist."})
		return
	}
	motion.SessionID = sessionID
	db.Create(&motion)
	c.JSON(201, motion)
}
