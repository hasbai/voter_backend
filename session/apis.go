package session

import (
	"github.com/gin-gonic/gin"
	"voter_backend/db"
	"voter_backend/utils"
)

type AddSessionModel struct {
	Name string `json:"name,omitempty"`
}

// AddSession
// @Summary Add A Session
// @Tags Session
// @Accept application/json
// @Produce application/json
// @Router /sessions [put]
// @Param json body AddSessionModel true "json"
// @Success 201 {object} Session
// @Success 200 {object} Session
func AddSession(c *gin.Context) {
	var session Session
	if err := utils.ValidateJSON(c, &session); err != nil {
		return
	}
	result := db.DB.Where(&session).Preload("Motions").FirstOrCreate(&session)
	var code int
	if result.RowsAffected > 0 {
		code = 201
	} else {
		code = 200
	}
	c.JSON(code, session)
}

// ListSessions
// @Summary List Sessions
// @Tags Session
// @Produce application/json
// @Router /sessions [get]
// @Success 200 {array} SimpleSession
func ListSessions(c *gin.Context) {
	var sessions []Session
	db.DB.Find(&sessions)
	c.JSON(200, sessions)
}

// GetSession
// @Summary Get Session
// @Tags Session
// @Produce application/json
// @Router /sessions/{id} [get]
// @Param id path int true "id"
// @Success 200 {object} Session
// @Failure 404 {object} utils.MessageModel
func GetSession(c *gin.Context) {
	var id utils.IDUri
	if err := utils.ValidateUri(c, &id); err != nil {
		return
	}
	var session Session
	if err := utils.Detect404(c, db.DB.Preload("Motions").First(&session, id.A)); err != nil {
		return
	}
	c.JSON(200, session)
}

// GetTheLatestSession
// @Summary Get The Last Session
// @Tags Session
// @Produce application/json
// @Router /session [get]
// @Success 200 {object} Session
// @Failure 404 {object} utils.MessageModel
func GetTheLatestSession(c *gin.Context) {
	var session Session
	if err := utils.Detect404(c, db.DB.Preload("Motions").Last(&session)); err != nil {
		return
	}
	c.JSON(200, session)
}
