package api

import (
	"database/sql"
	"errors"
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
// @Router /createSkill [POST]
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

type GetSkillByIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// / @Get a skill by ID
// @Tags skills
// @Accept  json
// @Produce  json
// @Param id path int64 true "id"
// @Success 200 {object} db.Skill "ok"
// @Failure 400 {object} ServerError  "Something is wrong"
// @Failure 404 {object} ServerError  "Something is wrong"
// @Failure 500 {object} ServerError  "Something is wrong"
// @Router /getskillbyid/{id} [get]
func (s *Server) GetSkillByIDHandler(ctx *gin.Context) {
	var request GetSkillByIDRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	skill, err := s.db.GetSkillByID(ctx, request.ID)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, NewServerError(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, skill)
	return
}
