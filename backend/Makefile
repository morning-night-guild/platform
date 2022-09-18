POSTGRES_VERSION = 14.5
APP_NAME = morning-night-guild-platform

export
POSTGRES_IMAGE := postgres:$(POSTGRES_VERSION)-alpine
POSTGRES_HOST := localhost
POSTGRES_DB := postgres
POSTGRES_USER := postgres
POSTGRES_PASS := postgres
POSTGRES_PORT := 5432
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
API_KEY := test

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: test
test:
	@go test \
	` \
	go list ./... | \
	grep -v github.com/morning-night-guild/platform/adapter/gateway/ent | \
	grep -v github.com/morning-night-guild/platform/cmd/migrate \
	` \
	-count=1

.PHONY: run
run:
	@go run main.go

.PHONY: gen
gen:
	@go generate ./...

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go

.PHONY: db
db:
	@docker run --rm -d \
		-p $(POSTGRES_PORT):5432 \
		-e TZ=UTC \
		-e LANG=ja_JP.UTF-8 \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASS) \
		-e POSTGRES_INITDB_ARGS=--encoding=UTF-8 \
		--name $(APP_NAME)-db \
		$(POSTGRES_IMAGE)

.PHONY: psql
psql:
	@docker exec -it $(APP_NAME)-db psql -U postgres

.PHONY: dbstop
dbstop:
	@docker stop $(APP_NAME)-db

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o main