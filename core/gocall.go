package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type GoCall struct {
	Action     types.NamedAction
	Parameters []meta.Generic
}

func (gc GoCall) Execute(run rt.Runtime) (err error) {
	id := ident.MakeId(gc.Action.String())
	if e := run.RunAction(id, run, gc.Parameters...); e != nil {
		err = errutil.New("GoCall failed", gc.Action, e)
	}
	return
}
