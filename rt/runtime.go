package rt

import (
	"github.com/ionous/sashimi/meta"
	"io"
)

type Scope interface {
	// FindValue returns one of the mars eval statements.
	FindValue(string) (meta.Generic, error)
}

type IndexInfo struct {
	Index           int
	IsFirst, IsLast bool
}

// basically RuntimeCore replacing game.Play
// with event data and hint of GameEventAdapter
type Runtime interface {
	GetObject(Reference) (meta.Instance, error)
	RunAction(string, Scope, Parameters) error
	// Say: api.Output-> ScriptSays, ActorSays, Log
	PushOutput(io.Writer)
	Print(...interface{}) error
	Println(...interface{}) error
	PopOutput()
	Execute(meta.Callback) error
	// push -- this shouldnt be an object, it should be an interface
	// actions should, for instance, allow tripart objects
	PushScope(Scope, *IndexInfo)
	GetScope() (Scope, *IndexInfo)
	PopScope()
	// Query?
	// Random
	// api.LookupParents
	// LookupParent(inst meta.Instance) (ret meta.Instance, rel meta.Property, okay bool)
	// api.Event?
	StopHere()
}
