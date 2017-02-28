package blocks

import (
	"github.com/ionous/sashimi/util/errutil"
)

type DocumentCursor struct {
	doc  Document
	curr *DocNode
}

func NewDocument() *DocumentCursor {
	return &DocumentCursor{doc: make(Document)}
}

func (dc *DocumentCursor) Document() Document {
	return dc.doc
}

func (dc *DocumentCursor) Top() *DocNode {
	return dc.curr
}

func (dc *DocumentCursor) Push(n *DocNode) error {
	dc.curr = dc.doc.AddElement(dc.curr, n)
	return nil
}

func (dc *DocumentCursor) Pop() (ret *DocNode, err error) {
	if n := dc.curr; n == nil {
		err = errutil.New("stack underflow", n.Path)
	} else {
		ret, dc.curr = n, n.Parent
	}
	return
}
