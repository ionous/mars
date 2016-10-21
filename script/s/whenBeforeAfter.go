package s

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/frag"
	E "github.com/ionous/sashimi/event"
	G "github.com/ionous/sashimi/game"
	S "github.com/ionous/sashimi/source"
)

// statement to declare an default action handler
func To(action string, cb rt.Execute) frag.Fragment {
	return DefaultActionFragment{action, cb}
}

type DefaultActionFragment struct {
	action string
	cb     rt.Execute
}

func (to DefaultActionFragment) Build(src script.Source, top frag.Topic) error {
	fields := S.RunFields{top.Subject, to.action, to.cb, E.TargetPhase}
	return src.NewActionHandler(fields, script.Unknown)
}

//
// FIX: itd be nice to have some sort of wrapper to detect if they are used outside of,
// or rather not consumed by, the(). the wrapper would error at script compile time.

// a shortcut for meaning at the target
// ( implemented as a capturing event )
func Before(event string) EventPhrase {
	return EventPhrase{[]string{event}, S.ListenTargetOnly | S.ListenCapture}
}

// a shortcut for meaning at the target
// ( queues the callback to run after the default actions have completed. )
func After(event string) EventPhrase {
	return EventPhrase{[]string{event}, S.ListenTargetOnly | S.ListenCapture | S.ListenRunAfter}
}

// a shortcut for meaning at the target
// ( implemented as a bubbling event )
func When(event string) EventPhrase {
	return EventPhrase{[]string{event}, S.ListenTargetOnly}
}

//
func (p EventPhrase) Or(event string) EventPhrase {
	p.events = append(p.events, event)
	return p
}

//
func (p EventPhrase) Always(cb rt.Execute) EventFinalizer {
	return EventFinalizer{p, cb}
}

func (p EventPhrase) Go(cb G.Callback) EventFinalizer {
	return EventFinalizer{p, cb}
}

//
func (ef EventFinalizer) Build(src script.Source, top frag.Topic) (err error) {
	for _, evt := range ef.events {
		fields := S.ListenFields{top.Subject, evt, ef.cb, ef.options}
		if e := src.NewEventHandler(fields, script.Unknown); e != nil {
			err = e
			break
		}
	}
	return err
}

//
type EventPhrase struct {
	events  []string // name of the event in question
	options S.ListenOptions
}

//
type EventFinalizer struct {
	EventPhrase
	cb G.Callback
}
