# go-common

## Synopsis

The Senzing go-common repository has packages containing
common interface definitions, data structures, helper functions, and test data
used by other Senzing Go language packages.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-common.svg)](https://pkg.go.dev/github.com/senzing/go-common)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-common)](https://goreportcard.com/report/github.com/senzing/go-common)
[![go-test.yaml](https://github.com/Senzing/go-common/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-common/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/go-common/blob/main/LICENSE)

## Overview

The following helpers have been created:

### Data structures

- `record` - A package for a Senzing record

### Helper functions

- `g2engineconfigurationjson`:  A package for creating the `SENZING_ENGINE_CONFIGURATION_JSON` Document.

### Test data

- `truthset` - A Go representation of [truth-sets](https://github.com/Senzing/truth-sets)

## References

1. [API documentation](https://pkg.go.dev/github.com/senzing/go-common)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
