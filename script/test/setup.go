package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/meta"
)

func NewSuite(name string, units ...Unit) Suite {
	return Suite{name, units}
}

func Setup(setup ...backend.Spec) Unit {
	return Unit{Setup: backend.SpecList{setup}}
}

func (u Unit) Try(name string, trials ...Trial) Unit {
	// FIX: move this
	u.Name = name
	u.Trials = append(u.Trials, trials...)
	return u
}

func Parse(input string) Trial {
	return Trial{Imp: Imp{Input: input}}
}

func Run(action string, args ...meta.Generic) Trial {
	if args == nil {
		args = make([]meta.Generic, 0, 1)
	}
	return Trial{Imp: Imp{Input: action, Args: args}}
}

func Execute(x rt.Execute) Trial {
	return Trial{Imp: Imp{Execute: x}}
}

func Expect(expect ...rt.BoolEval) Trial {
	return Trial{Post: expect}
}

func (h Trial) Match(match ...string) Trial {
	h.Imp.Match = match
	return h
}

func (h Trial) Expect(expect ...rt.BoolEval) Trial {
	h.Post = expect
	return h
}

func (h Trial) Else(fini rt.Execute) Trial {
	h.Fini = fini
	return h
}
