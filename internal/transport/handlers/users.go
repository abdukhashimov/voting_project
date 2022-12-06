package handlers

import (
	"net/http"

	"github.com/abdukhashimov/go_api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// @Router /users [post]
// @Summary API to create users
// @Tags users
// @Accept json
// @Produce json
// @Param payload body domain.UserCreate true "users-body"
// @Success 200 {object} domain.SuccessResp
// @Failure 400 {object} domain.ErrorResp
// @Failure 500 {object} domain.ErrorResp
func (s *Server) CreateUser(c *gin.Context) {
	var (
		payload  domain.UserCreate
		response domain.SuccessResp
	)

	resp, err := ExecuteWithResp(c, &payload, s.services.Users.Create)
	if handleError(c, err) {
		return
	}

	response.ID = resp

	c.JSON(http.StatusOK, response)
}

// @Router /users [get]
// @Summary API to get list of users
// @Tags users
// @Accept json
// @Produce json
// @Param filter query domain.GetAllParams false "filter-model"
// @Success 200 {object} domain.UserAllResp
func (s *Server) GetAllUsers(c *gin.Context) {
	payload := parseQueryParams(c)
	resp, err := QueryWithResp(c, payload, s.services.Users.GetAll)
	if handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /users/{id} [get]
// @Summary API to get a single user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "key-id"
// @Success 200 {object} domain.User
func (s *Server) GetSingleUser(c *gin.Context) {
	user, err := QueryWithSingleResp(c, idKeyParam, s.services.Users.Get)
	if handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Router /users/{id} [put]
// @Summary API to update users
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param payload body domain.UserCreate true "paylaod"
// @Success 200 {object} domain.SuccessResp
func (s *Server) UpdateUser(c *gin.Context) {
	var (
		payload  domain.User
		response domain.SuccessResp
	)

	payload.ID = c.Param(idKeyParam)

	err := Execute(c, &payload, s.services.Users.Update)
	if handleError(c, err) {
		return
	}

	response.ID = payload.ID

	c.JSON(http.StatusOK, response)
}

// @Router /users/{id} [delete]
// @Summary API to delete a single user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "key-id"
// @Success 200 {object} domain.SuccessResp
func (s *Server) DeleteUser(c *gin.Context) {
	err := ExecuteDelete(c, idKeyParam, s.services.Users.Delete)
	if handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResp{})
}
