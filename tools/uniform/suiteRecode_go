package uniform

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/test"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// we need more than the spec!

type SuiteData struct {
	Name  string     `json:"name,omitempty"`
	Units []UnitData `json:"units,omitempty"`
}

type UnitData struct {
	Name   string      `json:"name,omitempty"`
	Setup  []DataBlock `json:"setup,omitempty"`
	Trials []TrialData `json:"trials,omitempty"`
}

type TrialData struct {
	Name string      `json:"name,omitempty"`
	Imp  *ImpData    `json:"imp,omitempty"`
	Pre  []DataBlock `json:"pre,omitempty"`
	Post []DataBlock `json:"post,omitempty"`
	Fini []DataBlock `json:"fini,omitempty"`
}

type ImpData struct {
	Input   string      `json:"input,omitempty"`
	Match   []string    `json:"match,omitempty"`
	Args    []DataBlock `json:"args,omitempty"`
	Execute []DataBlock `json:"exec,omitempty"`
}

func addSuiteData(ue UniformEncoder, src test.Suite) (ret SuiteData, err error) {
	if us, e := addUnits(ue, src.Units); e != nil {
		err = errutil.New("couldn't add suite", src.Name, "because", e)
	} else {
		ret = SuiteData{src.Name, us}
	}
	return
}

func addUnits(ue UniformEncoder, src []test.Unit) (ret []UnitData, err error) {
	for _, u := range src {
		if newUnit, e := addUnit(ue, u); e != nil {
			err = errutil.New("couldn't add unit", u.Name, "because", e)
			break
		} else {
			ret = append(ret, newUnit)
		}
	}
	return
}

func addUnit(ue UniformEncoder, src test.Unit) (ret UnitData, err error) {
	if setupData, e := addSetup(ue, src); e != nil {
		err = errutil.New("couldn't add setup", e)
	} else if newTrials, e := addTrials(ue, src.Trials); e != nil {
		err = e
	} else {
		ret = UnitData{src.Name, setupData, newTrials}
	}
	return
}

func addSetup(ue UniformEncoder, src test.Unit) (ret []DataBlock, err error) {
	for _, s := range src.Setup {
		if d, e := ue.Compute(s); e != nil {
			err = e
			break
		} else {
			ret = append(ret, d)
		}
	}
	return
}

func addTrials(ue UniformEncoder, src []test.Trial) (ret []TrialData, err error) {
	for _, t := range src {
		if newTrial, e := addTrial(ue, t); e != nil {
			err = e
			break
		} else {
			ret = append(ret, newTrial)
		}
	}
	return
}

func addTrial(ue UniformEncoder, src test.Trial) (ret TrialData, err error) {
	if imp, e := addImp(ue, src.Imp); e != nil {
		err = e
	} else if pre, e := addConditions(ue, src.Pre); e != nil {
		err = e
	} else if post, e := addConditions(ue, src.Post); e != nil {
		err = e
	} else if fini, e := addExecute(ue, src.Fini); e != nil {
		err = e
	} else {
		ret = TrialData{src.Name, imp, pre, post, fini}
	}
	return
}

func addExecute(ue UniformEncoder, src rt.Execute) (ret []DataBlock, err error) {
	if src != nil {
		if c, e := ue.Compute(src); e != nil {
			err = e
		} else {
			ret = append(ret, c)
		}
	}
	return
}

func addConditions(ue UniformEncoder, src test.Conditions) (ret []DataBlock, err error) {
	for _, cond := range src {
		if cmd, e := ue.Compute(cond); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addImp(ue UniformEncoder, src test.Imp) (ret *ImpData, err error) {
	if args, e := addArgs(ue, src.Args); e != nil {
		err = e
	} else if cmds, e := addStatements(ue, src.Execute); e != nil {
		err = e
	} else if len(cmds) > 0 || len(args) > 0 || len(src.Match) > 0 || src.Input != "" {
		ret = &ImpData{src.Input, src.Match, args, cmds}
	}
	return
}

func addArgs(ue UniformEncoder, src []meta.Generic) (ret []DataBlock, err error) {
	for _, a := range src {
		if cmd, e := ue.Compute(a); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addStatements(ue UniformEncoder, src []rt.Execute) (ret []DataBlock, err error) {
	for _, a := range src {
		if cmd, e := ue.Compute(a); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func MakeUniformSuites(ue UniformEncoder, tests []test.Suite) (ret []SuiteData, err error) {
	for _, src := range tests {
		if s, e := addSuiteData(ue, src); e != nil {
			err = e
			break
		} else {
			ret = append(ret, s)
		}
	}
	return
}
