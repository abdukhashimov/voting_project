package psql

import (
	"context"

	"github.com/abdukhashimov/go_api/internal/pkg/logger"
	"github.com/jackc/pgx/v4"
)

type SQLStore struct {
	DB *pgx.Conn
}

func NewStore(ctx context.Context, psqlUri string) *SQLStore {
	logger.Log.Info("connecting to psql...")
	dbConn, err := pgx.Connect(ctx, psqlUri)
	if err != nil {
		logger.Log.Fatal("failed to connecto to psql", err)
	}

	logger.Log.Info("psql connected")
	return &SQLStore{
		DB: dbConn,
	}
}
