package rt

import (
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

type Runtime interface {
	meta.Model
	//
	FindValue(string) (meta.Generic, error)
	ScopePath() []string
	//
	Print(...interface{}) error
	Println(...interface{}) error

	// scope allows us to inject a new scope
	// MARS im not satisfied with this --  ex. what about a new print?
	// maybe revisit post lookupparent cahnges
	RunAction(action ident.Id, scope Scope, args ...meta.Generic) error

	// LookupParent is a nod to the stdlib: objects have a "parent" relation
	// but the parent mechanism is currently actually multiple properties.
	// MARS: once we can fully store machines in properties (and have class defaults),
	// the stdlib could store a parent eval; to satisfy its current implementation
	// ( multiple properties ) it would probably have to return a data object
	// designating the parent and the name of the relation.
	// Users/ionous/Dev/go/src/github.com/ionous/sashimi/play/api/lookupParents.go
	LookupParent(meta.Instance) (meta.Instance, meta.Property, bool)
}
