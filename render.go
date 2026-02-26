package tablewriter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"
)

func render(opts Options, rows [][]string) (string, error) {
	switch opts.Format {
	case FormatMarkdown:
		return renderMarkdown(opts, rows)
	case FormatCSV:
		return renderCSV(opts, rows)
	case FormatJSON:
		return renderJSON(opts, rows)
	case FormatSimple:
		return renderSimple(opts, rows)
	default:
		return renderPlain(opts, rows)
	}
}

// colWidths computes the max display width of each column across headers + rows.
func colWidths(opts Options, rows [][]string) []int {
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
	return widths
}

func applyCellOpts(v string, opts Options) string {
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
	return v
}

func alignCell(s string, width int, align Alignment) string {
	slen := utf8.RuneCountInString(s)
	pad := width - slen
	if pad <= 0 {
		return s
	}
	switch align {
	case AlignRight:
		return strings.Repeat(" ", pad) + s
	case AlignCenter:
		left := pad / 2
		right := pad - left
		return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
	default:
		return s + strings.Repeat(" ", pad)
	}
}

func getAlign(opts Options, col int) Alignment {
	if col < len(opts.Alignments) {
		return opts.Alignments[col]
	}
	return AlignLeft
}

func renderPlain(opts Options, rows [][]string) (string, error) {
	widths := colWidths(opts, rows)
	if len(widths) == 0 {
		return "", nil
	}
	var b bytes.Buffer

	// top border
	b.WriteString(buildHorizontalBorder(widths, "┌", "┬", "┐", "─"))

	// headers
	if len(opts.Headers) > 0 {
		b.WriteString("\n│")
		for i, h := range opts.Headers {
			h = applyCellOpts(h, opts)
			b.WriteString(" " + alignCell(h, widths[i], getAlign(opts, i)) + " │")
		}
		// fill missing header cols
		for i := len(opts.Headers); i < len(widths); i++ {
			b.WriteString(" " + strings.Repeat(" ", widths[i]) + " │")
		}
		b.WriteString("\n")
		b.WriteString(buildHorizontalBorder(widths, "├", "┼", "┤", "─"))
	}

	// rows
	for _, r := range rows {
		b.WriteString("\n│")
		for i := 0; i < len(widths); i++ {
			v := ""
			if i < len(r) {
				v = r[i]
			}
			v = applyCellOpts(v, opts)
			b.WriteString(" " + alignCell(v, widths[i], getAlign(opts, i)) + " │")
		}
	}

	// bottom border
	b.WriteString("\n")
	b.WriteString(buildHorizontalBorder(widths, "└", "┴", "┘", "─"))
	b.WriteString("\n")
	return b.String(), nil
}

func buildHorizontalBorder(widths []int, left, mid, right, fill string) string {
	var b strings.Builder
	b.WriteString(left)
	for i, w := range widths {
		b.WriteString(strings.Repeat(fill, w+2))
		if i < len(widths)-1 {
			b.WriteString(mid)
		}
	}
	b.WriteString(right)
	return b.String()
}

func renderMarkdown(opts Options, rows [][]string) (string, error) {
	widths := colWidths(opts, rows)
	if len(widths) == 0 {
		return "", nil
	}
	var b bytes.Buffer

	// header row
	b.WriteString("|")
	headers := opts.Headers
	for i := 0; i < len(widths); i++ {
		h := ""
		if i < len(headers) {
			h = headers[i]
		}
		h = applyCellOpts(h, opts)
		b.WriteString(" " + alignCell(h, widths[i], getAlign(opts, i)) + " |")
	}
	b.WriteString("\n")

	// separator row
	b.WriteString("|")
	for i, w := range widths {
		align := getAlign(opts, i)
		sep := strings.Repeat("-", w)
		switch align {
		case AlignRight:
			b.WriteString(" " + sep + ":|")
		case AlignCenter:
			b.WriteString(":" + sep + ":|")
		default:
			b.WriteString(" " + sep + " |")
		}
	}
	b.WriteString("\n")

	// data rows
	for _, r := range rows {
		b.WriteString("|")
		for i := 0; i < len(widths); i++ {
			v := ""
			if i < len(r) {
				v = r[i]
			}
			v = applyCellOpts(v, opts)
			b.WriteString(" " + alignCell(v, widths[i], getAlign(opts, i)) + " |")
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}

func renderCSV(opts Options, rows [][]string) (string, error) {
	var b bytes.Buffer
	if len(opts.Headers) > 0 {
		b.WriteString(csvRow(opts.Headers))
		b.WriteString("\n")
	}
	for _, r := range rows {
		cells := make([]string, len(r))
		for i, v := range r {
			cells[i] = applyCellOpts(v, opts)
		}
		b.WriteString(csvRow(cells))
		b.WriteString("\n")
	}
	return b.String(), nil
}

func csvRow(cols []string) string {
	escaped := make([]string, len(cols))
	for i, c := range cols {
		if strings.ContainsAny(c, `",`+"\n") {
			c = `"` + strings.ReplaceAll(c, `"`, `""`) + `"`
		}
		escaped[i] = c
	}
	return strings.Join(escaped, ",")
}

func renderJSON(opts Options, rows [][]string) (string, error) {
	if len(opts.Headers) == 0 {
		return "", ErrMissingHeaders
	}
	result := make([]map[string]string, 0, len(rows))
	for _, r := range rows {
		obj := make(map[string]string, len(opts.Headers))
		for i, h := range opts.Headers {
			v := ""
			if i < len(r) {
				v = applyCellOpts(r[i], opts)
			}
			obj[h] = v
		}
		result = append(result, obj)
	}
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("tablewriter: JSON marshal failed: %w", err)
	}
	return string(data) + "\n", nil
}

func renderSimple(opts Options, rows [][]string) (string, error) {
	widths := colWidths(opts, rows)
	if len(widths) == 0 {
		return "", nil
	}
	var b bytes.Buffer

	if len(opts.Headers) > 0 {
		for i := 0; i < len(widths); i++ {
			h := ""
			if i < len(opts.Headers) {
				h = opts.Headers[i]
			}
			h = applyCellOpts(h, opts)
			if i > 0 {
				b.WriteString("  ")
			}
			b.WriteString(alignCell(h, widths[i], getAlign(opts, i)))
		}
		b.WriteString("\n")
		for i, w := range widths {
			if i > 0 {
				b.WriteString("  ")
			}
			b.WriteString(strings.Repeat("-", w))
		}
		b.WriteString("\n")
	}

	for _, r := range rows {
		for i := 0; i < len(widths); i++ {
			v := ""
			if i < len(r) {
				v = r[i]
			}
			v = applyCellOpts(v, opts)
			if i > 0 {
				b.WriteString("  ")
			}
			b.WriteString(alignCell(v, widths[i], getAlign(opts, i)))
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}
