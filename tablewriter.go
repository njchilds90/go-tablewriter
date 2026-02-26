// Package tablewriter renders structured tabular data as plain text, Markdown,
// CSV, or JSON. It is designed for zero dependencies, deterministic output,
// and clean integration with AI agents and CLI tools.
//
// # Quick Start
//
//	t := tablewriter.New(tablewriter.Options{
//	    Headers: []string{"Name", "Age", "City"},
//	    Format:  tablewriter.FormatMarkdown,
//	})
//	t.AddRow("Alice", "30", "NYC")
//	t.AddRow("Bob",   "25", "LA")
//	fmt.Println(t.Render())
package tablewriter

import "errors"

// Format controls the output format of the rendered table.
type Format int

const (
	// FormatPlain renders a plain ASCII table with box-drawing characters.
	FormatPlain Format = iota
	// FormatMarkdown renders a GitHub-flavored Markdown table.
	FormatMarkdown
	// FormatCSV renders comma-separated values.
	FormatCSV
	// FormatJSON renders a JSON array of objects (requires headers).
	FormatJSON
	// FormatSimple renders a minimal table with no borders, only header separator.
	FormatSimple
)

// Alignment controls column text alignment.
type Alignment int

const (
	AlignLeft   Alignment = iota // AlignLeft aligns text to the left (default).
	AlignCenter                  // AlignCenter centers text within the column.
	AlignRight                   // AlignRight aligns text to the right.
)

// ErrMissingHeaders is returned when FormatJSON is used without headers.
var ErrMissingHeaders = errors.New("tablewriter: JSON format requires headers")

// ErrColumnMismatch is returned when a row has a different number of columns than expected.
var ErrColumnMismatch = errors.New("tablewriter: row column count does not match header count")

// Options configures table rendering behavior.
type Options struct {
	// Headers is the list of column names. Optional except for FormatJSON.
	Headers []string

	// Format controls the output format. Defaults to FormatPlain.
	Format Format

	// Alignments sets per-column alignment. If shorter than column count, AlignLeft is used.
	Alignments []Alignment

	// MaxColumnWidth truncates cell values longer than this. 0 = no limit.
	MaxColumnWidth int

	// NullPlaceholder is the string used for empty cells. Defaults to "".
	NullPlaceholder string

	// StrictColumnCount causes AddRow to return an error if column count mismatches.
	StrictColumnCount bool
}

// Table holds headers, rows, and rendering options.
type Table struct {
	opts Options
	rows [][]string
}

// New creates a new Table with the provided Options.
//
// Example:
//
//	t := tablewriter.New(tablewriter.Options{
//	    Headers: []string{"ID", "Status"},
//	    Format:  tablewriter.FormatMarkdown,
//	})
func New(opts Options) *Table {
	return &Table{opts: opts}
}

// AddRow appends a row of string values to the table.
// Returns ErrColumnMismatch if StrictColumnCount is true and counts differ.
//
// Example:
//
//	err := t.AddRow("1", "active")
func (t *Table) AddRow(cols ...string) error {
	if t.opts.StrictColumnCount && len(t.opts.Headers) > 0 {
		if len(cols) != len(t.opts.Headers) {
			return ErrColumnMismatch
		}
	}
	row := make([]string, len(cols))
	copy(row, cols)
	t.rows = append(t.rows, row)
	return nil
}

// AddRows appends multiple rows at once.
// Returns the first error encountered if StrictColumnCount is set.
//
// Example:
//
//	err := t.AddRows([][]string{{"Alice", "30"}, {"Bob", "25"}})
func (t *Table) AddRows(rows [][]string) error {
	for _, r := range rows {
		if err := t.AddRow(r...); err != nil {
			return err
		}
	}
	return nil
}

// Render returns the formatted table as a string.
// Returns an error string prefixed with "tablewriter error:" if rendering fails.
//
// Example:
//
//	output, err := t.RenderErr()
//	fmt.Println(output)
func (t *Table) Render() string {
	s, _ := t.RenderErr()
	return s
}

// RenderErr returns the formatted table and any rendering error.
//
// Example:
//
//	output, err := t.RenderErr()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (t *Table) RenderErr() (string, error) {
	return render(t.opts, t.rows)
}

// Reset clears all rows while preserving options and headers.
//
// Example:
//
//	t.Reset()
func (t *Table) Reset() {
	t.rows = nil
}

// RowCount returns the number of data rows currently in the table.
//
// Example:
//
//	n := t.RowCount()
func (t *Table) RowCount() int {
	return len(t.rows)
}

// Render is a package-level convenience function. It creates a table with the
// given options and rows and returns the rendered string.
//
// Example:
//
//	out, err := tablewriter.Render(
//	    tablewriter.Options{Headers: []string{"A","B"}, Format: tablewriter.FormatMarkdown},
//	    [][]string{{"x","y"},{"1","2"}},
//	)
func Render(opts Options, rows [][]string) (string, error) {
	return render(opts, rows)
}
