# go-tablewriter

[![CI](https://github.com/njchilds90/go-tablewriter/actions/workflows/ci.yml/badge.svg)](https://github.com/njchilds90/go-tablewriter/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-tablewriter.svg)](https://pkg.go.dev/github.com/njchilds90/go-tablewriter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](go.mod)

Zero-dependency Go library for rendering structured data as **plain text, Markdown, CSV, or JSON tables** — with a clean, chainable API designed for both humans and AI agents.

---

## Why go-tablewriter?

Go's stdlib `text/tabwriter` is primitive and produces no structured output. Existing community packages are stale and pre-generics. `go-tablewriter` fills that gap with:

- **4 output formats** — plain (Unicode box), Markdown, CSV, JSON
- **Chainable options API** — no configuration struct mutation
- **Zero external dependencies**
- **Deterministic output** — same input always produces same output
- **Structured error types** — machine-readable, no string parsing required
- **AI-agent friendly** — pure functions, no hidden state, predictable behavior

---

## Install
```bash
go get github.com/njchilds90/go-tablewriter@latest
```

---

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/njchilds90/go-tablewriter"
)

func main() {
    opts := tablewriter.DefaultOptions().
        WithHeaders("Name", "Age", "City").
        WithFormat(tablewriter.FormatPlain)

    t := tablewriter.New(opts)
    t.AddRow("Alice", "30", "New York")
    t.AddRow("Bob",   "25", "Los Angeles")
    t.AddRow("Carol", "35", "Chicago")

    fmt.Println(t.Render())
}
```

Output:
```
┌───────┬─────┬─────────────┐
│ Name  │ Age │ City        │
├───────┼─────┼─────────────┤
│ Alice │ 30  │ New York    │
│ Bob   │ 25  │ Los Angeles │
│ Carol │ 35  │ Chicago     │
└───────┴─────┴─────────────┘
```

---

## Output Formats

### Markdown
```go
opts := tablewriter.DefaultOptions().
    WithHeaders("Name", "Score").
    WithFormat(tablewriter.FormatMarkdown)

out, _ := tablewriter.Render(opts, [][]string{
    {"Alice", "95"},
    {"Bob",   "87"},
})
fmt.Print(out)
```

Output:
```
| Name  | Score |
| ----- | ----- |
| Alice | 95    |
| Bob   | 87    |
```

### CSV
```go
opts := tablewriter.DefaultOptions().
    WithHeaders("ID", "Status").
    WithFormat(tablewriter.FormatCSV)

out, _ := tablewriter.Render(opts, [][]string{
    {"1", "active"},
    {"2", "pending"},
})
fmt.Print(out)
```

Output:
```
ID,Status
1,active
2,pending
```

### JSON
```go
opts := tablewriter.DefaultOptions().
    WithHeaders("Name", "Role").
    WithFormat(tablewriter.FormatJSON)

out, _ := tablewriter.Render(opts, [][]string{
    {"Alice", "admin"},
    {"Bob",   "viewer"},
})
fmt.Print(out)
```

Output:
```json
[
  {"Name": "Alice", "Role": "admin"},
  {"Name": "Bob",   "Role": "viewer"}
]
```

### Simple (no borders)
```go
opts := tablewriter.DefaultOptions().
    WithHeaders("Col1", "Col2").
    WithFormat(tablewriter.FormatSimple)
```

---

## Options

| Option               | Description                                      | Default   |
|----------------------|--------------------------------------------------|-----------|
| `WithHeaders(...)`   | Column headers                                   | none      |
| `WithFormat(...)`    | Output format                                    | FormatPlain |
| `WithAlignments(...)` | Per-column alignment (Left, Right, Center)      | AlignLeft |
| `WithMaxColumnWidth(n)` | Truncate cells longer than n chars (adds ...) | 0 (off)   |
| `WithNullPlaceholder(s)` | String to use for empty cells               | ""        |
| `WithStrictColumnCount()` | Error if row column count mismatches header | false     |

---

## Column Alignment
```go
opts := tablewriter.DefaultOptions().
    WithHeaders("Name", "Amount", "Status").
    WithFormat(tablewriter.FormatPlain).
    WithAlignments(
        tablewriter.AlignLeft,
        tablewriter.AlignRight,
        tablewriter.AlignCenter,
    )
```

---

## Error Handling
```go
out, err := tablewriter.Render(opts, rows)
if err != nil {
    switch err {
    case tablewriter.ErrMissingHeaders:
        // JSON format requires headers
    case tablewriter.ErrColumnMismatch:
        // Strict mode: row doesn't match header count
    default:
        // Unexpected error
    }
}
```

---

## Stateless Convenience Function
```go
out, err := tablewriter.Render(
    tablewriter.DefaultOptions().WithHeaders("A","B").WithFormat(tablewriter.FormatMarkdown),
    [][]string{{"x","y"},{"1","2"}},
)
```

---

## Package-Level Reference

See full GoDoc at: https://pkg.go.dev/github.com/njchilds90/go-tablewriter

---

## License

MIT © Nicholas John Childs
