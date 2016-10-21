package frag

import (
	"github.com/ionous/mars/script"
)

type Topic struct {
	Target, Subject string
}

type Fragment interface {
	Build(script.Source, Topic) error
}

type Fragments []Fragment
