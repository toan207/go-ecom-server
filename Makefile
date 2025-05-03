GOOSE_DBSTRING = "root:123456@tcp(127.0.0.1:3306)/mydb"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER = mysql

APP_NAME = server

dev:
	ENV=dev go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:	
	docker compose down

upg: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downg: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetg: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: run upg downg resetg dev