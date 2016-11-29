package main

import (
	"encoding/json"
	"fmt"
	"github.com/ionous/mars"
	"github.com/ionous/mars/encode"
	"github.com/ionous/mars/script/test"
	"github.com/ionous/mars/std"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

type DataBlocks []encode.DataBlock

// we need more than the spec!

type Suite struct {
	Name  string
	Units []Unit
}

type Unit struct {
	Name   string
	Setup  encode.DataBlock
	Trials []Trial
}

type Trial struct {
	Name      string
	Imp       Imp
	Pre, Post DataBlocks
	Fini      encode.DataBlock
}

type Imp struct {
	Input   string
	Match   []string
	Args    DataBlocks
	Execute encode.DataBlock
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
	if newSetup, e := encode.Compute(src.Setup); e != nil {
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
	} else if fini, e := encode.Compute(src.Fini); e != nil {
		err = e
	} else {
		ret = Trial{src.Name, imp, pre, post, fini}
	}
	return
}

func addConditions(src test.Conditions) (ret DataBlocks, err error) {
	for _, cond := range src {
		if cmd, e := encode.Compute(cond); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addImp(src test.Imp) (ret Imp, err error) {
	if args, e := addArgs(src.Args); e != nil {
		err = e
	} else if cmd, e := encode.Compute(src.Execute); e != nil {
		err = e
	} else {
		ret = Imp{src.Input, src.Match, args, cmd}
	}
	return
}

func addArgs(src []meta.Generic) (ret DataBlocks, err error) {
	for _, a := range src {
		if cmd, e := encode.Compute(a); e != nil {
			err = e
			break
		} else {
			ret = append(ret, cmd)
		}
	}
	return
}

func addSuites(p *mars.Package) (ret []Suite, err error) {
	if cnt := len(p.Tests); cnt > 0 {
		fmt.Println("package", p.Name, "adding", cnt, "tests")
		for _, src := range p.Tests {
			if s, e := addSuite(src); e != nil {
				err = e
				break
			} else {
				ret = append(ret, s)
			}
		}
	}
	return
}

func addTests() (ret []Suite, err error) {
	m := make(encode.PackageMap)
	if packages, e := m.AddPackage(std.Std()); e != nil {
		err = e
	} else {
		for _, p := range packages {
			if ps, e := addSuites(p); e != nil {
				err = e
				break
			} else {
				ret = append(ret, ps...)
			}
		}
	}
	return
}

// Called("D"), Has("rank", 3))
// is failing. we need to be able to store integer/number literals
// thats true even if we were using make.go I -> rt.Number
// note: currently the compiler is making the number evals from numbers
// hrmmmm..... that means we have to support primitives. at leaste on parse
// even if not on set in the editors
func main() {
	if suites, e := addTests(); e != nil {
		fmt.Println("error:", e)
	} else if b, e := json.MarshalIndent(suites, "", " "); e != nil {
		fmt.Println("error:", e)
	} else {
		fmt.Println(string(b))
	}
}
