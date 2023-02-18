.PHONY: help
help: ## Display this help screen. This Makefile is prepared to automatically generate code from the api definition.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: tool
tool: ## Install tool.
	@aqua i && npm install

.PHONY: lint
lint: ## Lint file.
	@npx spectral lint openapi.yaml

.PHONY: backend
backend: ## Generate code for backend.
	@oapi-codegen -generate types -package openapi openapi.yaml > ../backend/pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi openapi.yaml > ../backend/pkg/openapi/server.gen.go
	@oapi-codegen -generate client -package openapi openapi.yaml > ../backend/pkg/openapi/client.gen.go
	@(cd ../backend && go mod tidy)

.PHONY: frontend
frontend: ## Generate code for frontend.
	@docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g typescript-fetch -o /local/openapi
	@rm -r ../frontend/src/openapi && mv openapi ../frontend/src/openapi

.PHONY: gen
gen: ## Generate code.
	@make backend
	@make frontend

.PHONY: doc
doc: ## Generate documentation.
	@rm -rf doc
	@mkdir -p doc
	@npx widdershins --omitHeader --code true openapi.yaml doc/openapi.md
