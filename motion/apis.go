package motion

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"voter_backend/db"
	"voter_backend/user"
	"voter_backend/utils"
	"voter_backend/ws"
)

type AddMotionModel struct {
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
// @Param json body AddMotionModel true "json"
// @Success 201 {object} Motion
// @Security ApiKeyAuth
func addMotion(c *gin.Context) {
	// validate motion
	var motion Motion
	if err := utils.ValidateJSON(c, &motion); err != nil {
		return
	}
	// get user
	u, err := user.GetUser(c)
	if err != nil {
		return
	}
	motion.UserID = u.ID
	// get the last session
	var sessionID int
	db.DB.Raw("select id from sessions order by id desc limit 1").Scan(&sessionID)
	if sessionID == 0 {
		c.JSON(400, gin.H{"message": "No sessions exist."})
		return
	}
	motion.SessionID = sessionID
	// create motion and return
	db.DB.Create(&motion)
	ws.BroadcastObject("motion", &motion)
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
	if err := utils.Detect404(c, db.DB.Last(&motion)); err != nil {
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
	var id utils.IDUri
	if err := utils.ValidateUri(c, &id); err != nil {
		return
	}
	if err := utils.Detect404(c, db.DB.Last(&motion, id.A)); err != nil {
		return
	}
	c.JSON(200, motion)
}

// voteMotion
// @Summary Vote A Motion
// @Tags Motion
// @Produce application/json
// @Router /motions/{id}/{type} [post]
// @Param id path int true "id"
// @Param type path string true "type"
// @Success 200 {object} Motion
// @Security ApiKeyAuth
func voteMotion(c *gin.Context) {
	// get type
	var uri utils.IdNameUri
	err := utils.ValidateUri(c, &uri)
	if err != nil {
		return
	}
	// get user
	u, err := user.GetUser(c)
	if err != nil {
		return
	}
	// get motion
	var motion Motion
	err = utils.Detect404(c, db.DB.First(&motion, uri.A))
	if err != nil {
		return
	}
	// vote
	locAbstain := slices.Index(motion.Abstain, u.ID)
	locFor := slices.Index(motion.For, u.ID)
	locAgainst := slices.Index(motion.Against, u.ID)
	switch uri.B {
	case "for":
		if locFor >= 0 || locAgainst >= 0 {
			break
		}
		if locAbstain >= 0 {
			motion.Abstain = slices.Delete(motion.Abstain, locAbstain, locAbstain+1)
		}
		motion.For = append(motion.For, u.ID)
	case "against":
		if locFor >= 0 || locAgainst >= 0 {
			break
		}
		if locAbstain >= 0 {
			motion.Abstain = slices.Delete(motion.Abstain, locAbstain, locAbstain+1)
		}
		motion.Against = append(motion.Against, u.ID)
	case "abstain":
		if locFor >= 0 || locAgainst >= 0 || locAbstain >= 0 {
			break
		}
		motion.Abstain = append(motion.Abstain, u.ID)
	}
	db.DB.Save(&motion)
	ws.BroadcastObject("motion", &motion)
	c.JSON(200, motion)
}

// resolveMotion
// @Summary Resolve A Motion
// @Tags Motion
// @Produce application/json
// @Router /motions/{id} [put]
// @Param id path int true "id"
// @Success 200 {object} Motion
func resolveMotion(c *gin.Context) {
	// get motion
	var uri utils.IDUri
	err := utils.ValidateUri(c, &uri)
	if err != nil {
		return
	}
	var motion Motion
	err = utils.Detect404(c, db.DB.First(&motion, uri.A))
	if err != nil {
		return
	}
	// resolve motion
	if len(motion.For) > len(motion.Against) {
		motion.Status = 1
	} else {
		motion.Status = -1
	}
	db.DB.Save(&motion)
	ws.BroadcastObject("motion", &motion)
	c.JSON(200, motion)
}
