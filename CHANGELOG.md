# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.8] - 2023-07-31

### Changed in 0.2.8

- Added `jsonutil.RedactJson()` function to recursively nullify JSON property values
- Added `jsonutil.RedactJsonWithMap()` function to recursively replace JSON property values
- Added test cases and examples for new redaction functions

## [0.2.7] - 2023-07-27

### Changed in 0.2.7

- Added `jsonutil.NormalizeAndSortJson()` function to sort JSON arrays in addition to normalization
- Added `-j` option to `main()` in `main.go` to allow command-line JSON normalization testing.

## [0.2.6] - 2023-07-26

### Changed in 0.2.6

- Added `jsonutil` package with `IsJson()` and `NormalizeJson` functions
- Added unit tests for `jsonutil`

## [0.2.5] - 2023-07-24

### Changed in 0.2.5

- Option package moved to go-cmdhelping

## [0.2.4] - 2023-07-19

### Changed in 0.2.4

- Added the option package

## [0.2.3] - 2023-07-17

### Changed in 0.2.3

- Remove test for `RESOURCEPATH`

## [0.2.2] - 2023-07-14

### Added in 0.2.2

- `g2engineconfigurationjson.VerifySenzingEngineConfigurationJson()`

## [0.2.1] - 2023-07-13

### Changed in 0.2.1

- Added support for a Senzing directory
- Update dependencies
  - github.com/senzing/go-logging v1.3.1

## [0.2.0] - 2023-07-12

### Added in 0.2.0

- Support for Mac/Windows

## [0.1.5] - 2023-07-06

### Changed in 0.1.5

- Update dependencies
  - github.com/senzing/go-logging v1.3.0

## [0.1.4] - 2023-05-16

### Added to 0.1.4

- Update dependencies
  - github.com/senzing/go-logging v1.2.6
  - github.com/stretchr/testify v1.8.4

## [0.1.3] - 2023-05-11

### Added to 0.1.3

- In `engineconfigurationjsonparser`, RedactedJson()
- Update dependencies
  - github.com/senzing/go-logging v1.2.3

## [0.1.2] - 2023-03-01

### Added to 0.1.2

- `engineconfigurationjsonparser` - Parses SENZING_ENGINE_CONFIGURATION_JSON

## [0.1.1] - 2023-02-10

### Added to 0.1.1

- `engineconfigurationjson` - Simple construction of SENZING_ENGINE_CONFIGURATION_JSON

### Deleted in 0.1.1

- `testrecords01`

## [0.1.0] - 2023-01-25

### Added to 0.1.0

- `truthset` - Truth set data
- `record` - Add Record type
