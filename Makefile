DEV_VARS = POSTGRES_HOST=127.0.0.1 POSTGRES_PORT=5434 POSTGRES_USER=eShop POSTGRES_PASSWORD=eShop POSTGRES_DBNAME=eShop_db

.PHONY: run
## Run server. Usage: 'make run'
run: ; $(info running server...) @
	@$(DEV_VARS) go run ./cmd/server/main.go

.PHONY: migrations
## Run migrations. Usage: 'make migrations'
migrations: ; $(info running migrations...) @
	@$(DEV_VARS) go run ./cmd/migrations/main.go