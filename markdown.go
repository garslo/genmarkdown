package genmarkdown

import (
	"bytes"
	"io"
)

// Node is the basic interface that all markdown types must
// implement. The returned io.Reader must provide all text that the
// type is composed of.
type Node interface {
	Markdown() io.Reader
}

type Nodes []Node

// Markdown implements the Node interface.
func (me Nodes) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	for _, node := range me {
		io.Copy(buf, node.Markdown())
	}
	return buf
}

// Add appends a node in-place.
func (me *Nodes) Add(n Node) {
	*me = append(*me, n)
}
