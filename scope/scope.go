package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
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

func (sr scopedRuntime) FindValue(s string) (ret meta.Generic, err error) {
	if r, e := sr.find.FindValue(s); e != nil {
		err = errutil.New("find value error", e, sr.find.ScopePath())
	} else {
		ret = r
	}
	return
}
