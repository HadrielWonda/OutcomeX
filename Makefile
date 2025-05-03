# Makefile

# Configuration
BIN_DIR := bin
COVERAGE_DIR := coverage
PKG := github.com/hadrielwonda/outcomex
VERSION := $(shell git describe --tags --always --dirty)
GO_LDFLAGS := -ldflags "-X main.version=$(VERSION)"
GO_TEST_FLAGS := -v -race -timeout 2m
GO_COVER_FLAGS := -coverprofile=$(COVERAGE_DIR)/coverage.out -coverpkg=./...

# Tools
GOLANGCI_LINT_VERSION := v1.55.2

.PHONY: all
all: build test lint ## Run build, test and lint (default)

# Build targets
.PHONY: build
build: ## Build all binaries
	@mkdir -p $(BIN_DIR)
	go build $(GO_LDFLAGS) -o $(BIN_DIR)/ ./...

.PHONY: install
install: ## Install dependencies
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

# Test targets
.PHONY: test
test: ## Run all tests
	@mkdir -p $(COVERAGE_DIR)
	go test $(GO_TEST_FLAGS) $(GO_COVER_FLAGS) ./...

.PHONY: test-html
test-html: test ## Generate HTML coverage report
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html

.PHONY: benchmark
benchmark: ## Run benchmarks
	go test -benchmem -run=^$$ -bench . ./internal/outcomex

# Linting
.PHONY: lint
lint: ## Run linters
	golangci-lint run ./... --timeout 5m

.PHONY: lint-fix
lint-fix: ## Fix linting issues
	golangci-lint run ./... --fix

# Code formatting
.PHONY: fmt
fmt: ## Format source code
	go fmt ./...

.PHONY: tidy
tidy: ## Tidy go.mod
	go mod tidy

# Examples
.PHONY: examples
examples: ## Build examples
	@cd examples/basic-usage && go build -o ../../$(BIN_DIR)/basic-usage

.PHONY: run-examples
run-examples: examples ## Run examples
	@$(BIN_DIR)/basic-usage

# Documentation
.PHONY: docs
docs: ## Generate documentation
	@mkdir -p docs
	godoc -http=:6060

.PHONY: preview-docs
preview-docs: ## Preview documentation
	@xdg-open http://localhost:6060/pkg/$(PKG) || open http://localhost:6060/pkg/$(PKG)

# Cleanup
.PHONY: clean
clean: ## Clean build artifacts
	@rm -rf $(BIN_DIR) $(COVERAGE_DIR)
	@go clean -testcache

# Help
.PHONY: help
help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)