# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 0.1.0 (2026-10-06)


### Features

* 1st commit ([9c12765](https://github.com/deploymenttheory/go-sdk-nexthink/commit/9c1276566601c6ffc0e64f3c0f2e92a33250dae7))
* add full LCRUD for web management resources ([4a4eac2](https://github.com/deploymenttheory/go-sdk-nexthink/commit/4a4eac2d8dcf6221c31b399b2b444588e65e4150))
* add full LCRUD for web management resources ([4117e0c](https://github.com/deploymenttheory/go-sdk-nexthink/commit/4117e0c8c99b892f811aeccd2736a485246fc55e))
* add headless local password authentication ([#56](https://github.com/deploymenttheory/go-sdk-nexthink/issues/56)) ([611dfac](https://github.com/deploymenttheory/go-sdk-nexthink/commit/611dfaccf2d525ffeb304b748b41eeaa978243fe))
* add web content lifecycles and fix DELETE payloads ([#51](https://github.com/deploymenttheory/go-sdk-nexthink/issues/51)) ([85cf633](https://github.com/deploymenttheory/go-sdk-nexthink/commit/85cf6333c23271f739d86b0127e653bcc7699fe8))
* complete captured browser API operations and examples ([#53](https://github.com/deploymenttheory/go-sdk-nexthink/issues/53)) ([d80327c](https://github.com/deploymenttheory/go-sdk-nexthink/commit/d80327cbd68e1094926d5c4bf7cb8e2500c0d22c))
* complete web API discovery leads and refresh quick start ([#55](https://github.com/deploymenttheory/go-sdk-nexthink/issues/55)) ([d47bab6](https://github.com/deploymenttheory/go-sdk-nexthink/commit/d47bab6f31933726e6c24b44b2fa136e14cccba7))
* enhance NQL service with new query builder, templates, and result processing features ([670f602](https://github.com/deploymenttheory/go-sdk-nexthink/commit/670f602d34ce3e8f6eed41c1ad885c0739b82cbe))
* expand browser identity support and helper APIs ([#54](https://github.com/deploymenttheory/go-sdk-nexthink/issues/54)) ([00b766e](https://github.com/deploymenttheory/go-sdk-nexthink/commit/00b766e854a993d9f34cf63a45f4b9ed8840e45d))
* expand web management APIs and complete examples ([4eaa429](https://github.com/deploymenttheory/go-sdk-nexthink/commit/4eaa429472b31ce6b873721c1d0a4938b9cd6d67))
* expand web management resources and complete examples ([9289d59](https://github.com/deploymenttheory/go-sdk-nexthink/commit/9289d59911e48b86f69fa52c7772d6b634b36af9))
* implement API versioning and refactor NQL service methods for clarity ([cfbc12d](https://github.com/deploymenttheory/go-sdk-nexthink/commit/cfbc12ddb5253cac72588b41027d83fb16a358a4))
* unify Nexthink API families and validate lab contracts ([6dd2088](https://github.com/deploymenttheory/go-sdk-nexthink/commit/6dd2088bfdd8274603aa5eb5ffaef2db361ad7b3))
* **web:** add integrations, knowledge uploads and nested content APIs ([#52](https://github.com/deploymenttheory/go-sdk-nexthink/issues/52)) ([b7b445c](https://github.com/deploymenttheory/go-sdk-nexthink/commit/b7b445c7f506da546fafb37cb42756ea553b1b03))


### Bug Fixes

* check Go packages directly in lint workflow ([14a33d2](https://github.com/deploymenttheory/go-sdk-nexthink/commit/14a33d220d45ef44c9941ab940c41c7f46df8461))
* run CI linters in Go module mode ([bf29d4a](https://github.com/deploymenttheory/go-sdk-nexthink/commit/bf29d4ae765bfa6a219922f87657640cf0cf2152))
* select module linting without conflicting flags ([776e869](https://github.com/deploymenttheory/go-sdk-nexthink/commit/776e8694e9fee80953beb3b560e1e1f7ac2e3c78))
* validate Nexthink API contracts and add experimental services ([6de5575](https://github.com/deploymenttheory/go-sdk-nexthink/commit/6de5575d0c5a7d533bd4a8b68ef625f4fd01ad1b))

## [Unreleased]

### Added

#### Browser integrations and content operations

- Add knowledge base uploads, legacy connector configuration/secrets, webhook configuration/tests, and Data Exporter configuration/tests through the shared `WebAPI` client.
- Add connector test polling, workflow connector/credential views, nested dashboard mutations and import/export/duplicate, checklist import/export/fields, and monitor import/export/library export.
- Add 56 runnable examples and synthetic JSON wire-contract fixtures with live lab validation.
- Fix raw response handling for plain-text save acknowledgments while retaining JSON validation for typed results.


#### NQL Service Enhancements

- **Query Builder**: Added fluent API for programmatic NQL query construction with type safety and validation
  - Method chaining for readable query construction
  - Support for all NQL constructs (table selection, time ranges, filters, aggregations)
  - Built-in validation with detailed error messages
  - IDE auto-completion support

- **Query Templates**: Added 18 pre-built query templates for common scenarios
  - Device health monitoring (crashes, memory usage, boot time)
  - User experience analysis (web errors, collaboration quality)
  - Application performance (error rates, crash analysis)
  - DEX score analysis (overall, by platform, low score users, component impact)
  - Network connectivity and performance monitoring
  - Workflow and remote action metrics

- **Result Set Processing**: Added type-safe result processing helpers
  - V1 and V2 result set wrappers with type-safe getters
  - Filter and map operations for data transformation
  - Row iteration utilities
  - V1 to V2 format conversion
  - JSON export capabilities

- **Export Workflow**: Added simplified workflow for large data exports
  - One-line export methods (ExportToCSV, ExportToJSON)
  - Progress tracking with customizable callbacks
  - Automatic polling and completion detection
  - Configurable timeouts and intervals
  - Human-readable result sizes and progress reporting

- **Data Model Constants**: Added comprehensive constants for type-safe query construction
  - 50+ table name constants
  - 100+ field name constants
  - 50+ value enumerations (platforms, hardware types, experience levels, etc.)
  - Time selection constants and helpers
  - Operator and function constants

- **Time Selection Helpers**: Added fluent API for time range specification
  - Predefined time constants (Past7Days, Past24Hours, etc.)
  - Relative and absolute time range builders
  - Time granularity constants for aggregations
  - High-resolution support for VDI data

- **Query Validation**: Added comprehensive client-side query validation
  - Syntax and structure validation
  - Operator compatibility checking
  - Comment balance validation
  - Detailed error reporting

- **Metadata Extraction**: Added helpers for execution and performance metrics
  - Query execution time tracking
  - Response duration and size metrics
  - Rate limit information extraction
  - Row count and status tracking

#### Examples

- Added QueryBuilder example demonstrating fluent query construction
- Added Templates example showing all 18 pre-built templates
- Added ResultSetProcessing example with type-safe data access patterns
- Added ExportWorkflow example with progress tracking
- Added ComprehensiveExample combining all enhancements

#### Documentation

- Added NQL Query Building Guide (complete guide to query builder)
- Added NQL Result Processing Guide (working with query results)
- Added NQL Export Workflow Guide (large data exports)
- Added NQL Templates Guide (pre-built query templates)
- Added NQL Best Practices Guide (optimization and patterns)
- Added NQL API Reference (complete reference documentation)
- Added NQL Enhancements summary document
- Updated README with comprehensive NQL enhancements section

### Changed

- Enhanced NQL service with developer-friendly features
- Updated README documentation to include new NQL capabilities

### Fixed

- N/A

## [1.1.0] - 2021-06-23

### Added

- Added x [@your_username](https://github.com/your_username)

### Changed

- Changed y [@your_username](https://github.com/your_username)

## [1.0.0] - 2021-06-20

### Added

- Inititated y [@your_username](https://github.com/your_username)
- Inititated z [@your_username](https://github.com/your_username)
