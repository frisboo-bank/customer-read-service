GOPATH:=$(shell go env GOPATH)

MODULE  := frisboo-bank/customers-service

# Project info
BUILD   := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')
NAME    := $(notdir $(MODULE))
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo v0.0.0)
GOOS    ?= $(shell go env GOOS)
GOARCH  ?= $(shell go env GOARCH)
MARCH   := $(GOOS)-$(GOARCH)

# Local DB info
DB_NAME = customers-service
DB_HOST = localhost
DB_USER = postgres
DB_PASS = postgres
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

# Tool versions
GCI_VERSION := latest
GO_VULN_CHECK_VERSION := latest
GOFUMPT_VERSION := latest
GOLANGCI_VERSION := latest
GOLINES_VERSION := latest
MOCKERY_VERSION  := latest
REVIVE_VERSION := latest
STATIC_CHECK_VERSION := latest

# Build Flags
LDFLAGS := \
	-X '$(MODULE)/internal/version.Name=$(NAME)' \
	-X '$(MODULE)/internal/version.Version=$(VERSION)' \
	-X '$(MODULE)/internal/version.Build=$(BUILD)' \
	-X '$(MODULE)/internal/version.Commit=$(COMMIT)'

# Go Parameters
GO := go
PKG := ./...
BOOTSTRAP := ./cmd/app/
BIN_DIR := bin
COVERAGE := /tmp/$(NAME)-coverage.out

#
# default target
#
.PHONY: all
all: help

#
# QUALITY
#
## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v
	go run github.com/segmentio/golines@$(GOLINES_VERSION) -m 120 -w --ignore-generated .
	go run github.com/daixiang0/gci@$(GCI_VERSION) write --skip-generated -s standard -s "prefix($(MODULE))" -s default -s blank -s dot --custom-order  .
	go run mvdan.cc/gofumpt@$(GOFUMPT_VERSION) -l -w .

## audit: Run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@$(STATIC_CHECK_VERSION) -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@$(GO_VULN_CHECK_VERSION) ./...
	go test -race -buildvcs -vet=off ./...

## lint: Run linters
.PHONY: lint
lint:
	go run github.com/mgechev/revive@$(REVIVE_VERSION) -config revive-config.toml -formatter friendly ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION) run ./...

#
# DEVELOPMENT
#
## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/coverage: run all test and display coverage
.PHONY: test/coverage
test/coverage:
	go test -v -race -buildvcs -coverprofile=$(COVERAGE)  ./...
	go tool cover -html=$(COVERAGE)

## build: build the application
.PHONY: build
build:
	@echo "Buiding $(NAME) $(VERSION) ($(COMMIT)) for $(MARCH)"
	@mkdir -p $(BIN_DIR)
	go build -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(NAME) $(BOOTSTRAP)

## run: run the application
.PHONY: run
run:
	go run $(BOOTSTRAP)

## watch: run the application in watch mode
.PHONY: watch
watch:
	@go run github.com/air-verse/air@latest \
		--root "." \
		--build.cmd "make build" \
	  --build.bin "" \
	  --build.full_bin "go run $(BOOTSTRAP)main.go" \
	  --build.delay "100" \
		--build.include_dir "../pkg/" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## db/create-container: create the db container
.PHONY: db/create-container
db/create-container:
	docker run --name $(NAME)-db -p $(DB_PORT)\:$(DB_PORT) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASS) -d postgres:alpine

## db/create: create the test database
.PHONY: db/create
db/create:
	docker exec -it $(NAME)-db createdb -U $(DB_USER) -O $(DB_USER) $(DB_NAME)

## db/drop: drop the test database
.PHONY: db/drop
db/drop:
	docker exec -it $(NAME)-db dropdb -U $(DB_USER) $(DB_NAME)

#
# PRODUCTION
#
## build/release: build optimized production application
.PHONY: build/release
build/release:
	make build GOOS=$(GOOS) GOARCH=$(GOARCH) -ldflags="$(LDFLAGS) -w -s"

## optimize: compress application binary
.PHONY: optimize
optimize:
	upx -9 -k $(BIN_DIR)/$(NAME) || echo "UPX not installed. Skip..."

## build/docker: build the docker image
.PHONY: build/docker
build/docker:
	docker build -D -t $(NAME):$(VERSION) .

#
# OPERATIONS
#
#.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit lint test no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: tidy audit lint test no-dirty build/release compress

#
# HELPERS
#
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort

.PHONY: version
version: ## Display build version info
	@echo "Application: $(NAME)"
	@echo "Version:     $(VERSION)"
	@echo "Commit:      $(COMMIT)"
	@echo "Build Time:  $(BUILD)"
	@echo "OS-Arch:     $(MARCH)"

.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code
