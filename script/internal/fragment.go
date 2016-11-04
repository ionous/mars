package internal

import S "github.com/ionous/sashimi/source"

// Topic targets a noun and or its type.
type Topic struct {
	Target, Subject string
}

// Fragment phrases appear in "The" phrases.
type Fragment interface {
	GenFragment(*S.Statements, Topic) error
}

// Fragments array
type Fragments []Fragment
