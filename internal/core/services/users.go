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
	_ ports.UsersService = (*UserService)(nil)
)

type UserService struct {
	repo repository.Store
}

func NewUserService(repo repository.Store) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(ctx context.Context, payload *domain.UserCreate) (string, error) {
	var dbPayload sqlc.CreateUserParams

	userID, err := processor.ExecuteWithResp(ctx, payload, dbPayload, u.repo.CreateUser)
	if err != nil {
		return "", nil
	}

	return userID, nil
}

func (u *UserService) Update(ctx context.Context, payload *domain.User) error {
	var dbPayload sqlc.UpdateUserParams
	return processor.Execute(ctx, payload, dbPayload, u.repo.UpdateUser)
}

func (u *UserService) GetAll(ctx context.Context, params *domain.GetAllParams) (domain.UserAllResp, error) {
	var (
		resp        domain.UserAllResp
		filter      sqlc.GetAllUsersParams
		filterCount sqlc.GetAllUsersCountParams
	)

	count, err := processor.ExecuteManyWithResp(ctx, params, filterCount, u.repo.GetAllUsersCount)
	if err != nil {
		return resp, err
	}

	users, err := processor.ExecuteManyWithResp(ctx, params, filter, u.repo.GetAllUsers)
	if err != nil {
		return resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, users, &resp.Users)
	if err != nil {
		return resp, err
	}

	resp.Count = int(count)

	return resp, nil
}

func (u *UserService) Get(ctx context.Context, id string) (*domain.User, error) {
	var (
		resp domain.User
	)

	user, err := processor.ExecuteWithResp(ctx, id, id, u.repo.GetSingleUserByID)
	if err != nil {
		return &resp, err
	}

	err = serialize.MarshalUnMarshal(ctx, user, &resp)
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (u *UserService) Delete(ctx context.Context, id string) error {
	return u.repo.SoftDeleteUser(ctx, sqlc.SoftDeleteUserParams{
		Status: 3,
		ID:     id,
	})
}
