package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/errutil"
)

func (c CanDoPhrase) And(doing string) RequiresWhatPhrase {
	c.EventName = doing
	return RequiresWhatPhrase(c)
}

func (c RequiresWhatPhrase) RequiresNothing() Fragment {
	c.Requires = &RequiresNothing{}
	return (*CanDoIt)(&c)
}

// FIX: class name must be singular right now :(
func (c RequiresWhatPhrase) RequiresTwo(class string) Fragment {
	c.Requires = &RequiresTwo{class}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOnly(target string) Fragment {
	c.Requires = &RequiresOnly{Target: target}
	return (*CanDoIt)(&c)
}

func (c RequiresWhatPhrase) RequiresOne(target string) RequiresMorePhrase {
	c.Requires = &Requires{Target: target}
	return RequiresMorePhrase(c)
}

func (c RequiresMorePhrase) AndOne(context string) Fragment {
	req := c.Requires.(*Requires)
	req.Context = context
	return (*CanDoIt)(&c)
}

//
type CanDoPhrase CanDoIt
type RequiresWhatPhrase CanDoIt
type RequiresMorePhrase CanDoIt

type CanDoIt struct {
	ActionName string `mars:"can [act];action"`
	EventName  string `mars:"and [acting];event"`
	Requires   ActionRequirements
}

type ActionRequirements interface {
	TargetClass() string
	ContextClass() string
}

type RequiresNothing struct {
}

type Requires struct {
	Target  string `mars:"one;class"`
	Context string `mars:"and one;class"`
}

type RequiresOnly struct {
	Target string `mars:";class"`
}

type RequiresTwo struct {
	Class string `mars:"classes;class"`
}

func (*RequiresNothing) TargetClass() string  { return "" }
func (*RequiresNothing) ContextClass() string { return "" }

func (r *Requires) TargetClass() string  { return r.Target }
func (r *Requires) ContextClass() string { return r.Context }

func (r *RequiresOnly) TargetClass() string  { return r.Target }
func (r *RequiresOnly) ContextClass() string { return "" }

func (r *RequiresTwo) TargetClass() string  { return r.Class }
func (r *RequiresTwo) ContextClass() string { return r.Class }

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
