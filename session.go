package main

import (
	"github.com/gin-gonic/gin"
)

type SessionAdd struct {
	Name string `json:"name,omitempty"`
}

// addSession
// @Summary Add A Session
// @Tags Session
// @Accept application/json
// @Produce application/json
// @Router /sessions [put]
// @Param json body SessionAdd true "json"
// @Success 201 {object} Session
// @Success 200 {object} Session
func addSession(c *gin.Context) {
	var session Session
	if err := validateJSON(c, &session); err != nil {
		return
	}
	result := db.Where(Session{Name: session.Name}).FirstOrCreate(&session)
	var code int
	if result.RowsAffected > 0 {
		code = 201
	} else {
		code = 200
	}
	c.JSON(code, session)
}

// listSessions
// @Summary List Sessions
// @Tags Session
// @Produce application/json
// @Router /sessions [get]
// @Success 200 {array} Session
func listSessions(c *gin.Context) {
	var sessions []Session
	db.Find(&sessions)
	c.JSON(200, sessions)
}

// getSession
// @Summary Get Session
// @Tags Session
// @Produce application/json
// @Router /sessions/{id} [get]
// @Param id path int true "id"
// @Success 200 {object} Session
// @Failure 404 {object} MessageModel
func getSession(c *gin.Context) {
	var id IDUri
	if err := validateUri(c, &id); err != nil {
		return
	}
	var session Session
	if id.A <= 0 {
		if err := detect404(c, db.Last(&session)); err != nil {
			return
		}
	} else {
		if err := detect404(c, db.First(&session, id.A)); err != nil {
			return
		}
	}
	c.JSON(200, session)
}

// getLastSession
// @Summary Get The Last Session
// @Tags Session
// @Produce application/json
// @Router /session [get]
// @Success 200 {object} Session
// @Failure 404 {object} MessageModel
func getLastSession(c *gin.Context) {
	var session Session
	if err := detect404(c, db.Last(&session)); err != nil {
		return
	}
	c.JSON(200, session)
}
