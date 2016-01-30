package genmarkdown

import (
	"bytes"
	"io"
)

type CodeBlock struct {
	Language string
	Code     string
}

func (me CodeBlock) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	buf.WriteString("```")
	buf.WriteString(me.Language + "\n")
	buf.WriteString(me.Code)
	buf.WriteString("\n```\n\n")
	return buf
}
