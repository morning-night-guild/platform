.PHONY: help
help: ## Display this help screen. This Makefile is prepared to automatically generate code from the api definition.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: tool
tool: ## Install tool.
	@aqua i && npm install

.PHONY: lint
lint: ## Lint proto file.
	@(cd proto && buf lint)

.PHONY: format
format: ## Format proto file.
	@buf format -w

.PHONY: backend
backend: ## Generate code for backend.
	@buf generate --template buf.backend.gen.yaml

.PHONY: frontend
frontend: ## Generate code for frontend.
	@buf generate --template buf.frontend.gen.yaml

.PHONY: gen
gen: ## Generate code.
	@make backend
	@make frontend

.PHONY: doc
doc: ## Generate documentation.
	@rm -rf doc && mkdir -p doc/api && protoc --doc_out=./doc/api --doc_opt=markdown,README.md "$(shell find . -name "*.proto")"
