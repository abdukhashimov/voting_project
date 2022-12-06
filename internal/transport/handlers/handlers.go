package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/domain"
	"github.com/abdukhashimov/go_api/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cast"
)

var (
	internalServiceError = "internal service error"
	objectNotFoundError  = "requested object with given key was not found"
)

var (
	idKeyParam = "id"
)

func makeContextWithTimeout() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(config.TimeoutDuration),
	)

	return ctx, cancel
}

func parseQueryParams(c *gin.Context) *domain.GetAllParams {
	var (
		resp domain.GetAllParams
	)

	resp.Limit = cast.ToInt(c.DefaultQuery("limit", "10"))
	resp.Offset = cast.ToInt(c.DefaultQuery("offset", "0"))
	resp.Search = c.DefaultQuery("search", "")
	resp.DistrictID = cast.ToInt(c.DefaultQuery("district_id", "0"))
	resp.RegionID = cast.ToInt(c.DefaultQuery("region_id", "0"))
	resp.Status = cast.ToInt(c.DefaultQuery("status", "0"))

	return &resp
}

func handleError(c *gin.Context, err error) bool {
	var (
		errorExists = false
	)

	if err != nil {
		logger.Log.Error("http request failed with error: ", err)
		errorExists = true

		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, domain.ErrorResp{
				Message: objectNotFoundError,
				Code:    "404",
			})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResp{
				Message: internalServiceError,
				Code:    "500",
			})
		}

	}

	return errorExists
}
