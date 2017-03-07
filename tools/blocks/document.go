package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

//
type DocNode struct {
	Parent   *DocNode             `json:"-"`
	Path     inspect.Path         `json:",omitempty"`
	Children []*DocNode           `json:",omitempty"`
	Slot     *inspect.CommandInfo `json:",omitempty"`
	Command  *inspect.CommandInfo `json:",omitempty"`
	Param    *inspect.ParamInfo   `json:",omitempty"`
	Data     interface{}          `json:",omitempty"` // or... text?
}

func (n *DocNode) NumChildren() int {
	return len(n.Children)
}

func (n *DocNode) MaxChildren() int {
	return cap(n.Children)
}

// Document: to support rebuilding the model from arbitrary points
// we need a pool of records for every path;
// to support match lookups we need a stack.
type Document map[string]*DocNode

// adds the node to the document and links up children.
func (d Document) AddElement(p, n *DocNode) *DocNode {
	if p != nil {
		if p == n {
			panic("cant add node to itself")
		}
		n.Parent, p.Children = p, append(p.Children, n)
		if parentPath := n.Path.ParentPath(); inspect.PathCompare(p.Path, parentPath) != 0 {
			panic("mismatched paths: parent=" + p.Path.String() + "; child=" + n.Path.String())
		}
	}
	d[n.Path.String()] = n
	return n
}
