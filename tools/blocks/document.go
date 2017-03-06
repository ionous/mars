package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

//
type DocNode struct {
	Parent   *DocNode             `json:"-"`
	Children []*DocNode           `json:",omitempty"`
	Path     inspect.Path         `json:",omitempty"`
	BaseType *inspect.CommandInfo `json:",omitempty"`
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

type DocNodes map[string]*DocNode

// Document: to support rebuilding the model from arbitrary points
// we need a pool of records for every path;
// to support match lookups we need a stack.
type Document DocNodes

// adds the node to the document and links up children.
func (d Document) AddElement(p, n *DocNode) *DocNode {
	if p != nil {
		n.Parent, p.Children = p, append(p.Children, n)
	}
	d[n.Path.String()] = n
	return n
}
