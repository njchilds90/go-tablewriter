package tablewriter

// Package tablewriter provides a simple way to render tables in various formats.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// ErrInvalidFormat is returned when an invalid format is provided.
var ErrInvalidFormat = errors.New("invalid format")

// ErrInvalidOptions is returned when invalid options are provided.
var ErrInvalidOptions = errors.New("invalid options")

// render renders a table based on the provided options and rows.
//
// render takes a context, options and rows as input, and returns the rendered table as a string, and an error if any.
func render(ctx context.Context, opts Options, rows [][]string) (string, error) {
	if ctx == nil {
		return "", errors.New("context is nil")
	}
	if opts == nil {
		return "", ErrInvalidOptions
	}
	if rows == nil {
		return "", errors.New("rows is nil")
	}
	switch opts.Format {
	case FormatMarkdown:
		return renderMarkdown(ctx, opts, rows)
	case FormatCSV:
		return renderCSV(ctx, opts, rows)
	case FormatJSON:
		return renderJSON(ctx, opts, rows)
	case FormatSimple:
		return renderSimple(ctx, opts, rows)
	default:
		return "", fmt.Errorf("invalid format: %s", opts.Format)
	}
}

// colWidths computes the max display width of each column across headers + rows.
//
// colWidths takes a context, options and rows as input, and returns the max display width of each column as a slice of integers, and an error if any.
func colWidths(ctx context.Context, opts Options, rows [][]string) ([]int, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	if opts == nil {
		return nil, ErrInvalidOptions
	}
	if rows == nil {
		return nil, errors.New("rows is nil")
	}
	numCols := len(opts.Headers)
	for _, r := range rows {
		if len(r) > numCols {
			numCols = len(r)
		}
	}
	widths := make([]int, numCols)
	for i, h := range opts.Headers {
		w := utf8.RuneCountInString(h)
		if w > widths[i] {
			widths[i] = w
		}
	}
	for _, r := range rows {
		for i, c := range r {
			c = applyCellOpts(c, opts)
			w := utf8.RuneCountInString(c)
			if w > widths[i] {
				widths[i] = w
			}
		}
	}
	if opts.MaxColumnWidth > 0 {
		for i := range widths {
			if widths[i] > opts.MaxColumnWidth {
				widths[i] = opts.MaxColumnWidth
			}
		}
	}
	return widths, nil
}

// applyCellOpts applies the cell options to the given value.
//
// applyCellOpts takes a value, and options as input, and returns the value with the cell options applied, and an error if any.
func applyCellOpts(v string, opts Options) (string, error) {
	if v == "" && opts.NullPlaceholder != "" {
		v = opts.NullPlaceholder
	}
	if opts.MaxColumnWidth > 0 && utf8.RuneCountInString(v) > opts.MaxColumnWidth {
		runes := []rune(v)
		if opts.MaxColumnWidth > 3 {
			v = string(runes[:opts.MaxColumnWidth-3]) + "..."
		} else {
			v = string(runes[:opts.MaxColumnWidth])
		}
	}
	return v, nil
}

// alignCell aligns the cell to the given width and alignment.
//
// alignCell takes a string, width, and alignment as input, and returns the aligned string, and an error if any.
func alignCell(s string, width int, align Alignment) (string, error) {
	slen := utf8.RuneCountInString(s)
	pad := width - slen
	if pad <= 0 {
		return s, nil
	}
	switch align {
	case AlignRight:
		return strings.Repeat(" ", pad) + s, nil
	case AlignCenter:
		left := pad / 2
		right := pad - left
		return strings.Repeat(" ", left) + s + strings.Repeat(" ", right), nil
	default:
		return s + strings.Repeat(" ", pad), nil
	}
}

// getAlign gets the alignment for the given column.
//
// getAlign takes a context, options, and column as input, and returns the alignment for the given column, and an error if any.
func getAlign(ctx context.Context, opts Options, col int) (Alignment, error) {
	if ctx == nil {
		return "", errors.New("context is nil")
	}
	if opts == nil {
		return "", ErrInvalidOptions
	}
	if col < 0 || col >= len(opts.Headers) {
		return "", fmt.Errorf("column out of range: %d", col)
	}
	// alignment logic
	return "", nil
}
