export
DATABASE_URL := postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
ENDPOINT := http://localhost:8080
API_KEY := local

.PHONY: tool
tool:
	@aqua i

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: modules
modules:
	@go list -u -m all

.PHONY: renovate
renovate:
	@go get -u -t ./...

.PHONY: compile
compile:
	@go build -v ./... && go clean

.PHONY: test
test:
	@go test ./internal/...

.PHONY: e2e
e2e:
	@go test ./e2e/... -tags e2e

.PHONY: gen
gen:
	@go generate ./... && go mod tidy

.PHONY: doc
doc:
	@rm -rf doc && tbls doc $(DATABASE_URL) doc/databases

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go
