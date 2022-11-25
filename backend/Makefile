export
DATABASE_URL := postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: tool
tool:
	@aqua i

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: package
package:
	@go get -u

.PHONY: test
test:
	@go test ./...

.PHONY: gen
gen:
	@rm -rf app/core/wire_gen.go && go generate ./... && (cd app/core && wire) && go mod tidy

.PHONY: doc
doc:
	@rm -rf doc && tbls doc $(DATABASE_URL) doc/databases

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go
