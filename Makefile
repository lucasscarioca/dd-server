include .env
export

build:
	@go build ./cmd/server/... -o tmp/main

start: build
	@./tmp/main

run:
	@go run ./cmd/server/...

test:
	@go test -v ./...

install:
	@go mod tidy

new_migration:
	migrate create -ext sql -dir internal/adapters/repository/migrations/ -seq $(name)

migrate_up:
	migrate -path internal/adapters/repository/migrations/ -database ${DB_URL} -verbose up

migrate_down:
	migrate -path internal/adapters/repository/migrations/ -database ${DB_URL} -verbose down

migrate_fix:
	migrate -path internal/adapters/repository/migrations/ -database ${DB_URL} force $(version)
