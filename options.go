package tablewriter

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

// WithHeaders returns a copy of Options with the given headers set.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithHeaders("Name", "Score")
func (o Options) WithHeaders(headers ...string) Options {
	o.Headers = headers
	return o
}

// WithFormat returns a copy of Options with the given format set.
//
// Example:
//
//	opts := tablewriter.DefaultOptions().WithFormat(tablewriter.FormatCSV)
func (o Options) WithFormat(f Format) Options {
	o.Format = f
	return o
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
func (o Options) WithMaxColumnWidth(w int) Options {
	o.MaxColumnWidth = w
	return o
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
