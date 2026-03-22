.PHONY: help run test migrate-up migrate-down migrate-version migrate-force migrate-create

help:
	@echo "Available targets:"
	@echo "  make run                  - run the API server"
	@echo "  make test                 - run all tests"
	@echo "  make migrate-up           - apply all pending migrations"
	@echo "  make migrate-down         - rollback one migration step"
	@echo "  make migrate-version      - show current migration version"
	@echo "  make migrate-force VERSION=1"
	@echo "  make migrate-create NAME=add_users_table"

run:
	go run .

test:
	go test ./...

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

migrate-version:
	go run ./cmd/migrate version

migrate-force:
	go run ./cmd/migrate force $(VERSION)

migrate-create:
	go run ./cmd/migrate create $(NAME)
