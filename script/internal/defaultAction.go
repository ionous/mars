package internal

import (
	"github.com/ionous/mars/rt"
	E "github.com/ionous/sashimi/event"
	S "github.com/ionous/sashimi/source"
)

type DefaultAction struct {
	Action string
	Call   rt.Execute
}

func NewDefaultAction(action string, call rt.Execute) Fragment {
	return DefaultAction{action, call}
}

func (to DefaultAction) GenFragment(src *S.Statements, top Topic) error {
	fields := S.RunFields{top.Subject, to.Action, to.Call, E.TargetPhase}
	return src.NewActionHandler(fields, S.UnknownLocation)
}
