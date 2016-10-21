package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/sbuf"
	"io"
	"reflect"
)

type Rtm struct {
	model  meta.Model
	output []*PrintMachine
	scope  []ScopeInfo
}

type ScopeInfo struct {
	rt.Scope
	info *rt.IndexInfo
}

func NewRtm(data meta.Model) *Rtm {
	return &Rtm{model: data}
}

func (run *Rtm) StopHere() {
	panic("not implemented")
}

func (run *Rtm) Execute(cb meta.Callback) (err error) {
	if exec, ok := cb.(rt.Execute); !ok {
		err = errutil.New("callback not of execute type", sbuf.Type{cb})
	} else {
		err = exec.Execute(run)
	}
	return
}

// from NewRuntimeAction
func (run *Rtm) RunAction(act string, scope rt.Scope, parms rt.Parameters) (err error) {
	if act, ok := run.model.GetAction(ident.MakeId(act)); !ok {
		err = errutil.New("unknown action", act)
	} else {
		types := act.GetNouns()
		switch diff := len(parms) + 1 - len(types); {
		case diff < 0:
			err = errutil.New("too few nouns specified for", act)
		case diff > 0:
			err = errutil.New("too many nouns specified for", act)
		default:
			if values, e := parms.Resolve(run); e != nil {
				err = e
			} else if cbs, ok := act.GetCallbacks(); ok {
				// FIX: how much of looping, etc. do you want to leak in?
				// maybe none; except for a very special "partials"?
				run.PushScope(ActionScope{run.model, types, values, scope}, nil)
				defer run.PopScope()

				for i := 0; i < cbs.NumCallback(); i++ {
					cb := cbs.CallbackNum(i)
					if exec, ok := cb.(rt.Execute); !ok {
						err = errutil.New("callback not of execute type", reflect.TypeOf(cb))
						break
					} else if e := exec.Execute(run); e != nil {
						err = e
						break
					}
				}
			}
		}
	}
	return
}

func (run *Rtm) LookupParent(inst meta.Instance) (meta.Instance, meta.Property, bool) {
	panic("not implemented")
	return nil, nil, false
}

func (run *Rtm) Print(args ...interface{}) (err error) {
	// get the top output, the one we want to write to
	if cnt := len(run.output); cnt > 0 {
		out := run.output[len(run.output)-1]
		err = out.Print(args...)
	} else {
		err = errutil.New("runtime lacks an output stream")
	}
	return
}

func (run *Rtm) Println(args ...interface{}) (err error) {
	out := run.output[len(run.output)-1]
	return out.Println(args...)
}

func (run *Rtm) PushOutput(out io.Writer) {
	run.output = append(run.output, &PrintMachine{flush: out})
}
func (run *Rtm) PopOutput() {
	run.output = run.output[:len(run.output)-1]
}
func (run *Rtm) Flush() error {
	out := run.output[len(run.output)-1]
	return out.Flush()
}

func (run *Rtm) PushScope(scope rt.Scope, info *rt.IndexInfo) {
	var cp *rt.IndexInfo
	if info != nil {
		idx := rt.IndexInfo(*info)
		cp = &idx
	}
	run.scope = append(run.scope, ScopeInfo{scope, cp})
}
func (run *Rtm) PopScope() {
	run.scope = run.scope[:len(run.scope)-1]
}

func (run *Rtm) GetScope() (scope rt.Scope, info *rt.IndexInfo) {
	if len(run.scope) > 0 {
		s := run.scope[len(run.scope)-1]
		scope, info = s.Scope, s.info
	} else {
		scope, info = xEmptyScope{}, nil
	}
	return
}

// FIX: maybe runtime will return this? ie. rt.GetObject(ref) instead of exposing the model
func (run *Rtm) GetObject(xr rt.Reference) (ret meta.Instance, err error) {
	if inst, ok := run.model.GetInstance(xr.Id()); !ok {
		err = errutil.New("instance not found", xr.Id)
	} else {
		ret = inst
	}
	return
}

// xEmptyScope provides a default implementation for Scope
type xEmptyScope struct {
}

func (xEmptyScope) FindValue(string) (ret meta.Generic, err error) {
	err = errutil.New("no scope is set")
	return
}
