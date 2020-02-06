GOPATH := $(PWD)/../../..
PATH := "$(PATH):$(GOPATH)/bin:/opt/golang/bin"

migration: ## Generate new migration
	@read -p "Enter migration name: " migration_name; \
	bee generate migration $${migration_name}

migrate: ## Run database migrations
	@bee migrate --conn="$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)?charset=utf8"

dep: ## Get the dependencies
	@dep ensure -v

test:
	@./scripts/unittest.sh

lint:
	@./scripts/lint.sh
	@./scripts/golangci-lint.sh
	
	