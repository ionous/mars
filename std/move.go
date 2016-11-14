package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Move(what rt.ObjEval) MoveToPhrase {
	return MoveToPhrase{what}
}

func (move MoveToPhrase) To(where rt.ObjEval) rt.Execute {
	return AssignParent(move.what, "whereabouts", where)
}

func (move MoveToPhrase) OutOfWorld() rt.Execute {
	return AssignParent(move.what, "whereabouts", core.Nothing())
}

type MoveToPhrase struct {
	what rt.ObjEval
}
