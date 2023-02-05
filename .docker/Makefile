include .env

.PNONY: init
init:
	@cp .tmp.env .env

.PHONY: up
up:
	@docker compose --project-name ${APP_NAME} up -d

.PHONY: down
down:
	@docker compose --project-name ${APP_NAME} down

.PHONY: psql
psql:
	@docker exec -it ${APP_NAME}-db psql -U postgres

.PHONY: rm
rm:
	@docker image rm ${APP_NAME}-app-core
