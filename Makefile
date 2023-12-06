include .env

DOCKER_COMPOSE_FILE ?= docker-compose.dev.yaml
POSTGRES_DB_URI ?= 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable'

service-up: ## Build and run service
	sudo docker-compose -f ${DOCKER_COMPOSE_FILE} up --build

service-run: ## Run service
	sudo docker-compose -f ${DOCKER_COMPOSE_FILE} up

service-down: ## Stop service
	sudo docker-compose -f ${DOCKER_COMPOSE_FILE} down

shell-postgres: ## Enter to database console
	sudo docker-compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d user_service_db

migration-create-tasks: ## Create a DB migration files for user
	migrate create -ext sql -dir db/migrations create_tasks_table

migration-up: ## Run migrations UP
	migrate -path db/migrations -database ${POSTGRES_DB_URI} up

migration-down: ## Rollback migrations
	migrate -path db/migrations -database ${POSTGRES_DB_URI} down

swag-gen:
	swag init -g cmd/main.go --parseDependency --parseInternal

test:
	go test ./... -v -cover -short

.PHONY: service-up service-run service-down shell-postgres migration-create-user migration-up migration-down swag-gen test