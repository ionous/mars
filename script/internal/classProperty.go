package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
	"strings"
)

// FIX: hacky, maybe should be HaveMany() instead.
const ListKind = " list"

type ClassProperty struct {
	Name types.NamedProperty // property field name
	Kind types.NamedClass    // property kind: primitive or user class
}

func (c ClassProperty) GenFragment(src *S.Statements, top Topic) error {
	isMany, kind := c.listKind()
	fields := S.PropertyFields{top.Subject.String(), c.Name.String(), kind, isMany}
	return src.NewProperty(fields, S.UnknownLocation)
}

func (c ClassProperty) listKind() (isMany bool, kind string) {
	kind = c.Kind.String()
	if i := strings.Index(kind, ListKind); i > 0 {
		if i+len(ListKind) == len(kind) {
			kind = kind[:i]
			isMany = true
		}
	}
	return
}
