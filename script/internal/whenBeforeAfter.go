package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type EventTiming interface {
	GetOptions() S.ListenOptions
}

type BeforeEvent struct{}

type AfterEvent struct{}

type WhenEvent struct{}

type HandleEvent struct {
	Run    EventTiming  `mars:"[handle event]"`
	Events []string     `mars:"[event name(s)]"`
	Calls  []rt.Execute `mars:"always: [run actions]."`
}

func (_ BeforeEvent) GetOptions() S.ListenOptions {
	return S.ListenTargetOnly | S.ListenCapture
}

func (_ AfterEvent) GetOptions() S.ListenOptions {
	return S.ListenTargetOnly | S.ListenCapture | S.ListenRunAfter
}

func (_ WhenEvent) GetOptions() S.ListenOptions {
	return S.ListenTargetOnly
}

type EventPartial struct {
	data *HandleEvent
}

func NewEvent(evt string, t EventTiming) EventPartial {
	h := &HandleEvent{Run: t}
	// for testing, allow empty event strings:
	if evt != "" {
		h.Events = []string{evt}
	}
	return EventPartial{h}
}

func (p EventPartial) Or(event string) EventPartial {
	p.data.Events = append(p.data.Events, event)
	return p
}

func (p EventPartial) Always(cb rt.Execute, cbs ...rt.Execute) Fragment {
	// for testing, allow empty blocks:
	if cb != nil {
		p.data.Calls = JoinCallbacks(cb, cbs)
	}
	return p.data
}

func (p *HandleEvent) GenFragment(src *S.Statements, top Topic) (err error) {
	opt := p.Run.GetOptions()
	for _, evt := range p.Events {
		fields := S.ListenFields{top.Subject, evt, p.Calls, opt}
		if e := src.NewEventHandler(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}
