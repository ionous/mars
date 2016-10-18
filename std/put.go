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
	return ChangeParent{core.R(p.prop), "support", core.R(supporter)}
}

type PutOntoPhrase struct {
	prop string
}
