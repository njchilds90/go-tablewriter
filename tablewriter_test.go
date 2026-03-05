package tablewriter_test

// Package tablewriter_test contains tests for the tablewriter package.
import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/njchilds90/go-tablewriter"
)

// ErrInvalidFormat is a sentinel error indicating an invalid format.
var ErrInvalidFormat = errors.New("invalid format")

// ErrInvalidOptions is a sentinel error indicating invalid options.
var ErrInvalidOptions = errors.New("invalid options")

func TestRenderPlain(t *testing.T) {
	// Test cases for RenderPlain.
	tests := []struct {
		name     string
		headers  []string
		rows     [][]string
		wantOut  string
		wantErr  bool
	}{{
		"valid input",
		[]string{"Name", "Age", "City"},
		[][]string{{"Alice", "30", "New York"}, {"Bob", "25", "Los Angeles"}},
		"Alice\nBob\n30\n25\nNew York\nLos Angeles",
		false,
	},{
		"invalid format",
		[]string{"Name"},
		[][]string{{"Alice"}},
		"",
		true,
	},{
		"empty input",
		[]string{},
		[][]string{},
		"",
		false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			opts := tablewriter.Options{
				Headers: tt.headers,
				Format:  tablewriter.FormatPlain,
			}
			out, err := tablewriter.RenderWithContext(ctx, opts, tt.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderPlain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(out, tt.wantOut) {
				t.Errorf("RenderPlain() got = %v, want %v", out, tt.wantOut)
			}
		})
	}
}

func TestRenderMarkdown(t *testing.T) {
		// Test cases for RenderMarkdown.
		tests := []struct {
			name     string
			headers  []string
			rows     [][]string
			wantOut  string
			wantErr  bool
		}{{
			"valid input",
			[]string{"Name", "Score"},
			[][]string{{"Alice", "95"}, {"Bob", "87"}},
			"| Name | Score |\n| --- | --- |\n| Alice | 95 |\n| Bob | 87 |",
			false,
		},{
			"invalid format",
			[]string{"Name"},
			[][]string{{"Alice"}},
			"",
			true,
		},{
			"empty input",
			[]string{},
			[][]string{},
			"",
			false,
		}}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctx := context.Background()
				opts := tablewriter.Options{
					Headers: tt.headers,
					Format:  tablewriter.FormatMarkdown,
				}
				out, err := tablewriter.RenderWithContext(ctx, opts, tt.rows)
				if (err != nil) != tt.wantErr {
					t.Errorf("RenderMarkdown() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !strings.Contains(out, tt.wantOut) {
					t.Errorf("RenderMarkdown() got = %v, want %v", out, tt.wantOut)
				}
			})
		}
}

func TestRenderCSV(t *testing.T) {
		// Test cases for RenderCSV.
		tests := []struct {
			name     string
			headers  []string
			rows     [][]string
			wantOut  string
			wantErr  bool
		}{{
			"valid input",
			[]string{"ID", "Status"},
			[][]string{{"1", "active"}, {"2", "pending"}},
			"ID,Status\n1,active\n2,pending",
			false,
		},{
			"invalid format",
			[]string{"Name"},
			[][]string{{"Alice"}},
			"",
			true,
		},{
			"empty input",
			[]string{},
			[][]string{},
			"",
			false,
		}}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctx := context.Background()
				opts := tablewriter.Options{
					Headers: tt.headers,
					Format:  tablewriter.FormatCSV,
				}
				out, err := tablewriter.RenderWithContext(ctx, opts, tt.rows)
				if (err != nil) != tt.wantErr {
					t.Errorf("RenderCSV() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !strings.Contains(out, tt.wantOut) {
					t.Errorf("RenderCSV() got = %v, want %v", out, tt.wantOut)
				}
			})
		}
}

func BenchmarkRenderPlain(b *testing.B) {
	// Benchmark RenderPlain.
	ctx := context.Background()
	opts := tablewriter.Options{
		Headers: []string{"Name", "Age", "City"},
		Format:  tablewriter.FormatPlain,
	}
	rows := [][]string{{"Alice", "30", "New York"}, {"Bob", "25", "Los Angeles"}}
	for i := 0; i < b.N; i++ {
		tablewriter.RenderWithContext(ctx, opts, rows)
	}
}

func ExampleRenderPlain() {
	// Example usage of RenderPlain.
	ctx := context.Background()
	opts := tablewriter.Options{
		Headers: []string{"Name", "Age", "City"},
		Format:  tablewriter.FormatPlain,
	}
	rows := [][]string{{"Alice", "30", "New York"}, {"Bob", "25", "Los Angeles"}}
	out, _ := tablewriter.RenderWithContext(ctx, opts, rows)
	// Output: Alice\nBob\n30\n25\nNew York\nLos Angeles
