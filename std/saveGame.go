package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/scope"
	"github.com/ionous/sashimi/play"
)

type SaveGame struct {
	AutoSave      rt.BoolEval
	Saved, Failed rt.Statements
}

func (op SaveGame) Execute(run rt.Runtime) (err error) {
	if autosave, e := op.AutoSave.GetBool(run); e != nil {
		err = e
	} else if msg, ok := play.SaveCurrentGame(run, autosave.Value); true {
		run := scope.Make(run, scope.NewValue(T(msg)), run)
		if ok {
			err = op.Saved.ExecuteList(run)
		} else {
			err = op.Failed.ExecuteList(run)
		}
	}
	return
}
