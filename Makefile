echo:
	echo "HELLO WORLD"

sqlc:
	sqlc generate

swag_init:
	swag init -g internal/transport/handlers/server.go -o api/openapi

mockgen:
	mockgen -package mockdb -destination internal/core/repository/psql/mock/store.go github.com/abdukhashimov/go_api/internal/core/repository/psql/sqlc Querier

image_build:
	docker build -t open_budget_core .

run_dokcer_image:
	docker run --network=host -e APPLICATION_MODE=dev -e PSQL_URI='postgres://postgres:postgres@localhost:5432/open_budget' open http --port 9090

.PHONY: dev_environment_start
dev_environment_start:
	docker compose -f docker-compose.dev.yml up -d

.PHONY: dev_environment_stop
dev_environment_stop:
	docker compose -f docker-compose.dev.yml down

.PHONY: dev_environment_remove
dev_environment_remove:
	docker compose -f docker-compose.dev.yml down --volumes