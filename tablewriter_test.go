package tablewriter_test

import (
	"strings"
	"testing"

	"github.com/njchilds90/go-tablewriter"
)

func TestRenderPlain(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"Name", "Age", "City"},
		Format:  tablewriter.FormatPlain,
	}
	rows := [][]string{
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "Alice") {
		t.Error("expected 'Alice' in output")
	}
	if !strings.Contains(out, "Los Angeles") {
		t.Error("expected 'Los Angeles' in output")
	}
	if !strings.Contains(out, "Name") {
		t.Error("expected 'Name' header in output")
	}
}

func TestRenderMarkdown(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"Name", "Score"},
		Format:  tablewriter.FormatMarkdown,
	}
	rows := [][]string{
		{"Alice", "95"},
		{"Bob", "87"},
	}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "| Name") {
		t.Errorf("expected markdown to start with '| Name', got: %s", out[:20])
	}
	if !strings.Contains(out, "---") {
		t.Error("expected separator row in markdown output")
	}
}

func TestRenderCSV(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"ID", "Status"},
		Format:  tablewriter.FormatCSV,
	}
	rows := [][]string{
		{"1", "active"},
		{"2", "pending"},
	}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 {
		t.Errorf("expected 3 lines (header + 2 rows), got %d", len(lines))
	}
	if lines[0] != "ID,Status" {
		t.Errorf("expected 'ID,Status', got '%s'", lines[0])
	}
}

func TestRenderJSON(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"Name", "Age"},
		Format:  tablewriter.FormatJSON,
	}
	rows := [][]string{
		{"Alice", "30"},
	}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `"Name"`) {
		t.Error("expected JSON key 'Name'")
	}
	if !strings.Contains(out, `"Alice"`) {
		t.Error("expected JSON value 'Alice'")
	}
}

func TestRenderJSONNoHeaders(t *testing.T) {
	opts := tablewriter.Options{
		Format: tablewriter.FormatJSON,
	}
	rows := [][]string{{"Alice", "30"}}
	_, err := tablewriter.Render(opts, rows)
	if err != tablewriter.ErrMissingHeaders {
		t.Errorf("expected ErrMissingHeaders, got %v", err)
	}
}

func TestRenderSimple(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"Col1", "Col2"},
		Format:  tablewriter.FormatSimple,
	}
	rows := [][]string{{"a", "b"}, {"c", "d"}}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "Col1") {
		t.Error("expected 'Col1' in output")
	}
	if !strings.Contains(out, "----") {
		t.Error("expected separator in simple output")
	}
}

func TestMaxColumnWidth(t *testing.T) {
	opts := tablewriter.Options{
		Headers:        []string{"Text"},
		Format:         tablewriter.FormatPlain,
		MaxColumnWidth: 10,
	}
	rows := [][]string{{"This is a very long string that should be truncated"}}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(out, "very long string that should be truncated") {
		t.Error("expected text to be truncated")
	}
	if !strings.Contains(out, "...") {
		t.Error("expected '...' in truncated output")
	}
}

func TestNullPlaceholder(t *testing.T) {
	opts := tablewriter.Options{
		Headers:         []string{"A", "B"},
		Format:          tablewriter.FormatPlain,
		NullPlaceholder: "N/A",
	}
	rows := [][]string{{"value", ""}}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "N/A") {
		t.Error("expected 'N/A' placeholder for empty cell")
	}
}

func TestStrictColumnCount(t *testing.T) {
	t.Run("mismatch returns error", func(t *testing.T) {
		tbl := tablewriter.New(tablewriter.Options{
			Headers:           []string{"A", "B", "C"},
			StrictColumnCount: true,
		})
		err := tbl.AddRow("x", "y")
		if err != tablewriter.ErrColumnMismatch {
			t.Errorf("expected ErrColumnMismatch, got %v", err)
		}
	})

	t.Run("matching count no error", func(t *testing.T) {
		tbl := tablewriter.New(tablewriter.Options{
			Headers:           []string{"A", "B"},
			StrictColumnCount: true,
		})
		err := tbl.AddRow("x", "y")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}

func TestTableBuilderReset(t *testing.T) {
	tbl := tablewriter.New(tablewriter.Options{
		Headers: []string{"X"},
	})
	_ = tbl.AddRow("one")
	_ = tbl.AddRow("two")
	if tbl.RowCount() != 2 {
		t.Errorf("expected 2 rows, got %d", tbl.RowCount())
	}
	tbl.Reset()
	if tbl.RowCount() != 0 {
		t.Errorf("expected 0 rows after reset, got %d", tbl.RowCount())
	}
}

func TestChainableOptions(t *testing.T) {
	opts := tablewriter.DefaultOptions().
		WithHeaders("Name", "Score").
		WithFormat(tablewriter.FormatMarkdown).
		WithMaxColumnWidth(20).
		WithNullPlaceholder("—").
		WithAlignments(tablewriter.AlignLeft, tablewriter.AlignRight)

	tbl := tablewriter.New(opts)
	_ = tbl.AddRow("Alice", "100")
	out := tbl.Render()
	if !strings.Contains(out, "Alice") {
		t.Error("expected 'Alice' in chained options render")
	}
}

func TestCSVQuoting(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"Field"},
		Format:  tablewriter.FormatCSV,
	}
	rows := [][]string{
		{`value with "quotes" and, comma`},
	}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `""quotes""`) {
		t.Error("expected escaped double-quotes in CSV output")
	}
}

func TestAlignments(t *testing.T) {
	opts := tablewriter.Options{
		Headers:    []string{"Left", "Right", "Center"},
		Format:     tablewriter.FormatPlain,
		Alignments: []tablewriter.Alignment{tablewriter.AlignLeft, tablewriter.AlignRight, tablewriter.AlignCenter},
	}
	rows := [][]string{{"a", "b", "c"}}
	out, err := tablewriter.Render(opts, rows)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "a") || !strings.Contains(out, "b") {
		t.Error("expected cell values in aligned output")
	}
}

func TestEmptyTable(t *testing.T) {
	opts := tablewriter.Options{
		Headers: []string{"A", "B"},
		Format:  tablewriter.FormatPlain,
	}
	out, err := tablewriter.Render(opts, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "A") {
		t.Error("expected headers in empty table")
	}
}

func TestAddRows(t *testing.T) {
	tbl := tablewriter.New(tablewriter.Options{Headers: []string{"A", "B"}})
	err := tbl.AddRows([][]string{{"1", "2"}, {"3", "4"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tbl.RowCount() != 2 {
		t.Errorf("expected 2 rows, got %d", tbl.RowCount())
	}
}
