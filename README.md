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

### Example Usage with Context
```go
func main() {
    ctx := context.Background()
    opts := tablewriter.DefaultOptions().
        WithHeaders("Name", "Age", "City").
        WithFormat(tablewriter.FormatPlain)

    t := tablewriter.New(opts)
    t.AddRow("Alice", "30", "New York")
    t.AddRow("Bob",   "25", "Los Angeles")
    t.AddRow("Carol", "35", "Chicago")

    result, err := t.RenderWithContext(ctx)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}
```

---

## API Documentation
For more information on the available functions and types, please refer to the [Go Reference documentation](https://pkg.go.dev/github.com/njchilds90/go-tablewriter).