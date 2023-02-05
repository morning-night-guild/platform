include .env

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PNONY: env
env: ## create .env file
	@cp .env.local .env

.PHONY: up
up: ## create development environment
	@docker compose --project-name ${APP_NAME} up -d

.PHONY: down
down: ## delete development environment
	@docker compose --project-name ${APP_NAME} down

.PHONY: psql
psql: ## connect local postgres
	@docker exec -it ${APP_NAME}-db psql -U postgres
