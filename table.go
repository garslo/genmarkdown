package genmarkdown

import (
	"bytes"
	"fmt"
	"io"
	"text/tabwriter"
)

type Row struct {
	Values []string
}

type Table struct {
	Columns []string
	Rows    []Row
}

// Markdown implements the Node interface. It formats the table in a
// human-friendly manner; all rows are aligned and padded properly.
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
