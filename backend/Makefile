POSTGRES_VERSION = 14.5

export
POSTGRES_IMAGE := postgres:$(POSTGRES_VERSION)-alpine
POSTGRES_HOST := localhost
POSTGRES_DB := postgres
POSTGRES_USER := postgres
POSTGRES_PASS := postgres
POSTGRES_PORT := 5432
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: test
test:
	@go test ./...

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
		--name article-share-db \
		$(POSTGRES_IMAGE)

.PHONY: psql
psql:
	@docker exec -it article-share-db psql -U postgres

.PHONY: dbstop
dbstop:
	@docker stop article-share-db

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o main