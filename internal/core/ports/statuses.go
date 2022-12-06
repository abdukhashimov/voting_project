package ports

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
)

type StatusService interface {
	Create(ctx context.Context, payload *domain.StatusCreate) (string, error)
	Update(ctx context.Context, payload *domain.Status) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.StatusAllResp, error)
	Get(ctx context.Context, id string) (*domain.Status, error)
	Delete(ctx context.Context, id string) error
}
