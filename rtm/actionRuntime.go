package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/sbuf"
	"reflect"
)

func NewActionRuntime(run rt.Runtime, id ident.Id, scp rt.Scope, params []meta.Generic) (ret *ActionRuntime, err error) {
	if act, ok := run.GetAction(id); !ok {
		err = errutil.New("NewActionRuntime unknown action", id)
	} else {
		types := act.GetNouns()
		switch diff := len(params) - len(types); {
		case diff < 0:
			err = errutil.New("NewActionRuntime too few nouns", len(params), "of", len(types), "specified for", act)
		case diff > 0:
			err = errutil.New("NewActionRuntime too many nouns", len(params), "of", len(types), "specified for", act)
		default:
			run := scope.Make(run, scp)
			// zip: resolve parameters into primitive values
			// noting that parameters andprimitives are both stored as evalues :\oP
			values := make([]meta.Generic, len(params))
			// in the future, we want to allow any type of value.
			// its here we would new our action temporary data and "zip" the parameters into it
			// verifying as we go --
			for i, p := range params {
				if eval, ok := p.(rt.ObjEval); !ok {
					err = errutil.New("NewActionRuntime only objects are supported", i, sbuf.Type{p})
					break
				} else if obj, e := eval.GetObject(run); e != nil {
					err = errutil.New("NewActionRuntime failed to find object", i, e)
					break
				} else {
					want, have := types[i], obj.GetParentClass()
					if !run.AreCompatible(have, want) {
						err = errutil.New("NewActionRuntime type mismatch", obj, i, "is", have, ", expected", want)
						break
					} else {
						values[i] = obj
					}
				}
			}
			if err == nil {
				chain := scope.NewListener(run, types, values)
				ret = &ActionRuntime{run, act, chain, values, nil}
			}
		}
	}
	return
}

type ActionRuntime struct {
	rt.Runtime
	meta.Action
	scope  *scope.ListenerScope
	values []meta.Generic
	after  []QueuedCallback
}

type QueuedCallback struct {
	cb   meta.Callback
	hint ident.Id
}

func (ap *ActionRuntime) FindValue(name string) (ret meta.Generic, err error) {
	return ap.scope.FindValue(name)
}

func (ap *ActionRuntime) Values() []meta.Generic {
	return ap.values
}

// GetTarget can return nil
func (ap *ActionRuntime) GetTarget() (ret meta.Instance) {
	return ap.getObject(0)
}

// GetContext can return nil
func (ap *ActionRuntime) GetContext() (ret meta.Instance) {
	return ap.getObject(1)
}

func (ap *ActionRuntime) getObject(i int) (ret meta.Instance) {
	if cnt := len(ap.values); i < cnt {
		v := ap.values[i]
		if eval, ok := v.(rt.ObjEval); ok {
			/// what to do with error!?
			if obj, e := eval.GetObject(ap); e == nil {
				ret = obj.Instance
			}
		}
	}
	return
}

func (ap *ActionRuntime) run(cb meta.Callback) (err error) {
	if calls, ok := cb.([]rt.Execute); !ok {
		err = errutil.New("ActionRuntime", ap.GetId(), "callback not of execute type", reflect.TypeOf(cb))
	} else {
		for _, c := range calls {
			if e := c.Execute(ap); e != nil {
				err = e
				break
			}
		}
	}
	return
}

func (ap *ActionRuntime) RunNow(cb meta.Callback, hint ident.Id) (err error) {
	ap.scope.SetHint(hint)
	if e := ap.run(cb); e != nil {
		err = errutil.New("RunNow", e)
	}
	return err
}

func (ap *ActionRuntime) RunLater(cb meta.Callback, hint ident.Id) (err error) {
	ap.after = append(ap.after, QueuedCallback{cb, hint})
	return
}

// MARS: perhaps meta.Generic could be an interface Resolve() which returns interface{}
// then we'd have some type-safety at least
func (ap *ActionRuntime) RunDefault() (err error) {
	if cbs, ok := ap.GetCallbacks(); ok {
		ap.scope.SetHint(ident.Empty())
		for i := 0; i < cbs.NumCallback(); i++ {
			if e := ap.run(cbs.CallbackNum(i)); e != nil {
				err = errutil.New("RunDefault", i, e)
				break
			}
		}
	}
	return
}

// run "after" actions queued by RunCallbackLater
func (ap *ActionRuntime) RunAfterActions() (err error) {
	if after := ap.after; len(after) > 0 {
		for i, qa := range after {
			cb, hint := qa.cb, qa.hint
			ap.scope.SetHint(hint)
			if e := ap.run(cb); e != nil {
				err = errutil.New("RunAfterActions", i, e)
				break
			}
		}
	}
	return
}
