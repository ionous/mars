package frag

import "github.com/ionous/mars/script"

type Exists struct {
}

func (Exists) Build(script.Source, Topic) error {
	return nil
}
