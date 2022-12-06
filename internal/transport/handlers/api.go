package handlers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) SetUpAPI() {
	v1 := s.engine.Group("/v1")
	{
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetAllUsers)
		v1.GET("/users/:id", s.GetSingleUser)
		v1.PUT("/users/:id", s.UpdateUser)
		v1.DELETE("/users/:id", s.DeleteUser)

		v1.POST("/boards", s.CreateBoard)
		v1.GET("/boards/:id", s.GetSingleBoard)
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
