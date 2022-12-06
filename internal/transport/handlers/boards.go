package handlers

import (
	"net/http"

	"github.com/abdukhashimov/go_api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @Router /boards [post]
// @Summary API to create boards
// @Tags boards
// @Accept json
// @Produce json
// @Param payload body domain.BoardCreate true "payload-body"
// @Success 200 {object} domain.SuccessResp
func (s *Server) CreateBoard(c *gin.Context) {
	var (
		payload  domain.BoardCreate
		response domain.SuccessResp
	)

	resp, err := ExecuteWithResp(c, &payload, s.services.Boards.Create)
	if handleError(c, err) {
		return
	}

	response.ID = resp

	c.JSON(http.StatusOK, response)
}

// @Router /boards/{id} [get]
// @Summary API to get a single board
// @Tags boards
// @Accept json
// @Produce json
// @Param id path string true "key-id"
// @Success 200 {object} domain.Board
func (s *Server) GetSingleBoard(c *gin.Context) {
	resp, err := QueryWithSingleResp(c, idKeyParam, s.services.Boards.Get)
	
	if handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}
