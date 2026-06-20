# kops - Go CLI with Cobra

A command-line interface built with Go and the Cobra framework.

## Building

```bash
# Build the project
make build

# Build for multiple platforms
make build-all

# Install the locally built binary
make install

# Uninstall the binary
make uninstall

# Clean build artifacts
make clean

# Download and tidy dependencies
make deps

# Format, lint, and vet code
make fmt
make lint
make vet

# Generate .gitignore
make gitignore

# Show all available commands
make help

# Print the module dependency graph
make mod-graph

# Run tests
make test

# With coverage report
make test-cover

# Update dependencies
make update-deps

# Development build with race detection
make dev-build

# Using Go
go build -o kops
```

### Installing dependencies

```bash
# Using Go
go mod tidy

# Using Makefile
make deps
```

### Code Quality

```bash
# Format code
make fmt

# Vet code
make vet

# Lint code (requires golangci-lint)
make lint
```
