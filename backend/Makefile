export
DATABASE_URL := postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
ENDPOINT := http://localhost:8080
API_KEY := local

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: tool
tool: ## Install tool.
	@aqua i

.PHONY: fmt
fmt: ## Format code.
	@go fmt ./...

.PHONY: lint
lint: ## Lint code.
	@golangci-lint run --fix

.PHONY: modules
modules: ## List modules with dependencies.
	@go list -u -m all

.PHONY: renovate
renovate: ## Update modules with dependencies.
	@go get -u -t ./...

.PHONY: compile
compile: ## Compile code.
	@go build -v ./... && go clean

.PHONY: test
test: ## Run unit test. If you want to invalidate the cache, please specify an argument like `make test c=c`.
	@$(call _test,${c})

define _test
if [ -z "$1" ]; then \
	go test ./internal/... ; \
else \
	go test ./internal/... -count=1 ; \
fi
endef

.PHONY: e2e
e2e: ## Run e2e test. If you want to invalidate the cache, please specify an argument like `make e2e c=c`.
	@$(call _e2e,${c})

define _e2e
if [ -z "$1" ]; then \
	go test ./e2e/... -tags e2e ; \
else \
	go test ./e2e/... -tags e2e -count=1 ; \
fi
endef

.PHONY: gen
gen: ## Generate code.
	@go generate ./... && go mod tidy

.PHONY: doc
doc: ## Generate documentation.
	@rm -rf doc && tbls doc $(DATABASE_URL) doc/databases
