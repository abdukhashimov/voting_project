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
	_ ports.InitiativesService = (*InitiativeService)(nil)
)

type InitiativeService struct {
	repo repository.Store
}

func NewInitiativeService(repo repository.Store) *InitiativeService {
	return &InitiativeService{
		repo: repo,
	}
}

func (i *InitiativeService) Create(ctx context.Context, payload *domain.InitiativeCreate) (string, error) {
	var dbPayload sqlc.CreateInitiativeParams

	initiativeID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, i.repo.CreateInitiative)
	if err != nil {
		return "", nil
	}

	return initiativeID, nil
}

func (i *InitiativeService) Update(ctx context.Context, payload *domain.Initiative) error {
	var dbPayload sqlc.UpdateInitiativeByIDParams
	return processor.Execute(ctx, payload, dbPayload, i.repo.UpdateInitiativeByID)
}

func (i *InitiativeService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.InitiativeAllResp, error) {
	var (
		resp        domain.InitiativeAllResp
		filter      sqlc.GetAllInitiativesParams
		filterCount sqlc.GetAllIniativesCountParams
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, i.repo.GetAllIniativesCount)
	if err != nil {
		return resp, err
	}

	initiatives, err := processor.ExecuteManyWithResp(ctx, params, filter, i.repo.GetAllInitiatives)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, initiatives, &resp.Initiatives)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (i *InitiativeService) Get(ctx context.Context, id string) (*domain.Initiative, error) {
	var (
		resp domain.Initiative
	)

	initiative, err := processor.ExecuteWithResp(ctx, id, id, i.repo.GetInitiativeByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, initiative, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (i *InitiativeService) Delete(ctx context.Context, id string) error {
	return i.repo.SoftDeleteInitiative(ctx, sqlc.SoftDeleteInitiativeParams{
		Status: 3,
		ID:     id,
	})
}
