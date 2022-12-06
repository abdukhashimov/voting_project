package ports

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
)

type BoardsService interface {
	Create(ctx context.Context, payload *domain.BoardCreate) (string, error)
	Update(ctx context.Context, payload *domain.Board) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.BoardAllResp, error)
	Get(ctx context.Context, id string) (*domain.Board, error)
	Delete(ctx context.Context, id string) error
}
