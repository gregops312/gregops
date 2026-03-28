# GREGOPS - Go CLI with Cobra

A command-line interface built with Go and the Cobra framework.

## Build

```bash
# Using Makefile (recommended)
make build
```

## Install

```bash
# Using Makefile (installs to /usr/local/bin)
make install
```

## Develop

The project includes a Makefile with common development tasks:

```bash
# Build the project
make build

# Run tests
make test

# Clean build artifacts
make clean

# Format and vet code
make fmt
make vet

# Install to system
make install

# Build for multiple platforms
make build-all

# Show all available commands
make help

# Using Makefile
make test

# With coverage report
make test-cover
```

### Building

```bash
# Using Go
go build -o gregops

# Using Makefile (recommended)
make build

# Development build with race detection
make dev-build
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
