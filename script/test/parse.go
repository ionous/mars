package test

import (
	"github.com/ionous/mars/rt"
)

func Parse(input string) TrialHelper {
	return TrialHelper{Trial{Imp: Imp{Input: input}}}
}

func Run(match string, x rt.Execute) Trial {
	return Trial{Imp: Imp{Match: match, Execute: x}}
}

type TrialHelper struct {
	trial Trial
}

func (h TrialHelper) Match(match string) TrialHelper {
	h.trial.Imp.Match = match
	return h
}

func (h TrialHelper) Expect(expect ...rt.BoolEval) Trial {
	h.trial.Post = expect
	return h.trial
}
