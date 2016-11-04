package internal

import (
	S "github.com/ionous/sashimi/source"
	"strings"
)

// FIX: hacky, maybe should be HaveMany() instead.
const ListKind = " list"

func NewClassProperty(n, k string) Fragment {
	return ClassProperty{n, k}
}

type ClassProperty struct {
	Name string // property field name
	Kind string // property kind: primitive or user class
}

func (c ClassProperty) GenFragment(src *S.Statements, top Topic) error {
	isMany, kind := c.listKind()
	fields := S.PropertyFields{top.Subject, c.Name, kind, isMany}
	return src.NewProperty(fields, S.UnknownLocation)
}

func (c ClassProperty) listKind() (isMany bool, kind string) {
	kind = c.Kind
	if i := strings.Index(kind, ListKind); i > 0 {
		if i+len(ListKind) == len(kind) {
			kind = kind[:i]
			isMany = true
		}
	}
	return
}
