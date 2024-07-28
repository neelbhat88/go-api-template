PROJNAME := api-template
DB_CONTAINER := ${PROJNAME}-db
DB_CONNECTION_URI := 'postgres://user:password@localhost:5437/mydb?sslmode=disable'
BUILD_DIR := build
APP_NAME := application

.PHONY: build application test

default: build

all: clean run

clean:
	go clean

run: build
	./$(BUILD_DIR)/$(APP_NAME)

mocks:
	go generate ./...

application: cmd/*.go
	go mod tidy -compat=1.21
	go build -o $(BUILD_DIR)/$(APP_NAME) $^

build: mocks application

db-up:
	docker-compose up -d $(DB_CONTAINER)

db-down:
	docker-compose down

create-migration:
	@echo "Please enter a filename..."; \
	read FILE; \
	migrate create -ext sql -dir internal/data/postgres/migrations -seq $$FILE;

test:
	time go test -v ./...

start: docker
	docker-compose up

docker: mocks
	docker build --progress=plain . --tag $(APP_NAME) --build-arg SSH_PRIVATE_KEY="`cat $$SSH_KEY_FILEPATH`"

migrate-up:
	 migrate -path internal/data/postgres/migrations -database $(DB_CONNECTION_URI) up

migrate-down:
	 migrate -path internal/data/postgres/migrations -database $(DB_CONNECTION_URI) down