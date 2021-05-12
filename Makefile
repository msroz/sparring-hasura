.DEFAULT_GOAL := help

.PHONY: run-backend
run-backend: ## Run hasura-engine and database with docker-compose
	-docker-compose up

.PHONY: run-console
run-console: ## Run hasura console
	cd hasura && hasura console

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
