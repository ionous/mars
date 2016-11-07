package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

func NewEvent(options S.ListenOptions, events ...string) EventPartial {
	return EventPartial{events, options}
}

// EventPartial expands to an EventPhrase. Partials are not serialized.
type EventPartial struct {
	Events  []string // name of the event in question
	Options S.ListenOptions
}

//
type EventPhrase struct {
	EventPartial
	Execute rt.Execute
}

// FIX: itd be nice to have some sort of wrapper to detect if they are used outside of,
// or rather not consumed by, the(). the wrapper would error at script compile time.
func (p EventPartial) Or(event string) EventPartial {
	p.Events = append(p.Events, event)
	return p
}

//
func (p EventPartial) Always(cb rt.Execute, cbs ...rt.Execute) EventPhrase {
	return EventPhrase{p, JoinCallbacks(cb, cbs)}
}

func (p EventPartial) Go(cb rt.Execute, cbs ...rt.Execute) EventPhrase {
	return EventPhrase{p, JoinCallbacks(cb, cbs)}
}

//
func (ef EventPhrase) GenFragment(src *S.Statements, top Topic) (err error) {
	for _, evt := range ef.Events {
		fields := S.ListenFields{top.Subject, evt, ef.Execute, ef.Options}
		if e := src.NewEventHandler(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}
