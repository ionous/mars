package rt

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
	"io"
)

// basically RuntimeCore replacing game.Play
// with event data and hint of GameEventAdapter
type Runtime interface {
	RunAction(ident.Id, []meta.Generic) error
	// Say: api.Output-> ScriptSays, ActorSays, Log
	Print(...interface{}) error
	Println(...interface{}) error
	//
	GetObject(ident.Id) (Object, error)
	FindValue
	//
	PushOutput(io.Writer)
	PopOutput()
	// Query?
	// Random
	// api.LookupParents
	// LookupParent(inst meta.Instance) (ret meta.Instance, rel meta.Property, okay bool)
	// api.Event?
}
