package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/g"
)

func Move(obj string) MoveToPhrase {
	return MoveToPhrase{g.The(obj)}
}

func (move MoveToPhrase) To(dest string) rt.Execute {
	return AssignTo(move.obj, "whereabouts", g.The(dest))
}

func (move MoveToPhrase) OutOfWorld() rt.Execute {
	return AssignTo(move.obj, "whereabouts", core.NullRef())
}

type MoveToPhrase struct {
	obj rt.ObjEval
}
