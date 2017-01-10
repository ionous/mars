package internal

import (
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

type EventTiming interface {
	GetOptions() S.ListenOptions
}

type BeforeEvent struct{}

type AfterEvent struct{}

type WhenEvent struct{}

type HandleEvent struct {
	Run    EventTiming  `mars:"[handle event]"`
	Events []string     `mars:"[events]"`
	Calls  []rt.Execute `mars:"always:"`
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
	events := []string{evt}
	return EventPartial{&HandleEvent{t, events, nil}}
}

func (p EventPartial) Or(event types.NamedEvent) EventPartial {
	p.data.Events = append(p.data.Events, event.String())
	return p
}

func (p EventPartial) Always(cb rt.Execute, cbs ...rt.Execute) Fragment {
	p.data.Calls = JoinCallbacks(cb, cbs)
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
