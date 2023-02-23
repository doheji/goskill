package api

import (
	"github.com/gin-gonic/gin"
	db "goskill/db/sqlc"
	"net/http"
)

type CreateSkillRequest struct {
	Learner string `json:"learner" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

// / @Create a skill
// @Tags skills
// @Accept  json
// @Produce  json
// @Param CreateSkill body CreateSkillRequest true "Create Skill"
// @Success 200 {object} db.Skill "ok"
// @Failure 400 {object} ServerError  "We require all fields"
// @Router /skill/createSkill [POST]
func (s *Server) CreateSkillHandler(ctx *gin.Context) {
	var request CreateSkillRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	params := db.CreateSkillParams{
		Learner: request.Learner,
		Name:    request.Name,
	}

	skill, err := s.db.CreateSkill(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, skill)
	return
}
