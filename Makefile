.DEFAULT_GOAL := help
MAKEFILE_DIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: run-backend
run-backend: ## Run hasura-engine and database with docker-compose
	docker-compose up

.PHONY: run-firebase
run-firebase: ## Exec `firebase emulators:start`
	cd firebase && GOOGLE_APPLICATION_CREDENTIALS=$(MAKEFILE_DIR)/firebase/sparring-hasura-service-account.json firebase emulators:start

.PHONY: run-app
run-app: ## Run React App in local
	cd app && yarn start

.PHONY: run-local-console
run-local-console: ## Run hasura console in local environment
	cd hasura && hasura console --envfile env/.local.env

.PHONY: run-dev-console
run-dev-console: ## Run hasura console in development environment on Hasura Cloud
	cd hasura && hasura console --envfile env/.dev.env

.PHONY: run-prod-console
run-prod-console: ## Run hasura console in production environment on Hasura Cloud
	cd hasura && hasura console --envfile env/.prod.env

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
