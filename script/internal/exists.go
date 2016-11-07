package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type Exists struct{}

func (Exists) GenFragment(*S.Statements, Topic) error {
	return nil
}
