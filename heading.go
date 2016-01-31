package genmarkdown

import (
	"bytes"
	"io"
)

// Heading represents a markdown header. The Level indicates the
// number of leading '#' characters.
type Heading struct {
	Text  string
	Level int
}

// Markdown implements the Node interface.
func (me Heading) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	for i := 0; i < me.Level; i++ {
		buf.Write([]byte("#"))
	}
	buf.WriteString(" " + me.Text + "\n\n")
	return buf
}
