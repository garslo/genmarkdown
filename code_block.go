package genmarkdown

import (
	"bytes"
	"io"
)

// CodeBlock represents a markdown code block. The Code is inserted
// as-is; the Language is optional.
type CodeBlock struct {
	Language string
	Code     string
}

// Markdown implements the Node interface.
func (me CodeBlock) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	buf.WriteString("```")
	buf.WriteString(me.Language + "\n")
	buf.WriteString(me.Code)
	buf.WriteString("\n```\n\n")
	return buf
}
