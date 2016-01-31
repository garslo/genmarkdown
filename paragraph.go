package genmarkdown

import (
	"bytes"
	"io"
	"strings"
)

// Paragraph represents a markdown paragraph. All lines in the
// rendered paragraph will have length <= MaxLineLength; each
// paragraph will have an empty line at the end.
type Paragraph struct {
	Text string
}

// Maximum length of each paragraph line.
const MaxLineLength = 80

// Markdown implements the Node interface.
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
