package ports

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
)

type RegionsService interface {
	Create(ctx context.Context, payload *domain.RegionCreate) (int32, error)
	Update(ctx context.Context, payload *domain.Region) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.RegionAllResp, error)
	Get(ctx context.Context, id int32) (*domain.Region, error)
	Delete(ctx context.Context, id int32) error
}

type DistrictsService interface {
	Create(ctx context.Context, payload *domain.DistrictCreate) (int32, error)
	Update(ctx context.Context, payload *domain.District) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.DistrictAllResp, error)
	Get(ctx context.Context, id int32) (*domain.District, error)
	Delete(ctx context.Context, id int32) error
}

type QuartersService interface {
	Create(ctx context.Context, payload *domain.QuarterCreate) (int32, error)
	Update(ctx context.Context, payload *domain.Quarter) error
	GetAll(ctx context.Context, payload *domain.GetAllParams) (domain.QuarterAllResp, error)
	Get(ctx context.Context, id int32) (*domain.Quarter, error)
	Delete(ctx context.Context, id int32) error
}
