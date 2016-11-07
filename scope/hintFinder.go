package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

type HintFinder struct {
	hp     *ident.Id
	values ValueFinder
}

func ClassHint(model meta.Model, hp *ident.Id, values []meta.Generic) (ret rt.Scope) {
	return &ClassScope{model, &HintFinder{hp, values}}
}

func (hf *HintFinder) FindClass(id ident.Id) (ret meta.Generic, err error) {
	if string(id) != string(*hf.hp) {
		err = ClassNotFound(string(id))
	} else if inst, e := hf.values.getValue(0); e != nil {
		err = e
	} else {
		ret = inst
	}
	return
}
