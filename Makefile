# Variables
BINARY_NAME=kops
MAIN_PACKAGE=.
BUILD_DIR=build

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

.PHONY: all build build-all clean deps dev-build fmt gitignore help install lint mod-graph test test-cover uninstall update-deps vet

# Default target
all: clean deps fmt vet test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) $(BUILD_FLAGS) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Install the locally built binary
install: build
	@GOPATH="$$(go env GOPATH)"; \
	if [ -z "$$GOPATH" ]; then \
		echo "GOPATH is not set; cannot install $(BINARY_NAME)."; \
		exit 1; \
	fi; \
	echo "Installing $(BINARY_NAME) to $$GOPATH/bin..."; \
	mkdir -p "$$GOPATH/bin"; \
	install -m 755 $(BINARY_NAME) "$$GOPATH/bin/$(BINARY_NAME)"

# Remove the installed binary
uninstall:
	@GOPATH="$$(go env GOPATH)"; \
	if [ -z "$$GOPATH" ]; then \
		echo "GOPATH is not set; nothing to remove."; \
		exit 0; \
	fi; \
	echo "Removing $$GOPATH/bin/$(BINARY_NAME)..."; \
	rm -f "$$GOPATH/bin/$(BINARY_NAME)"

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

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	gofmt -s -w .

# Generate .gitignore
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

# Show available targets
help:
	@echo "Available targets:"
	@echo "  all          Clean, download deps, format, vet, test, and build"
	@echo "  build        Build the binary"
	@echo "  build-all    Build binaries for multiple platforms"
	@echo "  clean        Remove build artifacts"
	@echo "  deps         Download and tidy Go modules"
	@echo "  dev-build    Build with race detection"
	@echo "  fmt          Format code"
	@echo "  gitignore    Generate .gitignore"
	@echo "  help         Show this help"
	@echo "  install      Install the binary to GOPATH/bin"
	@echo "  lint         Run golangci-lint"
	@echo "  mod-graph    Print the module dependency graph"
	@echo "  test         Run tests"
	@echo "  test-cover   Run tests with coverage"
	@echo "  uninstall    Remove the installed binary"
	@echo "  update-deps  Update Go dependencies"
	@echo "  vet          Run go vet"

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

# Update dependencies
update-deps:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

# Vet code
vet:
	@echo "Vetting code..."
	$(GOCMD) vet ./...
