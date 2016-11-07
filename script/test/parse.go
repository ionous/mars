package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

func Parse(input string) TrialHelper {
	return TrialHelper{Trial{Imp: Imp{Input: input}}}
}

func Run(action string, args ...meta.Generic) TrialHelper {
	if args == nil {
		args = make([]meta.Generic, 0, 1)
	}
	return TrialHelper{Trial{Imp: Imp{Input: action, Args: args}}}
}

func Execute(match string, x rt.Execute) Trial {
	return Trial{Imp: Imp{Match: match, Execute: x}}
}

type TrialHelper struct {
	trial Trial
}

func (h TrialHelper) Match(match string) Trial {
	h.trial.Imp.Match = match
	return h.trial
}

func (h Trial) Expect(expect ...rt.BoolEval) Trial {
	h.Post = expect
	return h
}
