# =============================================================================
# Includes
# =============================================================================

include .env

#GO_PORT = ${MAKEFILE_PORT}
DATABASE_HOST = ${MAKEFILE_DATABASE_HOST}
DATABASE_USER = ${MAKEFILE_DATABASE_USER}
DATABASE_PASSWORD = ${MAKEFILE_DATABASE_PASSWORD}
DATABASE_DBNAME = ${MAKEFILE_DATABASE_DBNAME}
DATABASE_PORT = ${MAKEFILE_DATABASE_PORT}

# =============================================================================
# Variables
# =============================================================================

GO_BIN_AIR = github.com/air-verse/air
GO_BIN_SWAGGER = github.com/go-swagger/go-swagger/cmd/swagger
GO_BIN_MIGRATE = -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
GO_BINARIES = $(GO_BIN_AIR) $(GO_BIN_SWAGGER) $(GO_BIN_MIGRATE)

PROJ_EXE = bin/api.exe
PROJ_MAIN = cmd/app/main.go

DB_PG_MIGRATE_URL = "postgres://${DATABASE_USER}:$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)?sslmode=disable"
DB_PG_MIGRATE_URL_DB = "postgres://${DATABASE_USER}:$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_DBNAME)?sslmode=disable"
MIGRATE_PATH = db/migrations

DOCKER_CONTAINER_GO = go_app
DOCKER_CONTAINER_POSTGRESQL = go_db

#.EXPORT_ALL_VARIABLES:
export DOCKER_CLI_HINTS=false

# =============================================================================
# General
# =============================================================================

all: build

install:
	go mod tidy
	go install github.com/air-verse/air@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

clean:
	rm bin/api.exe

air_w:
	air --build.cmd "go build -o bin/api.exe cmd/app/main.go" --build.bin ".\bin\api"

# =============================================================================
# Database
# =============================================================================

pg_migrate:
	migrate -database $(DB_PG_MIGRATE_URL_DB) -path db/migrations up

db_pg_create:
	psql $(DB_PG_MIGRATE_URL) -c "create database $(DATABASE_DBNAME);"

db_pg_remove:
	psql $(DB_PG_MIGRATE_URL) -c "drop database $(DATABASE_DBNAME);"

# =============================================================================
# GO
# =============================================================================

run:
	go run cmd/app/main.go

build:
	go build cmd/app/main.go

build_w:
	go build -o ./bin/api.exe .\cmd\app\main.go

tidy:
	go mod tidy

test:
	go test --short -coverprofile=docs/cover.out -v ./...

test-v:
	go test -v ./...

test-cover:
	go test --short -cover -coverprofile=cover.out -v ./...

# =============================================================================
# Swagger
# =============================================================================

swagger_doc:
	echo 1

swagger_mock:
	echo 1

# =============================================================================
# Docker
# =============================================================================

docker:
	docker-compose up -d

docker_restart: docker_db_clean docker_clean docker_build_c

docker_i:
	docker-compose up

docker_build_c:
	docker-compose up -d --no-deps --build

docker_build:
	docker build .

docker_db_clean:
	docker exec -it $(DOCKER_CONTAINER_POSTGRESQL) psql $(DB_PG_MIGRATE_URL) -c "DROP SCHEMA public CASCADE;"
	docker exec -it $(DOCKER_CONTAINER_POSTGRESQL) psql $(DB_PG_MIGRATE_URL) -c "CREATE SCHEMA public;"

docker_clean:
	docker container stop $(DOCKER_CONTAINER_GO)
	docker container rm $(DOCKER_CONTAINER_GO)
	docker container stop $(DOCKER_CONTAINER_POSTGRESQL)
	docker container rm $(DOCKER_CONTAINER_POSTGRESQL)

docker_connect_go:
	docker exec -it $(DOCKER_CONTAINER_GO) sh

docker_connect_db:
	docker exec -it $(DOCKER_CONTAINER_POSTGRESQL) bash

