package internal

import (
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/errutil"
)

func NewCanDo(ActionName string) CanDoPhrase {
	return CanDoPhrase{ActionName}
}

//starts a requirements phrase for deciding how to provide nouns...
func (f CanDoPhrase) And(doing string) RequiresWhatPhrase {
	return RequiresWhatPhrase{f, doing}
}

// the Target will be the same as the source
func (f RequiresWhatPhrase) RequiresNothing() Fragment {
	return ActionAssertion{RequiresWhatPhrase: f}
}

// the Target and the Context will be input by the user, and will both be of the passed class
// FIX: class must be singular right now :(
func (f RequiresWhatPhrase) RequiresTwo(class string) Fragment {
	return ActionAssertion{RequiresWhatPhrase: f, Target: class, Context: class}
}

// the Target will be input by the user, and will of the passed class
func (f RequiresWhatPhrase) RequiresOne(class string) ActionAssertion {
	return ActionAssertion{RequiresWhatPhrase: f, Target: class}
}

// the Context will be input by the user, and will of the passed class
func (f ActionAssertion) AndOne(class string) Fragment {
	f.Context = class
	return f
}

//
type CanDoPhrase struct {
	ActionName string
}
type RequiresWhatPhrase struct {
	CanDoPhrase
	EventName string
}
type ActionAssertion struct {
	RequiresWhatPhrase
	Target, Context string
}

func (f ActionAssertion) GenFragment(src *S.Statements, top Topic) (err error) {
	if top.Subject == "" {
		err = errutil.New("action", f.ActionName, "has no subject")
	} else {
		fields := S.ActionAssertionFields{
			f.ActionName, f.EventName,
			top.Subject, f.Target, f.Context}
		err = src.NewActionAssertion(fields, S.UnknownLocation)
	}
	return
}
