# go-common

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## Synopsis

The Senzing go-common repository has packages containing
common interface definitions, data structures, helper functions, and test data
used by other Senzing Go language packages.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-common.svg)](https://pkg.go.dev/github.com/senzing-garage/go-common)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-common)](https://goreportcard.com/report/github.com/senzing-garage/go-common)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-common/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/go-common/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/go-common/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/go-common/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/go-common/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/go-common/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/go-common/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/go-common/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/go-common/actions/workflows/go-test-windows.yaml)

## Overview

The following helpers have been created:

### Data structures

- `record` - A package for a Senzing record

### Helper functions

- `g2engineconfigurationjson`:  A package for creating the `SENZING_ENGINE_CONFIGURATION_JSON` Document.

### Test data

- `truthset` - A Go representation of [truth-sets](https://github.com/senzing-garage/truth-sets)

## References

1. [API documentation](https://pkg.go.dev/github.com/senzing-garage/go-common)
1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
