package ports

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
)

type InitiativesService interface {
	Create(ctx context.Context, payload *domain.InitiativeCreate) (string, error)
	Update(ctx context.Context, payload *domain.Initiative) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.InitiativeAllResp, error)
	Get(ctx context.Context, id string) (*domain.Initiative, error)
	Delete(ctx context.Context, id string) error
}
