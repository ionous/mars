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
	// LookupParent is a nod to the stdlib: objects have a "parent" relation
	// but the parent mechanism is currently actually multiple properties.
	// MARS: once we can fully store machines in properties (and have class defaults),
	// the stdlib could store a parent eval; to satisfy its current implementation
	// ( multiple properties ) it would probably have to return a data object
	// designating the parent and the name of the relation.
	LookupParent(meta.Instance) (meta.Instance, meta.Property, bool)
}
