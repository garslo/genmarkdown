package genmarkdown

import (
	"bytes"
	"io"
	"os"
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

type Emitter struct {
	Nodes Nodes
}

func (me Emitter) Emit() error {
	io.Copy(os.Stdout, me.Nodes.Markdown())
	return nil
}
