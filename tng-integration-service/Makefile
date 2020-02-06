GOPATH := $(PWD)
PATH := "$(PATH):$(GOPATH)/bin:/opt/golang/bin"
SUB_PROJECTS := tng/h5-integration-service tng/user-profile-service
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v /pkg/ | grep -v _test.go)
PKG_LIST := $(shell go list tng/${PKG}/... | grep -v /vendor/)

.PHONY: all prepare test lint build

all: build test lint

prepare: ## Download and install tools
	@go get -v github.com/golangci/golangci-lint/cmd/golangci-lint
	@go get -v github.com/beego/bee
	@go get -v github.com/golang/dep/cmd/dep
	@go get -v github.com/jstemmer/go-junit-report
	@go get -v github.com/vektra/mockery/.../
	@go get -u golang.org/x/lint/golint

prepare-migrate:
	@go get -v github.com/beego/bee
	@go get -v github.com/golang/dep/cmd/dep

lint: ## Lint the files
	@./scripts/golangci-lint.sh $(SUB_PROJECTS)
	@./scripts/lint.sh $(SUB_PROJECTS)

build:
	@for project in $(SUB_PROJECTS) ; do \
		echo "Building $$project" && cd src/$$project && make build && cd - ; \
	done

test: ## Run unit tests
	@./scripts/unittest.sh $(SUB_PROJECTS)

dep: ## Get the dependencies
	@cd src/tng/ && dep ensure -v

fmt: ## Format all *.go package
	@cd src/ && go fmt $(PKG_LIST) && cd -
