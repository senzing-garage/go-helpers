# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [0.6.1] - 2024-09-05

### Changed in 0.6.1

- Update documentation
- Refactor to [Go template]

## [0.6.0] - 2024-08-19

### Changed in 0.6.0

- Change from `g2` to `sz`/`er`

## [0.5.2] - 2024-06-10

### Changed in 0.5.2

- Changed `engineconfigurationjson` to `settings`
- Changed `engineconfigurationjsonparser` to `settingsparser`

## [0.5.1] - 2024-04-19

### Changed in 0.5.1

- Update dependencies

## [0.5.0] - 2024-04-05

### Changed in 0.5.0

- Move from "G2" to "Sz" prefix
- Renamed module to `github.com/senzing-garage/go-helpers`

## [0.4.0] - 2024-01-02

### Changed in 0.4.0

- Renamed module to `github.com/senzing-garage/go-common`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-logging v1.4.0

## [0.3.1] - 2023-10-17

### Changed in 0.3.1

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-logging v1.3.3

## [0.3.0] - 2023-08-31

### Fixed in 0.3.0

- In `engineconfigurationjson`, Windows paths

### Removed in 0.3.0

- `engineconfigurationjson.BuildSimpleSystemConfigurationJson()`

## [0.2.14] - 2023-08-24

### Fixed in 0.2.14

- Windows: missing import

## [0.2.13] - 2023-08-11

### Changed in 0.2.13

- Added `fileutil` package with `fileutil.CopyFile()` function
- Added tests and examples for `fileutil.CopyFile()` function

## [0.2.12] - 2023-08-09

### Changed in 0.2.12

- Added `jsonutil.Strip()` function for removing properties from JSON objects.
- Added tests and example for `jsonutil.Strip()`

## [0.2.11] - 2023-08-04

### Changed in 0.2.11

- Refactored `jsonutil` method names to remove the "Json" suffix since it was redundant
- Fixed the `test` target for Make so it is now the same on all platforms
- Modified `jsonutil.Flatten()` method to always return JSON text

## [0.2.10] - 2023-08-03

### Changed in 0.2.10

- Refactored to template

## [0.2.9] - 2023-08-02

### Changed in 0.2.9

- Added `jsonutil.Flatten()` function to coalesce a string and error tuple into a string
- Added test cases and examples for new flatten function

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

- `engineconfigurationjson.VerifySenzingEngineConfigurationJson()`

## [0.2.1] - 2023-07-13

### Changed in 0.2.1

- Added support for a Senzing directory
- Update dependencies
  - github.com/senzing-garage/go-logging v1.3.1

## [0.2.0] - 2023-07-12

### Added in 0.2.0

- Support for Mac/Windows

## [0.1.5] - 2023-07-06

### Changed in 0.1.5

- Update dependencies
  - github.com/senzing-garage/go-logging v1.3.0

## [0.1.4] - 2023-05-16

### Added to 0.1.4

- Update dependencies
  - github.com/senzing-garage/go-logging v1.2.6
  - github.com/stretchr/testify v1.8.4

## [0.1.3] - 2023-05-11

### Added to 0.1.3

- In `engineconfigurationjsonparser`, RedactedJson()
- Update dependencies
  - github.com/senzing-garage/go-logging v1.2.3

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

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
[Go template]: https://github.com/senzing-garage/template-go
