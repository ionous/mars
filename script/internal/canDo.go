package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	. "github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
)

func NewCanDo(action ActionName) CanDoPhrase {
	return CanDoPhrase{ActionName: action}
}

func (c CanDoPhrase) And(doing EventName) RequiresWhatPhrase {
	c.EventName = doing
	return RequiresWhatPhrase(c)
}

func (c RequiresWhatPhrase) RequiresNothing() Fragment {
	c.Requires = &RequiresNothing{}
	return (*CanDoIt)(&c)
}

// FIX: class name must be singular right now :(
func (c RequiresWhatPhrase) RequiresTwo(class ClassName) Fragment {
	c.Requires = &RequiresTwo{class}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOnly(target ClassName) Fragment {
	c.Requires = &RequiresOnly{Target: target}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOne(target ClassName) RequiresMorePhrase {
	c.Requires = &Requires{Target: target}
	return RequiresMorePhrase(c)
}

func (c RequiresMorePhrase) AndOne(context ClassName) Fragment {
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
	ActionName ActionName `mars:"can [act]"`
	EventName  EventName  `mars:"and [acting]"`
	Requires   ActionRequirements
}

type ActionRequirements interface {
	TargetClass() ClassName
	ContextClass() ClassName
}

type RequiresNothing struct {
}

type Requires struct {
	Target  ClassName `mars:"one"`
	Context ClassName `mars:"and one"`
}

type RequiresOnly struct {
	Target ClassName
}

type RequiresTwo struct {
	Class ClassName `mars:"classes"`
}

func (*RequiresNothing) TargetClass() ClassName  { return "" }
func (*RequiresNothing) ContextClass() ClassName { return "" }

func (r *Requires) TargetClass() ClassName  { return r.Target }
func (r *Requires) ContextClass() ClassName { return r.Context }

func (r *RequiresOnly) TargetClass() ClassName  { return r.Target }
func (r *RequiresOnly) ContextClass() ClassName { return "" }

func (r *RequiresTwo) TargetClass() ClassName  { return r.Class }
func (r *RequiresTwo) ContextClass() ClassName { return r.Class }

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
