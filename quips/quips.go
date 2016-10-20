package quips

import (
	//	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
)

func Introduce(greeter string) GreeterPhrase {
	return GreeterPhrase{greeter: greeter}
}

func (g GreeterPhrase) To(greeted string) GreetedPhrase {
	g.greeted = greeted
	return GreetedPhrase(g)
}

func (g GreetedPhrase) WithDefault() rt.Execute {
	return g.With("")
}

func (g GreetedPhrase) With(greeting string) rt.Execute {
	// FIX: with inference compiler, in theory, this check wouldnt be needed.
	panic("with")
	// greeter, greeted, greets := R(g.greeter), R(g.greeted), R(greeting)
	// return Choose{
	// 	If: All(
	// 		Equals{R("player"), greeter},
	// 		Exists{greeted}),
	// 	True:  GoCall{Who: greeter, Run: "be greeted by", What: greeted, With: greets},
	// 	False: Error{"invalid greeting " + g.greeter + " " + g.greeted + " " + greeting}}
}

type greetingData struct {
	greeter, greeted string
}
type GreeterPhrase greetingData
type GreetedPhrase greetingData
