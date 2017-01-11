package encode

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/test"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// we need more than the spec!

type Suite struct {
	Name  string `json:"name,omitempty"`
	Units []Unit `json:"units,omitempty"`
}

type Unit struct {
	Name   string     `json:"name,omitempty"`
	Setup  DataBlocks `json:"setup,omitempty"`
	Trials []Trial    `json:"trials,omitempty"`
}

type Trial struct {
	Name string     `json:"name,omitempty"`
	Imp  *Imp       `json:"imp,omitempty"`
	Pre  DataBlocks `json:"pre,omitempty"`
	Post DataBlocks `json:"post,omitempty"`
	Fini DataBlocks `json:"fini,omitempty"`
}

type Imp struct {
	Input   string     `json:"input,omitempty"`
	Match   []string   `json:"match,omitempty"`
	Args    DataBlocks `json:"args,omitempty"`
	Execute DataBlocks `json:"exec,omitempty"`
}

func addSuite(src test.Suite) (ret Suite, err error) {
	if us, e := addUnits(src.Units); e != nil {
		err = errutil.New("couldnt add suite", src.Name, "because", e)
	} else {
		ret = Suite{src.Name, us}
	}
	return
}

func addUnits(src []test.Unit) (ret []Unit, err error) {
	for _, u := range src {
		if newUnit, e := addUnit(u); e != nil {
			err = errutil.New("couldnt add unit", u.Name, "because", e)
			break
		} else {
			ret = append(ret, newUnit)
		}
	}
	return
}

func addUnit(src test.Unit) (ret Unit, err error) {
	if newSetup, e := Computes(src.Setup); e != nil {
		err = errutil.New("couldnt add setup", e)
	} else if newTrials, e := addTrials(src.Trials); e != nil {
		err = e
	} else {
		ret = Unit{src.Name, newSetup, newTrials}
	}
	return
}

func addTrials(src []test.Trial) (ret []Trial, err error) {
	for _, t := range src {
		if newTrial, e := addTrial(t); e != nil {
			err = e
			break
		} else {
			ret = append(ret, newTrial)
		}
	}
	return
}

func addTrial(src test.Trial) (ret Trial, err error) {
	if imp, e := addImp(src.Imp); e != nil {
		err = e
	} else if pre, e := addConditions(src.Pre); e != nil {
		err = e
	} else if post, e := addConditions(src.Post); e != nil {
		err = e
	} else if fini, e := Computes(src.Fini); e != nil {
		err = e
	} else {
		ret = Trial{src.Name, imp, pre, post, fini}
	}
	return
}

func addConditions(src test.Conditions) (ret DataBlocks, err error) {
	for _, cond := range src {
		if cmd, e := Compute(cond); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addImp(src test.Imp) (ret *Imp, err error) {
	if args, e := addArgs(src.Args); e != nil {
		err = e
	} else if cmds, e := addStatements(src.Execute); e != nil {
		err = e
	} else if len(cmds) > 0 || len(args) > 0 || len(src.Match) > 0 || src.Input != "" {
		ret = &Imp{src.Input, src.Match, args, cmds}
	}
	return
}

func addArgs(src []meta.Generic) (ret DataBlocks, err error) {
	for _, a := range src {
		if cmd, e := Compute(a); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addStatements(src []rt.Execute) (ret DataBlocks, err error) {
	for _, a := range src {
		if cmd, e := Compute(a); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addSuites(tests []test.Suite) (ret []Suite, err error) {
	for _, src := range tests {
		if s, e := addSuite(src); e != nil {
			err = e
			break
		} else {
			ret = append(ret, s)
		}
	}
	return
}