package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

type GoCall struct {
	Action     ident.Id
	Parameters []meta.Generic
}

func (gc GoCall) Execute(run rt.Runtime) error {
	return run.RunAction(gc.Action, gc.Parameters)
}
