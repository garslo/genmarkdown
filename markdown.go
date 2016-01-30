package genmarkdown

import (
	"bytes"
	"io"
)

type Node interface {
	Markdown() io.Reader
}

type Nodes []Node

func (me Nodes) Markdown() io.Reader {
	buf := &bytes.Buffer{}
	for _, node := range me {
		io.Copy(buf, node.Markdown())
	}
	return buf
}

func (me *Nodes) Add(n Node) {
	*me = append(*me, n)
}
