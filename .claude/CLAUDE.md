# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-helpers is a Go library providing common interface definitions, data structures, helper functions, and test data used by other Senzing Go packages. It is not a standalone application.

## Build and Development Commands

```bash
# Run all tests with formatted output
make test

# Run a single test
go test -v -run TestFunctionName ./package/...

# Run tests for a specific package
go test -v ./settings/...

# Lint the codebase (golangci-lint + govulncheck + cspell)
make lint

# Run individual linters
make golangci-lint
make govulncheck
make cspell

# Auto-fix lint issues
make fix

# Generate test coverage report (opens in browser)
make coverage

# Check coverage against thresholds
make check-coverage

# Update dependencies
make dependencies

# Install development tools
make dependencies-for-development

# Clean build artifacts and caches
make clean
```

## Testing Requirements

Tests require a Senzing SDK installation and test database:

- Set `LD_LIBRARY_PATH=/opt/senzing/er/lib` (Linux)
- Test database: `sqlite3://na:na@nowhere/tmp/sqlite/G2C.db`
- License file needed at `/etc/opt/senzing/g2.lic` or `testdata/senzing-license/g2.lic`

## Package Architecture

- **settings** - Creates `SENZING_ENGINE_CONFIGURATION_JSON` for Senzing runtime configuration. Has OS-specific implementations (`settings_linux.go`, `settings_darwin.go`, `settings_windows.go`).

- **settingsparser** - Parses Senzing engine configuration JSON strings. Handles both single-database and multi-database (clustered) configurations.

- **record** - Validates and parses Senzing records (entity data with DATA_SOURCE, RECORD_ID, and entity attributes).

- **truthset** - Go representation of Senzing truth-set test data (customers, watchlist, reference data).

- **jsonutil** - JSON manipulation utilities (normalize, sort, redact, flatten, truncate, pretty-print).

- **fileutil** - File management utilities.

- **wraperror** - Error wrapping utilities.

- **env** - Environment variable utilities.

- **tls** - TLS configuration utilities with OS-specific implementations.

- **testfixtures** - Fixture records for testing.

## Linting Configuration

Uses extensive golangci-lint configuration at `.github/linters/.golangci.yaml` with 100+ enabled linters. Key exclusions are configured for test files and specific struct types in `exhaustruct`.

## JSON Tag Convention

Uses `upperSnake` case for JSON struct tags (configured in tagliatelle linter).
