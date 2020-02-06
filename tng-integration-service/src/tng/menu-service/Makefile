GOPATH := $(PWD)/../../..
PATH := "$(PATH):$(GOPATH)/bin:/opt/golang/bin"
PROJECT_NAME := "menu-service"
PKG := "menu-service"
PKG_LIST := $(shell export GOPATH=${PWD}/../../.. && go list tng/${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

BUILD_FLAGS = -ldflags "-X $(PROJECT_NAME)/app.BuildDate=$(BUILD_DATE)\
						-X $(PROJECT_NAME)/app.BuildCommitHash=$(SHORT_SHA)\
						-X $(PROJECT_NAME)/app.VersionName=$(VERSION_NAME)\
						-X $(PROJECT_NAME)/app.VersionCode=$(BUILD_ID)"

.PHONY: all test clean

all: pack

version: ## Show Beego and Go version
	@bee version

getbee: ## Download and install beego
	@go get github.com/astaxie/beego

init:
	@dep init -v

run: ## Run development server
	@bee run

docs: ## Generate swagger document
	@bee generate docs

lint: ## Lint the files
	@golangci-lint run ./...

test: ## Run unit tests
 	#@go test -v ${PKG_LIST}

dep: ## Get the dependencies
	@dep ensure -v

build: ## Build the binary file
	@make getbee
	@bee generate docs
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(BUILD_FLAGS) -o ./$(PROJECT_NAME)

clean: ## Remove previous build and test results
	@rm -f $(PROJECT_NAME)
	@rm -rf coverage
	@rm -rf test-results

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

migrate: ## Run database migrations
	@bee migrate --conn="$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)?charset=utf8"