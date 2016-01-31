package genmarkdown

import (
	"bytes"
	"fmt"
	"io"
	"text/tabwriter"
)

// Row represents a table row.
type Row struct {
	Values []string
}

// Table represents a markdown table. The table will be formatted in a
// human-friendly manner; all rows are aligned and padded properly.
type Table struct {
	Columns []string
	Rows    []Row
}

// Markdown implements the Node interface.
func (me Table) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	w := tabwriter.NewWriter(buf, 1, 4, 0, ' ', 0)
	for _, col := range me.Columns {
		fmt.Fprintf(w, "| %s \t", col)
	}
	fmt.Fprintf(w, "|\n")
	for _, _ = range me.Columns {
		fmt.Fprintf(w, "| --- \t")
	}
	fmt.Fprintf(w, "|\n")
	for _, row := range me.Rows {
		for _, val := range row.Values {
			fmt.Fprintf(w, "| %s \t", val)
		}
		fmt.Fprintf(w, "|\n")
	}
	w.Flush()
	buf.WriteString("\n")
	return buf
}
