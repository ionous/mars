package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"reflect"
)

func MakeProvider(run *Rtm, act meta.Action, values []meta.Generic) ActionProvider {
	nouns := act.GetNouns()
	chain := scope.NewChain(
		run.ModelScope,
		scope.NewParamScope(values),
		scope.ExactClass(run.model, nouns, values),
		scope.SimilarClass(run.model, nouns, values))
	return ActionProvider{run, act, values, chain}
}

type ActionProvider struct {
	rt.Runtime
	meta.Action
	values []meta.Generic
	chain  rt.FindValue
}

// GetTarget can return nil
func (ap ActionProvider) GetTarget() (ret meta.Instance) {
	return ap.getObject(0)
}

// GetContext can return nil
func (ap ActionProvider) GetContext() (ret meta.Instance) {
	return ap.getObject(1)
}

func (ap ActionProvider) getObject(i int) (ret meta.Instance) {
	if cnt := len(ap.values); i < cnt {
		v := ap.values[i]
		if eval, ok := v.(rt.ObjEval); ok {
			if obj, e := eval.GetObject(ap.Runtime); e == nil {
				ret = obj.Instance
			}
		}
	}
	return
}

func (ap ActionProvider) FindValue(name string) (meta.Generic, error) {
	return ap.chain.FindValue(name)
}
func (ap ActionProvider) ScopePath() []string {
	parts := ap.chain.ScopePath()
	return append(parts, "ActionProvider", ap.GetId().String())
}

// MARS: perhaps meta.Generic could be an interface Resolve() which returns interface{}
// then we'd have some type-safety at least
func (ap ActionProvider) RunDefault() (err error) {
	if cbs, ok := ap.GetCallbacks(); ok {
		for i := 0; i < cbs.NumCallback(); i++ {
			cb := cbs.CallbackNum(i)
			if exec, ok := cb.(rt.Execute); !ok {
				err = errutil.New("RunAction: callback not of execute type", reflect.TypeOf(cb))
				break
			} else if e := exec.Execute(ap); e != nil {
				err = e
				break
			}
		}
	}
	return
}
