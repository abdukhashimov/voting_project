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
	_ ports.StatusService = (*StatusService)(nil)
)

type StatusService struct {
	repo repository.Store
}

func NewStatusService(repo repository.Store) *StatusService {
	return &StatusService{
		repo: repo,
	}
}

func (b *StatusService) Create(ctx context.Context, payload *domain.StatusCreate) (string, error) {
	var dbPayload sqlc.CreateBoardParams

	boardID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, b.repo.CreateBoard)
	if err != nil {
		return "", nil
	}

	return boardID, nil
}

func (b *StatusService) Update(ctx context.Context, payload *domain.Status) error {
	var dbPayload sqlc.UpdateBoardParams
	return processor.Execute(ctx, payload, dbPayload, b.repo.UpdateBoard)
}

func (b *StatusService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.StatusAllResp, error) {
	var (
		resp        domain.StatusAllResp
		filter      sqlc.GetAllBoardsParams
		filterCount string
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, b.repo.GetAllBoardsCount)
	if err != nil {
		return resp, err
	}

	objects, err := processor.ExecuteManyWithResp(ctx, params, filter, b.repo.GetAllBoards)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, objects, &resp.Statuss)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (b *StatusService) Get(ctx context.Context, id string) (*domain.Status, error) {
	var (
		resp domain.Status
	)

	board, err := processor.ExecuteWithResp(ctx, id, id, b.repo.GetSingleUserByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, board, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (b *StatusService) Delete(ctx context.Context, id string) error {
	return b.repo.SoftDeleteBoard(ctx, sqlc.SoftDeleteBoardParams{
		Status: 3,
		ID:     id,
	})
}
