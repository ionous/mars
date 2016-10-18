package std

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Move(obj string) MoveToPhrase {
	return MoveToPhrase{obj: obj}
}

func (move MoveToPhrase) To(dest string) rt.Execute {
	return ChangeParent{core.R(move.obj), "whereabouts", core.R(dest)}
}

func (move MoveToPhrase) OutOfWorld() rt.Execute {
	return ChangeParent{core.R(move.obj), "whereabouts", core.NullRef()}
}

type MoveToPhrase struct {
	obj string
}
