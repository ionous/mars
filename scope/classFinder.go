package scope

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/lang"
	"github.com/ionous/sashimi/util/sbuf"
)

type ClassFinder interface {
	FindClass(ident.Id) (meta.Generic, error)
}

type ClassScope struct {
	meta.Model
	ClassFinder
}

func (cs *ClassScope) FindValue(name string) (ret meta.Generic, err error) {
	clsid := ident.MakeId(cs.Model.Pluralize(lang.StripArticle(name)))
	if v, e := cs.FindClass(clsid); e == nil {
		ret = v
	} else if c, ok := e.(ClassNotFound); !ok {
		err = e
	} else {
		err = NotFound(cs, string(c))
	}
	return
}

func (cs *ClassScope) ScopePath() (parts []string) {
	return append(parts, sbuf.Type{cs.ClassFinder}.String())
}
