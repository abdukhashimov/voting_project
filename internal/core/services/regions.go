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
	_ ports.RegionsService = (*RegionService)(nil)
)

type RegionService struct {
	repo repository.Store
}

func NewRegionService(repo repository.Store) *RegionService {
	return &RegionService{
		repo: repo,
	}
}

func (r *RegionService) Create(ctx context.Context, payload *domain.RegionCreate) (int32, error) {
	var dbPayload string

	regionID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, r.repo.CreateRegion)
	if err != nil {
		return 0, err
	}

	return regionID, nil
}

func (r *RegionService) Update(ctx context.Context, payload *domain.Region) error {
	var dbPayload sqlc.UpdateBoardParams
	return processor.Execute(ctx, payload, dbPayload, r.repo.UpdateBoard)
}

func (r *RegionService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.RegionAllResp, error) {
	var (
		resp        domain.RegionAllResp
		filter      sqlc.GetAllRegionsParams
		filterCount string
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, r.repo.GetAllBoardsCount)
	if err != nil {
		return resp, err
	}

	objects, err := processor.ExecuteManyWithResp(ctx, params, filter, r.repo.GetAllRegions)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, objects, &resp.Regions)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (r *RegionService) Get(ctx context.Context, id int32) (*domain.Region, error) {
	var (
		resp domain.Region
	)

	region, err := processor.ExecuteWithResp(ctx, id, id, r.repo.GetRegionByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, region, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (r *RegionService) Delete(ctx context.Context, id int32) error {
	return r.repo.DeleteRegion(ctx, id)
}
