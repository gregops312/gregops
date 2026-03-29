# Variables
BINARY_NAME=gregops
GO_VERSION=1.26
MAIN_PACKAGE=.
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-s -w"
BUILD_FLAGS=-v

.PHONY: all build clean test deps fmt vet install uninstall help

# Default target
all: clean deps fmt vet test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(BUILD_FLAGS) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Build for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)

.PHONY: gitignore
gitignore:
	@echo "Generating .gitignore file..."; \
	if [ -f .gitignore ]; then \
			sed '/^### Generated .gitignore contents, place custom entries above this line ###/,$$d' .gitignore > .gitignore.header; \
	else \
			touch .gitignore.header; \
	fi; \
	curl -sSL https://www.toptal.com/developers/gitignore/api/git,go,macos,vim,visualstudiocode > .gitignore.generated; \
	cat .gitignore.header > .gitignore.tmp; \
	echo "### Generated .gitignore contents, place custom entries above this line ###" >> .gitignore.tmp; \
	cat .gitignore.generated >> .gitignore.tmp; \
	mv .gitignore.tmp .gitignore; \
	rm -f .gitignore.header .gitignore.generated; \
	echo ".gitignore file generated."

# Run tests (use TEST=pattern to run specific tests with -run flag)
test:
	@echo "Running tests..."
ifdef TEST
	$(GOTEST) -v -run "$(TEST)" ./...
else
	$(GOTEST) -v ./...
endif

# Run tests with coverage (use TEST=pattern to run specific tests with -run flag)
test-cover:
	@echo "Running tests with coverage..."
ifdef TEST
	$(GOTEST) -v -run "$(TEST)" -coverprofile=coverage.out ./...
else
	$(GOTEST) -v -coverprofile=coverage.out ./...
endif
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -s -w .

# Vet code
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...

# Development build (with race detection)
dev-build:
	@echo "Building $(BINARY_NAME) for development..."
	$(GOBUILD) -race $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Lint the code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo >&2 "golangci-lint is required but not installed. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; exit 1; }
	golangci-lint run

# Generate Go modules graph
mod-graph:
	@echo "Generating module dependency graph..."
	$(GOMOD) graph

# Update dependencies
update-deps:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy
