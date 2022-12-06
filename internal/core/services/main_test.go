package services_test

import (
	"context"
	"os"
	"testing"

	"github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc"
	"github.com/abdukhashimov/go_api/internal/core/services"
	"github.com/abdukhashimov/go_api/mocks"
	"github.com/jackc/pgx/v4"
	"github.com/jaswdr/faker"
)

var (
	fake            faker.Faker
	GLOBAL_PSQL_URL = "postgres://postgres:postgres@localhost:5432/open_budget"
	svs             *services.Services
)

func TestMain(m *testing.M) {
	mocks.MockAppLogger()

	fake = faker.New()

	dbConn, err := pgx.Connect(context.Background(), GLOBAL_PSQL_URL)
	if err != nil {
		panic(err)
	}

	db := sqlc.New(dbConn)
	svs = services.NewServices(db)

	os.Exit(m.Run())
}
