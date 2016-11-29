package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
)

func NewCanDo(action types.NamedAction) CanDoPhrase {
	return CanDoPhrase{ActionName: action}
}

func (c CanDoPhrase) And(doing types.NamedEvent) RequiresWhatPhrase {
	c.EventName = doing
	return RequiresWhatPhrase(c)
}

func (c RequiresWhatPhrase) RequiresNothing() Fragment {
	c.Requires = &RequiresNothing{}
	return (*CanDoIt)(&c)
}

// FIX: class name must be singular right now :(
func (c RequiresWhatPhrase) RequiresTwo(class types.NamedClass) Fragment {
	c.Requires = &RequiresTwo{class}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOnly(target types.NamedClass) Fragment {
	c.Requires = &RequiresOnly{Target: target}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOne(target types.NamedClass) RequiresMorePhrase {
	c.Requires = &Requires{Target: target}
	return RequiresMorePhrase(c)
}

func (c RequiresMorePhrase) AndOne(context types.NamedClass) Fragment {
	req := c.Requires.(*Requires)
	req.Context = context
	return (*CanDoIt)(&c)
}

//
type CanDoPhrase CanDoIt
type RequiresWhatPhrase CanDoIt
type RequiresMorePhrase CanDoIt

type ActionAssertion struct {
	RequiresWhatPhrase
	Target, Context string
}

type CanDoIt struct {
	ActionName types.NamedAction `mars:"can [act]"`
	EventName  types.NamedEvent  `mars:"and [acting]"`
	Requires   ActionRequirements
}

type ActionRequirements interface {
	TargetClass() types.NamedClass
	ContextClass() types.NamedClass
}

type RequiresNothing struct {
}

type Requires struct {
	Target  types.NamedClass `mars:"one"`
	Context types.NamedClass `mars:"and one"`
}

type RequiresOnly struct {
	Target types.NamedClass
}

type RequiresTwo struct {
	Class types.NamedClass `mars:"classes"`
}

func (*RequiresNothing) TargetClass() types.NamedClass  { return "" }
func (*RequiresNothing) ContextClass() types.NamedClass { return "" }

func (r *Requires) TargetClass() types.NamedClass  { return r.Target }
func (r *Requires) ContextClass() types.NamedClass { return r.Context }

func (r *RequiresOnly) TargetClass() types.NamedClass  { return r.Target }
func (r *RequiresOnly) ContextClass() types.NamedClass { return "" }

func (r *RequiresTwo) TargetClass() types.NamedClass  { return r.Class }
func (r *RequiresTwo) ContextClass() types.NamedClass { return r.Class }

func (c *CanDoIt) GenFragment(src *S.Statements, top Topic) (err error) {
	if top.Subject == "" {
		err = errutil.New("action", c.ActionName, "has no subject")
	} else {
		fields := S.ActionAssertionFields{
			c.ActionName, c.EventName,
			top.Subject, c.Requires.TargetClass(), c.Requires.ContextClass()}
		err = src.NewActionAssertion(fields, S.UnknownLocation)
	}
	return
}
