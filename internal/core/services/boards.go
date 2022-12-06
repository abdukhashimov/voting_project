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
	_ ports.BoardsService = (*BoardService)(nil)
)

type BoardService struct {
	repo repository.Store
}

func NewBoardService(repo repository.Store) *BoardService {
	return &BoardService{
		repo: repo,
	}
}

func (b *BoardService) Create(ctx context.Context, payload *domain.BoardCreate) (string, error) {
	var dbPayload sqlc.CreateBoardParams

	boardID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, b.repo.CreateBoard)
	if err != nil {
		return "", nil
	}

	return boardID, nil
}

func (b *BoardService) Update(ctx context.Context, payload *domain.Board) error {
	var dbPayload sqlc.UpdateBoardParams
	return processor.Execute(ctx, payload, dbPayload, b.repo.UpdateBoard)
}

func (b *BoardService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.BoardAllResp, error) {
	var (
		resp        domain.BoardAllResp
		filter      sqlc.GetAllBoardsParams
		filterCount string
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, b.repo.GetAllBoardsCount)
	if err != nil {
		return resp, err
	}

	boards, err := processor.ExecuteManyWithResp(ctx, params, filter, b.repo.GetAllBoards)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, boards, &resp.Boards)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (b *BoardService) Get(ctx context.Context, id string) (*domain.Board, error) {
	var (
		resp domain.Board
	)

	board, err := processor.ExecuteWithResp(ctx, id, id, b.repo.GetBoardByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, board, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (b *BoardService) Delete(ctx context.Context, id string) error {
	return b.repo.SoftDeleteBoard(ctx, sqlc.SoftDeleteBoardParams{
		Status: 3,
		ID:     id,
	})
}
