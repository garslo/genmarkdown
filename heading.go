package genmarkdown

import (
	"bytes"
	"io"
)

type Heading struct {
	Text  string
	Level int
}

func (me Heading) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	for i := 0; i < me.Level; i++ {
		buf.Write([]byte("#"))
	}
	buf.WriteString(" " + me.Text + "\n\n")
	return buf
}
