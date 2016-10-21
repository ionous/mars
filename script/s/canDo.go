package s

import (
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/frag"
	S "github.com/ionous/sashimi/source"
)

// active verb:
func Can(verb string) CanDoPhrase {
	return CanDoPhrase{verb}
}

//starts a requirements phrase for deciding how to provide nouns...
func (canDo CanDoPhrase) And(doing string) RequiresWhatPhrase {
	return RequiresWhatPhrase{canDo, doing}
}

// the target will be the same as the source
func (canDo RequiresWhatPhrase) RequiresNothing() frag.Fragment {
	return ActionAssertionFragment{RequiresWhatPhrase: canDo}
}

// the target and the context will be input by the user, and will both be of the passed class
// FIX: class must be singular right now :(
func (canDo RequiresWhatPhrase) RequiresTwo(class string) frag.Fragment {
	return ActionAssertionFragment{RequiresWhatPhrase: canDo, target: class, context: class}
}

// the target will be input by the user, and will of the passed class
func (canDo RequiresWhatPhrase) RequiresOne(class string) ActionAssertionFragment {
	return ActionAssertionFragment{RequiresWhatPhrase: canDo, target: class}
}

// the context will be input by the user, and will of the passed class
func (canDo ActionAssertionFragment) AndOne(class string) frag.Fragment {
	canDo.context = class
	return canDo
}

//
type CanDoPhrase struct {
	actionName string
}
type RequiresWhatPhrase struct {
	CanDoPhrase
	eventName string
}
type ActionAssertionFragment struct {
	RequiresWhatPhrase
	target, context string
}

func (canDo ActionAssertionFragment) Build(src script.Source, top frag.Topic) error {
	fields := S.ActionAssertionFields{
		canDo.actionName, canDo.eventName,
		top.Subject, canDo.target, canDo.context}
	return src.NewActionAssertion(fields, script.Unknown)
}
