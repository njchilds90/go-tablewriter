
// Package tablewriter provides a writer for table-based data with various format options.
//
// Users can configure table writer options using the DefaultOptions function and
// a series of with- functions to modify the options.
package tablewriter

import (
	"context"
	"errors"
	"fmt"
	"io"
)

// ErrInvalidOptions is returned when options are invalid.
var ErrInvalidOptions = errors.New("invalid options")

// ErrInvalidColumnWidth is returned when a column width is invalid.
var ErrInvalidColumnWidth = errors.New("invalid column width")

// ValidFormat determines if a given format is valid.
func ValidFormat(f Format) bool {
	switch f {
	case FormatPlain, FormatMarkdown, FormatCSV:
		return true
	default:
		return false
	}
}

// DefaultOptions returns a sensible default Options configuration.
//
// Example:
//
//	opts := tablewriter.DefaultOptions()
//	opts.Format = tablewriter.FormatMarkdown
func DefaultOptions() Options {
	return Options{
		Format:          FormatPlain,
		MaxColumnWidth:  0,
		NullPlaceholder: "",
	}
}

// Options represents a set of configurable options for a table writer.
//
// The Zero value of Options is invalid and will cause errors if used directly.
type Options struct {
	Format          Format
	Headers         []string
	Alignments      []Alignment
	MaxColumnWidth  int
	NullPlaceholder string
	StrictColumnCount bool
}

// WithHeaders returns a copy of Options with the given headers set.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithHeaders("Name", "Score", "Age")
func (o Options) WithHeaders(headers ...string) Options {
	if len(headers) == 0 {
		return o
	}
	o.Headers = headers
	return o
}

// WithFormat returns a copy of Options with the given format set.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithFormat(tablewriter.FormatCSV)
func (o Options) WithFormat(f Format) (Options, error) {
	if !ValidFormat(f) {
		return o, fmt.Errorf("invalid format: %w", ErrInvalidOptions)
	}
	o.Format = f
	return o, nil
}

// WithAlignments returns a copy of Options with the given per-column alignments.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithAlignments(tablewriter.AlignLeft, tablewriter.AlignRight)
func (o Options) WithAlignments(a ...Alignment) Options {
	o.Alignments = a
	return o
}

// WithMaxColumnWidth returns a copy of Options with the given max column width.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithMaxColumnWidth(20)
func (o Options) WithMaxColumnWidth(w int) (Options, error) {
	if w < 0 {
		return o, fmt.Errorf("invalid max column width: %w", ErrInvalidColumnWidth)
	}
	o.MaxColumnWidth = w
	return o, nil
}

// WithNullPlaceholder returns a copy of Options with the given null placeholder string.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithNullPlaceholder("N/A")
func (o Options) WithNullPlaceholder(s string) Options {
	o.NullPlaceholder = s
	return o
}

// WithStrictColumnCount returns a copy of Options with strict column count validation enabled.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithStrictColumnCount()
func (o Options) WithStrictColumnCount() Options {
	o.StrictColumnCount = true
	return o
}

// NewTableWriter returns a new table writer using the given options.
func NewTableWriter(ctx context.Context, w io.Writer, opts Options) (io.Writer, error) {
	if opts.Format == 0 {
		return nil, fmt.Errorf("invalid options: %w", ErrInvalidOptions)
	}
	// implement table writer logic here
	return w, nil
}
