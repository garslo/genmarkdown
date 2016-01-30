package genmarkdown

import (
	"bytes"
	"io"
)

type Paragraph struct {
	Text string
}

func (me Paragraph) Markdown() io.Reader {
	return bytes.NewBuffer([]byte(me.Text + "\n\n"))
}
