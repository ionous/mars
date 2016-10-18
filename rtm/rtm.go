package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"io"
)

type Rtm struct {
	model     meta.Model
	output    []*PrintMachine
	scope     []ScopeInfo
	callbacks Callbacks
}

// we really want more, source line, etc.
type Callbacks map[ident.Id]rt.Execute

type ScopeInfo struct {
	rt.Scope
	info *rt.IndexInfo
}

func NewRtm(data meta.Model, code Callbacks) *Rtm {
	return &Rtm{model: data, callbacks: code}
}

func (r *Rtm) Model() meta.Model {
	return r.model
}

func (r *Rtm) StopHere() {
	panic("not implemented")
}

// from NewRuntimeAction
func (r *Rtm) RunAction(act ident.Id, scope rt.Scope, parms rt.Parameters) (err error) {
	if act, ok := r.model.GetAction(act); !ok {
		err = errutil.New("unknown action", act)
	} else {
		types := act.GetNouns()
		switch diff := len(parms) + 1 - len(types); {
		case diff < 0:
			err = errutil.New("too few nouns specified for", act)
		case diff > 0:
			err = errutil.New("too many nouns specified for", act)
		default:
			if values, e := parms.Resolve(r); e != nil {
				err = e
			} else if cbs, ok := act.GetCallbacks(); ok {
				// FIX: how much of looping, etc. do you want to leak in?
				// maybe none; except for a very special "partials"?
				r.PushScope(ActionScope{r.model, types, values, scope}, nil)
				defer r.PopScope()

				for i := 0; i < cbs.NumCallback(); i++ {
					cb := cbs.CallbackNum(i)
					if cb, ok := r.callbacks[cb]; !ok {
						err = errutil.New("unknown action", cb)
						break
					}
				}
			}
		}
	}
	return
}

func (r *Rtm) LookupParent(inst meta.Instance) (meta.Instance, meta.Property, bool) {
	panic("not implemented")
	return nil, nil, false
}

func (r *Rtm) Print(args ...interface{}) (err error) {
	// get the top output, the one we want to write to
	out := r.output[len(r.output)-1]
	return out.Print(args...)
}

func (r *Rtm) Println(args ...interface{}) (err error) {
	out := r.output[len(r.output)-1]
	return out.Println(args...)
}

func (r *Rtm) PushOutput(out io.Writer) {
	r.output = append(r.output, &PrintMachine{flush: out})
}
func (r *Rtm) PopOutput() {
	r.output = r.output[:len(r.output)-1]
}
func (r *Rtm) Flush() error {
	out := r.output[len(r.output)-1]
	return out.Flush()
}

func (r *Rtm) PushScope(scope rt.Scope, info *rt.IndexInfo) {
	var cp *rt.IndexInfo
	if info != nil {
		idx := rt.IndexInfo(*info)
		cp = &idx
	}
	r.scope = append(r.scope, ScopeInfo{scope, cp})
}
func (r *Rtm) PopScope() {
	r.scope = r.scope[:len(r.scope)-1]
}

func (r *Rtm) GetScope() (scope rt.Scope, info *rt.IndexInfo) {
	if len(r.scope) > 0 {
		s := r.scope[len(r.scope)-1]
		scope, info = s.Scope, s.info
	} else {
		scope, info = xEmptyScope{}, nil
	}
	return
}

// xEmptyScope provides a default implementation for Scope
type xEmptyScope struct {
}

func (xEmptyScope) FindValue(string) (ret rt.Value, err error) {
	err = errutil.New("no scope is set")
	return
}
