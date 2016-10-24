package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
	"github.com/ionous/sashimi/util/sbuf"
	"io"
)

type Rtm struct {
	model meta.Model
	scope.ModelScope
	output []*PrintMachine
}

func NewRtm(model meta.Model) *Rtm {
	ms := scope.NewModelScope(model)
	return &Rtm{ModelScope: ms, model: model}
}

func (run *Rtm) RunAction(id ident.Id, params []meta.Generic) (err error) {
	if act, e := run.GetAction(id, params); e != nil {
		err = e
	} else {
		err = act.RunDefault()
	}
	return
}

func (run *Rtm) GetAction(id ident.Id, params []meta.Generic) (ret ActionProvider, err error) {
	if act, ok := run.model.GetAction(id); !ok {
		err = errutil.New("RunAction: unknown action", id)
	} else {
		types := act.GetNouns()
		switch diff := len(params) - len(types); {
		case diff < 0:
			err = errutil.New("RunAction: too few nouns specified for", act)
		case diff > 0:
			err = errutil.New("RunAction: too many nouns specified for", act)
		default:
			// zip:
			// in the future, we want to allow any type of value.
			// its in here we would new our action temporary data and "zip" the parameters into it
			// verifying as we go --
			vals := make([]meta.Generic, len(params))
			for i, p := range params {
				if eval, ok := p.(rt.ObjEval); !ok {
					err = errutil.New("RunAction: only objects are supported", eval, i, sbuf.Type{p})
					break
				} else if obj, e := eval.GetObject(run); e != nil {
					err = e
					break
				} else {
					want, have := types[i], obj.GetParentClass()
					if !run.model.AreCompatible(have, want) {
						err = errutil.New("RunAction: type mismatch", obj, i, "is", have, ", expected", want)
						break
					} else {
						vals[i] = obj
					}
				}
			}
			if err == nil {
				ret = MakeProvider(run, act, vals)
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

func (run *Rtm) GetObject(id ident.Id) (ret rt.Object, err error) {
	if id.Empty() {
		err = errutil.New("rtm.GetObject(id) is nil")
	} else if inst, ok := run.model.GetInstance(id); !ok {
		err = errutil.New("rtm.GetObject(id) not found", id)
	} else {
		ret = rt.Object{inst}
	}
	return
}
