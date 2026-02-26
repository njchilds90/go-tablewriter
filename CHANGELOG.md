# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-02-26

### Added
- `FormatPlain` — Unicode box-drawing table renderer
- `FormatMarkdown` — GitHub-flavored Markdown table renderer
- `FormatCSV` — RFC 4180-compliant CSV renderer with proper quoting
- `FormatJSON` — JSON array-of-objects renderer
- `FormatSimple` — Borderless table with header separator only
- `Table` builder with `AddRow`, `AddRows`, `Reset`, `RowCount`, `Render`, `RenderErr`
- Package-level `Render(opts, rows)` convenience function
- Chainable `Options` API: `WithHeaders`, `WithFormat`, `WithAlignments`, `WithMaxColumnWidth`, `WithNullPlaceholder`, `WithStrictColumnCount`
- Per-column alignment: `AlignLeft`, `AlignCenter`, `AlignRight`
- `MaxColumnWidth` with `...` truncation
- `NullPlaceholder` for empty cells
- `StrictColumnCount` with `ErrColumnMismatch` error type
- `ErrMissingHeaders` for JSON format without headers
- Full GoDoc with examples on all exported symbols
- Table-driven tests with race detector coverage
- GitHub Actions CI (Go 1.21, 1.22, 1.23)
- Zero external dependencies
