package g

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/rt"
)

// ReflectToTarget runs the passed action, flipping the source and target.
func ReflectToTarget(action string) Execute {
	return The("action.Target").Go(action, Id("action.Source"))
}

// ReflectToLocation invokes the passed action on the actor's current whereabouts.
// TODO: will have to become more sophisticated for being inside a box.
func ReflectToLocation(action string) Execute {
	return The("actor").Object("whereabouts").Go(action, Id("actor"))
}

// ReflectWithContext runs the passed action, shifting to target, context, source.
// FIX: i think it'd be better to first use ReflectToTarget, keeping the context as the third parameter
// and then ReflectToContext, possibly re-swapping source and target.
func ReflectWithContext(action string) Execute {
	return The("action.Target").Go(action, Id("action.Context"), Id("action.Source"))
}
