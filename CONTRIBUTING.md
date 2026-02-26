# Contributing

Thank you for your interest in contributing to go-tablewriter!

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a branch: `git checkout -b my-feature`
4. Make your changes
5. Run tests: `go test -v -race ./...`
6. Run vet: `go vet ./...`
7. Commit and push your branch
8. Open a Pull Request

## Guidelines

- All new exported symbols must have GoDoc comments with an `Example:` block
- All new features require table-driven tests with good coverage
- Keep zero external dependencies — stdlib only
- Follow idiomatic Go naming and patterns
- Maintain deterministic, pure-function behavior wherever possible

## Bug Reports

Open a GitHub Issue with:
- Go version (`go version`)
- Minimal reproducing example
- Expected vs actual output

## License

By contributing you agree your contributions are licensed under MIT.
```

---

## 4. Release & Verification Instructions
```
RELEASE STEPS — GitHub Web UI Only
====================================

1. VERIFY CI PASSES
   - Go to: https://github.com/njchilds90/go-tablewriter/actions
   - Confirm the latest CI run is green on all three Go versions

2. CREATE TAG v1.0.0
   - Go to: https://github.com/njchilds90/go-tablewriter
   - Click "Releases" in the right sidebar (or go to /releases)
   - Click "Create a new release"
   - Click "Choose a tag" → type: v1.0.0 → click "Create new tag: v1.0.0 on publish"
   - Target branch: main

3. FILL IN RELEASE DETAILS
   - Title: v1.0.0 — Initial Release
   - Description (paste this):

     ## go-tablewriter v1.0.0

     Zero-dependency Go library for rendering structured data as plain text,
     Markdown, CSV, and JSON tables.

     ### Features
     - 5 output formats: Plain (Unicode box), Markdown, CSV, JSON, Simple
     - Chainable options API
     - Per-column alignment (left, center, right)
     - MaxColumnWidth with truncation
     - NullPlaceholder for empty cells
     - StrictColumnCount validation
     - Structured error types (ErrMissingHeaders, ErrColumnMismatch)
     - Package-level Render() convenience function
     - Zero external dependencies
     - Full GoDoc + table-driven tests + CI (Go 1.21–1.23)

   - ✅ Set as the latest release
   - Click "Publish release"

4. TRIGGER pkg.go.dev INDEXING (takes ~5–10 min)
   Visit this URL in your browser to request indexing:
   https://sum.golang.org/lookup/github.com/njchilds90/go-tablewriter@v1.0.0

5. VERIFY ON PKG.GO.DEV
   After ~10 minutes, visit:
   https://pkg.go.dev/github.com/njchilds90/go-tablewriter

   You should see full GoDoc with all exported symbols, examples, and README.

SEMANTIC VERSIONING GUIDANCE
==============================
v1.0.x — Bug fixes, doc improvements, no API changes
v1.1.0 — New formats, new options fields (backward compatible additions)
v2.0.0 — Only if a breaking API change is required (avoid as long as possible)
