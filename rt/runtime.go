package rt

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
	"io"
)

type Execute interface {
	Execute(Runtime) error
}

type ParameterSource interface {
	// Push should call push target.
	Resolve(Runtime) (Value, error)
}

type Parameters []ParameterSource

func (ps Parameters) Resolve(r Runtime) (ret []Value, err error) {
	for _, p := range ps {
		if v, e := p.Resolve(r); e != nil {
			err = e
			break
		} else {
			ret = append(ret, v)
		}
	}
	return
}

// basically RuntimeCore replacing game.Play
// with event data and hint of GameEventAdapter
type Runtime interface {
	// The, R, A
	Model() meta.Model
	RunAction(ident.Id, Scope, Parameters) error
	// Say: api.Output-> ScriptSays, ActorSays, Log
	PushOutput(io.Writer)
	Print(...interface{}) error
	Println(...interface{}) error
	PopOutput()
	// push -- this shouldnt be an object, it should be an interface
	// actions should, for instance, allow tripart objects
	PushScope(Scope, *IndexInfo)
	GetScope() (Scope, *IndexInfo)
	PopScope()
	//

	// Query?
	// Random
	// api.LookupParents
	LookupParent(inst meta.Instance) (ret meta.Instance, rel meta.Property, okay bool)
	// api.Event?
	StopHere()
}
