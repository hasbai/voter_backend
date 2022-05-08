package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

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
	db.Preload(clause.Associations).Create(&motion)
	c.JSON(201, motion)
}

// getLastMotion
// @Summary Get The Last Motion
// @Tags Motion
// @Produce application/json
// @Router /motion [get]
// @Success 200 {object} Motion
func getLastMotion(c *gin.Context) {
	var motion Motion
	if err := detect404(c, db.Preload(clause.Associations).Last(&motion)); err != nil {
		return
	}
	c.JSON(200, motion)
}

// getMotion
// @Summary Get A Motion
// @Tags Motion
// @Produce application/json
// @Router /motions/{id} [get]
// @Param id path int true "id"
// @Success 200 {object} Motion
func getMotion(c *gin.Context) {
	var motion Motion
	var id IDUri
	if err := validateUri(c, &id); err != nil {
		return
	}
	if err := detect404(c, db.Preload(clause.Associations).Last(&motion, id.A)); err != nil {
		return
	}
	c.JSON(200, motion)
}
