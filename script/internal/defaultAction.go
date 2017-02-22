package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	E "github.com/ionous/sashimi/event"
	S "github.com/ionous/sashimi/source"
)

type DefaultAction struct {
	Action string `mars:"to [act];action"`
	Calls  []rt.Execute
}

func (to DefaultAction) GenFragment(src *S.Statements, top Topic) error {
	fields := S.RunFields{top.Subject, to.Action, to.Calls, E.TargetPhase}
	return src.NewActionHandler(fields, S.UnknownLocation)
}
