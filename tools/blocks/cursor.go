package blocks

import (
	"github.com/ionous/sashimi/util/errutil"
)

type Cursor struct {
	doc        Document
	root, curr *DocNode
}

func NewDocument() *Cursor {
	root := &DocNode{}
	return &Cursor{make(Document), root, root}
}

func (dc *Cursor) Document() Document {
	return dc.doc
}

func (dc *Cursor) Root() *DocNode {
	return dc.root
}

func (dc *Cursor) Top() *DocNode {
	return dc.curr
}

func (dc *Cursor) Push(n *DocNode) error {
	dc.curr = dc.doc.AddElement(dc.curr, n)
	return nil
}

func (dc *Cursor) Pop() (ret *DocNode, err error) {
	if n := dc.curr; n == nil {
		err = errutil.New("stack underflow", n.Path)
	} else {
		ret, dc.curr = n, n.Parent
	}
	return
}

func (dc *Cursor) Flush() (err error) {
	for dc.Top() != nil {
		if _, e := dc.Pop(); e != nil {
			err = e
			break
		}
	}
	return
}
