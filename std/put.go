package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Put(prop string) PutOntoPhrase {
	return PutOntoPhrase{prop: prop}
}

func (p PutOntoPhrase) Onto(supporter string) rt.Execute {
	// FIX: validate that the supporter is a supporter?
	return ChangeParent{core.Name(p.prop), "support", core.Name(supporter)}
}

type PutOntoPhrase struct {
	prop string
}
