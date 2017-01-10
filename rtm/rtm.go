package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type Rtm struct {
	model  meta.Model
	output OutputStack
	scope  ScopeStack
	// should use the stack for this
	lineWait bool
}

func (rtm *Rtm) Runtime() rt.Runtime {
	return localRuntime{rtm.model, rtm}
}

func NewRtm(model meta.Model) *Rtm {
	rtm := &Rtm{model: model}
	rtm.scope.PushScope(scope.NewModelScope(model))
	return rtm
}

func (rtm *Rtm) Flush() error {
	return rtm.output.Flush()
}

func (rtm *Rtm) RunAction(id ident.Id, scp rt.Scope, args ...meta.Generic) (err error) {
	if isArray(args) {
		err = errutil.New("RunAction parameters passed as array")
	} else if act, e := NewActionRuntime(rtm.Runtime(), id, scp, args); e != nil {
		err = e
	} else {
		err = act.RunDefault()
	}
	return
}

func isArray(args []meta.Generic) (ret bool) {
	if paramcnt := len(args); paramcnt == 1 {
		if _, isArray := args[0].([]meta.Generic); isArray {
			ret = isArray
		}
	}
	return
}
