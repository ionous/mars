package quips

import (
	. "github.com/ionous/mars/core"
)

func Introduce(greeter string) GreeterPhrase {
	return GreeterPhrase{greeter: greeter}
}

func (g GreeterPhrase) To(greeted string) GreetedPhrase {
	g.greeted = greeted
	return GreetedPhrase(g)
}

func (g GreetedPhrase) WithDefault() Execute {
	return g.With("")
}

func (g GreetedPhrase) With(greeting string) Execute {
	// FIX: with inference compiler, in theory, this check wouldnt be needed.
	greeter, greeted, greets := R(g.greeter), R(g.greeted), R(greeting)
	return Choose{
		If: All(
			Equals{R("player"), greeter},
			Exists{greeted}),
		True:  Go{Who: greeter, Run: "be greeted by", What: greeted, With: greets},
		False: Error{"invalid greeting " + g.greeter + " " + g.greeted + " " + greeting}}
}

type greetingData struct {
	greeter, greeted string
}
type GreeterPhrase greetingData
type GreetedPhrase greetingData
