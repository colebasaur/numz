# Binary name
BINARY_NAME=numz

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Build directory
BUILD_DIR=bin

.PHONY: all build clean test coverage run help install dev fmt lint

# Default target
all: test build

## build: Build the binary
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

## clean: Clean build files
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## coverage: Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -cover -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

## run: Run the application
run: build
	@$(BUILD_DIR)/$(BINARY_NAME)

## install: Install the binary to GOPATH/bin
install:
	@echo "Installing..."
	$(GOCMD) install

## dev: Run without building binary
dev:
	$(GOCMD) run main.go

## fmt: Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

## pre-lint: Install golangci-lint
pre-lint:
	@echo "Installing golangci-lint..."
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest

## lint: Run linter (installs golangci-lint if not present)
lint:
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@echo "Running linter..."
	golangci-lint run ./...

## tidy: Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
