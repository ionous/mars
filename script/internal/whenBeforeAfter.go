package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type EventPhrase struct {
	Events types.NamedEvents `mars:"the [events]"`
	Calls  []rt.Execute      `mars:"always"`
}

type BeforeEvent struct {
	EventPhrase
}
type AfterEvent struct {
	EventPhrase
}
type WhenEvent struct {
	EventPhrase
}

func (ev *BeforeEvent) GenFragment(src *S.Statements, top Topic) error {
	return ev.GenEvents(src, top, S.ListenTargetOnly|S.ListenCapture)
}

func (ev *AfterEvent) GenFragment(src *S.Statements, top Topic) error {
	return ev.GenEvents(src, top, S.ListenTargetOnly|S.ListenCapture|S.ListenRunAfter)
}

func (ev *WhenEvent) GenFragment(src *S.Statements, top Topic) error {
	return ev.GenEvents(src, top, S.ListenTargetOnly)
}

type EventPartial struct {
	fragment Fragment
	data     *EventPhrase
}

func NewEvent(evt types.NamedEvent, f Fragment, p *EventPhrase) EventPartial {
	p.Events = types.NamedEvents{evt.String()}
	return EventPartial{f, p}
}

func (p EventPartial) Or(event types.NamedEvent) EventPartial {
	p.data.Events = append(p.data.Events, event.String())
	return p
}

func (p EventPartial) Always(cb rt.Execute, cbs ...rt.Execute) Fragment {
	p.data.Calls = JoinCallbacks(cb, cbs)
	return p.fragment
}

func (p *EventPhrase) GenEvents(src *S.Statements, top Topic, opt S.ListenOptions) (err error) {
	for _, evt := range p.Events {
		fields := S.ListenFields{top.Subject.String(), evt, p.Calls, opt}
		if e := src.NewEventHandler(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}
