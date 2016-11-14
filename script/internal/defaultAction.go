package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	E "github.com/ionous/sashimi/event"
	S "github.com/ionous/sashimi/source"
)

type DefaultAction struct {
	Action string
	Calls  []rt.Execute
}

func NewDefaultAction(action string, calls []rt.Execute) Fragment {
	return DefaultAction{action, calls}
}

func (to DefaultAction) GenFragment(src *S.Statements, top Topic) error {
	fields := S.RunFields{string(top.Subject), to.Action, to.Calls, E.TargetPhase}
	return src.NewActionHandler(fields, S.UnknownLocation)
}
