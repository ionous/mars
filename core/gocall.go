package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/ident"
)

type GoCall struct {
	Action     ident.Id
	Parameters []meta.Generic
}

func (gc GoCall) Execute(run rt.Runtime) (err error) {
	if e := run.RunAction(gc.Action, run, gc.Parameters...); e != nil {
		err = errutil.New("GoCall failed", gc.Action, e)
	}
	return
}
