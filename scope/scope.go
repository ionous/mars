package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

func Make(run rt.Runtime, find ...rt.Scope) rt.Runtime {
	return scopedRuntime{
		run,
		find,
	}
}

type scopedRuntime struct {
	rt.Runtime
	find ScopeChain
}

func (sr scopedRuntime) FindValue(s string) (meta.Generic, error) {
	return sr.find.FindValue(s)
}
