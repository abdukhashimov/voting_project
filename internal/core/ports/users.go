package ports

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
)

type UsersService interface {
	Create(ctx context.Context, payload *domain.UserCreate) (string, error)
	Update(ctx context.Context, payload *domain.User) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.UserAllResp, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}
