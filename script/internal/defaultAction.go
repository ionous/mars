package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	E "github.com/ionous/sashimi/event"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type DefaultAction struct {
	Action types.NamedAction `mars:"to [act]"`
	Calls  []rt.Execute
}

func (to DefaultAction) GenFragment(src *S.Statements, top Topic) error {
	fields := S.RunFields{top.Subject, to.Action.String(), to.Calls, E.TargetPhase}
	return src.NewActionHandler(fields, S.UnknownLocation)
}
