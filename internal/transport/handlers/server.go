package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/abdukhashimov/go_api/api/openapi"
	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/repository"
	"github.com/abdukhashimov/go_api/internal/core/services"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine   *gin.Engine
	services *services.Services
	config   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	var (
		server Server
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(cfg.Project.Timout),
	)
	defer cancel()

	server.config = cfg

	switch cfg.Project.Mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	server.engine = gin.Default()

	respos := repository.New(ctx, cfg)

	server.services = services.NewServices(respos)

	server.SetUpAPI()

	return &server
}

// @title Open budget API
// @description This API contains the source for the Open budget app

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// @BasePath /v1

// Run initializes http server
func (s *Server) Run(port string) {
	var (
		portNumber int
	)

	if port != "" {
		portNum, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}

		portNumber = portNum
	} else {
		portNumber = s.config.Http.HTTP_PORT
	}

	if s.config.Project.SwaggerEnabled {
		openapi.SwaggerInfo.Version = s.config.Project.Version

		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.URL(fmt.Sprintf(
				"%s/%d/swagger/docs.json",
				s.config.Http.HTTP_HOST,
				s.config.Http.HTTP_PORT,
			)),
			ginSwagger.DefaultModelsExpandDepth(-1),
		)
	}

	s.engine.Run(fmt.Sprintf("%s:%d", s.config.Http.HTTP_HOST, portNumber))
}
