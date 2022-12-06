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
	_ ports.QuartersService = (*QuarterService)(nil)
)

type QuarterService struct {
	repo repository.Store
}

func NewQuarterService(repo repository.Store) *QuarterService {
	return &QuarterService{
		repo: repo,
	}
}

func (q *QuarterService) Create(ctx context.Context, payload *domain.QuarterCreate) (int32, error) {
	var dbPayload sqlc.CreateQuarterParams

	quarterID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, q.repo.CreateQuarter)
	if err != nil {
		return 0, err
	}

	return quarterID, nil
}

func (q *QuarterService) Update(ctx context.Context, payload *domain.Quarter) error {
	var dbPayload sqlc.UpdateQuarterParams
	return processor.Execute(ctx, payload, dbPayload, q.repo.UpdateQuarter)
}

func (q *QuarterService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.QuarterAllResp, error) {
	var (
		resp        domain.QuarterAllResp
		filter      sqlc.GetAllQuartersParams
		filterCount sqlc.GetAllQuartersCountParams
	)

	objects, err := processor.ExecuteManyWithResp(ctx, params, filter, q.repo.GetAllQuarters)
	if err != nil {
		return resp, err
	}

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, q.repo.GetAllQuartersCount)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, objects, &resp.Quarters)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (q *QuarterService) Get(ctx context.Context, id int32) (*domain.Quarter, error) {
	var (
		resp domain.Quarter
	)

	quarter, err := processor.ExecuteWithResp(ctx, id, id, q.repo.GetQuarterByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, quarter, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (q *QuarterService) Delete(ctx context.Context, id int32) error {
	return q.repo.DeleteQuearterByID(ctx, id)
}
