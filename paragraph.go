package genmarkdown

import (
	"bytes"
	"io"
	"strings"
)

type Paragraph struct {
	Text string
}

const MaxLineLength = 80

// Markdown implements the Node interface. It ensures that all lines
// are <= 80 characters in length and that each paragraph will have an
// empty line at the end
func (me Paragraph) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	chunked := strings.Fields(me.Text)
	thisLineLen := 0
	startOfLine := true

	for _, chunk := range chunked {
		chunkLen := len(chunk)
		if thisLineLen+chunkLen > MaxLineLength {
			buf.WriteString("\n")
			thisLineLen = 0
			startOfLine = true
		}
		if !startOfLine {
			buf.WriteString(" ")
		}
		startOfLine = false
		buf.WriteString(chunk)
		thisLineLen += chunkLen + 1
	}
	buf.WriteString("\n\n")
	return buf
}
