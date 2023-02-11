.PHONY: help
help: ## Display this help screen. This Makefile is prepared for static analysis of yaml files in the repository.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: aqua
<<<<<<< HEAD
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
=======
aqua: ## Put the path in your environment variables. ex) export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
>>>>>>> main
	@go run github.com/aquaproj/aqua-installer@latest --aqua-version v1.32.3

.PHONY: tool
tool: ## Install tool.
	@aqua i

.PHONY: fmt
fmt: ## Format yaml file.
	@yamlfmt

.PHONY: lint
lint: ## Lint yaml file.
	@yamlfmt -lint
