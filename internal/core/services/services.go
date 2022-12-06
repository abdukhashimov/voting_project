package services

import (
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/core/repository"
)

type Services struct {
	Users    ports.UsersService
	Boards   ports.BoardsService
	Regions  ports.RegionsService
	Statuses ports.StatusService

	Storage repository.Store
}

func NewServices(repos repository.Store) *Services {
	userService := NewUserService(repos)
	boardService := NewBoardService(repos)
	statusService := NewStatusService(repos)
	regionsService := NewRegionService(repos)

	return &Services{
		Storage:  repos,
		Users:    userService,
		Boards:   boardService,
		Statuses: statusService,
		Regions:  regionsService,
	}
}
