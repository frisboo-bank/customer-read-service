GOPATH:=$(shell go env GOPATH)

MODULE  := frisboo-bank/customer-service

# Project info
BUILD   := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')
NAME    := $(notdir $(MODULE))
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo v0.0.0)
ARCH    := $(shell go env GOARCH)
MARCH   := $(shell go env GOOS)-$(shell go env GOARCH)

# Tool versions
GO_VULN_CHECK_VERSION := latest
GOLANGCI_VERSION := latest
MOCKERY_VERSION  := latest
REVIVE_VERSION := latest
STATIC_CHECK_VERSION := latest

# Build Flags
LDFLAGS := -ldflags "\
	-X '$(MODULE)/internal/version.Name=$(NAME)' \
	-X '$(MODULE)/internal/version.Version=$(VERSION)' \
	-X '$(MODULE)/internal/version.Build=$(BUILD)' \
	-X '$(MODULE)/internal/version.Commit=$(COMMIT)'"

# Go Parameters
GO := go
PKG := ./...
BOOTSTRAP := ./cmd/service
TAGS    := netgo
BIN_DIR := bin

#
# default target
#
.PHONY: all
all: help

#
# QUALITY
#
.PHONY: tidy
tidy: ## tidy modfiles and format .go files
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit: ## Run quality control checks
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

.PHONY: lint
lint: ## Run linters
	go run github.com/mgechev/revive@latest -config revive-config.toml -formatter friendly ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

#
# DEVELOPMENT
#
.PHONY: build
build: ## build the application
	@echo "Buiding $(NAME) $(VERSION) ($(COMMIT)) for $(MARCH)"
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 go build -tags $(TAGS) $(LDFLAGS) -o $(BIN_DIR)/$(NAME) $(BOOTSTRAP)

.PHONY: run
run: ## run the application
	go run $(BOOTSTRAP)

.PHONY: watch
watch: ## run the application in watch mode
	go run github.com/air-verse/air@latest \
		--build.cmd "echo 'Skipping build step in Air...'"
	  --build.bin ""
	  --build.full_bin "go run $(BOOTSTRAP)/main.go"
	  --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

#
# HELPERS
#
.PHONY: help
help:
	@echo "Usage:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

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

#
# .PHONY: run_customer_service
# run_customer_service:
# 	@go run ./cmd/service/main.go
#
# .PHONY: build_customer_service
# build_customer_service:
# 	@go build ./cmd/service/main.go
#
#
# .PHONY: lint
# lint:
#
# .PHONY: format
# format:
# 	golines -m 120 -w --ignore-generated .
# 	# gci write --skip-generated -s standard -s "prefix(github.com/mehdihadeli/go-food-delivery-microservices)" -s default -s blank -s dot --custom-order  .
# 	gofumpt -l -w .
#
# .PHONY: update
# update:
# 	@go get -u
#
# .PHONY: deps-reset
# deps-reset:
# 	git checkout -- go.mod
# 	go mod tidy
#
# .PHONY: deps-upgrade
# deps-upgrade:
# 	go get -u -t -d -v ./...
# 	go mod tidy
#
# .PHONY: deps-cleancache
# deps-cleancache:
# 	go clean -modcache
