LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ### Run docker-compose
	docker-compose up --build -d postgres rabbitmq && docker-compose logs -f
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

mock-data: ### run mockgen
	go run migrations/crm_mock/crm_mock.go && go run migrations/auth_mock/auth_mock.go
.PHONY: mock

migrate-up: ### migration up
	go run migrations/auth/migrate.go && go run migrations/crm_core/migrate.go

.PHONY: migrate-up

migrate-down: ### migration down
	go run migrations/auth_down/migrate_down.go && go run migrations/crm_core_down/migrate_down.go
.PHONY: migrate-down

start-auth:
	go run cmd/auth/main.go
.PHONY: start-auth

start-crm:
	go run cmd/crm_core/main.go
.PHONY: start-crm

build-dockerfile-auth:
	docker build -f dockerfile-auth -t auth-service .
./PHONY: build-dockerfile-auth

build-dockerfile-crm:
	docker build -f dockerfile-crm -t crm-service .
./PHONY: build-dockerfile-crm

swag-auth:
	cd internal/auth && swag init --parseDependency --parseInternal -g ../../cmd/auth/main.go
./PHONY: swag-auth

swag-crm:
	cd internal/crm_core && swag init --parseDependency --parseInternal -g ../../cmd/crm_core/main.go
./PHONY: swag-crm

gen-proto:
	protoc --go_out ./gw --go_opt paths=source_relative --go-grpc_out ./gw --go-grpc_opt paths=source_relative --grpc-gateway_out=./gw --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true --proto_path=../../ --proto_path=./ auth.proto
./PHONY: gen-proto

linter-go:
	golangci-lint run
./PHONY: linter-go