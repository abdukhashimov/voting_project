package repository

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/config"
	"github.com/abdukhashimov/go_api/internal/core/repository/psql"
	"github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, cfg *config.Config) Store {
	dbConn := psql.NewStore(ctx, cfg.PSQL.URI)

	return sqlc.New(dbConn.DB)
}
