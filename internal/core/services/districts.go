package services

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/core/domain"
	"github.com/abdukhashimov/go_api/internal/core/ports"
	"github.com/abdukhashimov/go_api/internal/core/repository"
	"github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc"
	"github.com/abdukhashimov/go_api/internal/pkg/processor"
	"github.com/abdukhashimov/go_api/internal/pkg/serialize"
)

var (
	_ ports.DistrictsService = (*DistrictService)(nil)
)

type DistrictService struct {
	repo repository.Store
}

func NewDistrictService(repo repository.Store) *DistrictService {
	return &DistrictService{
		repo: repo,
	}
}

func (d *DistrictService) Create(ctx context.Context, payload *domain.DistrictCreate) (int32, error) {
	var dbPayload sqlc.CreateDistrictParams

	districtID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, d.repo.CreateDistrict)
	if err != nil {
		return 0, err
	}

	return districtID, nil
}

func (d *DistrictService) Update(ctx context.Context, payload *domain.District) error {
	var dbPayload sqlc.UpdateDistrictParams
	return processor.Execute(ctx, payload, dbPayload, d.repo.UpdateDistrict)
}

func (d *DistrictService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.DistrictAllResp, error) {
	var (
		resp        domain.DistrictAllResp
		filter      sqlc.GetAllDistrictsParams
		filterCount sqlc.GetAllDistrictsCountParams
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, d.repo.GetAllDistrictsCount)
	if err != nil {
		return resp, err
	}

	objects, err := processor.ExecuteManyWithResp(ctx, params, filter, d.repo.GetAllDistricts)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, objects, &resp.Districts)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (d *DistrictService) Get(ctx context.Context, id int32) (*domain.District, error) {
	var (
		resp domain.District
	)

	district, err := processor.ExecuteWithResp(ctx, id, id, d.repo.GetDistrictByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, district, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (r *DistrictService) Delete(ctx context.Context, id int32) error {
	return r.repo.DeleteDistrictByID(ctx, id)
}
