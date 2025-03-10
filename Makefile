GOOSE_DBSTRING ?= "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

# name app
APP_NAME := server

dev:
	go run ./cmd/${APP_NAME}/

run:
	doocker compose up -d && go run ./cmd/${APP_NAME}/
up:
	docker-compose up -d
down:
	docker-compose down

docker_up:
	docker-compose -f environment/docker-compose-dev.yml up

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
# create new a migration
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse run up down swag

.PHONY: air