package backend

import (
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

// Topic targets a noun and or its type.
type Topic struct {
	Target  string
	Subject types.NamedSubject
}

// Fragment phrases appear in "The" phrases.
type Fragment interface {
	GenFragment(*S.Statements, Topic) error
}

// Fragments array
type Fragments []Fragment
