package test

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

type Suite struct {
	Name  string
	Units []Unit
}

type Unit struct {
	Name   string
	Setup  backend.Declaration
	Trials []Trial
}

func (s Suite) String() string {
	return s.Name
}

type Trial struct {
	Name      string
	Imp       Imp
	Pre, Post Conditions
	Fini      rt.Execute
}

func (u Unit) Test(try Trytime) (err error) {
	for _, t := range u.Trials {
		if e := t.Test(try); e != nil {
			err = e
			break
		}
	}
	return
}

func (t Trial) Test(try Trytime) (err error) {
	if e := t.Pre.Test(try); e != nil {
		err = e
	} else if e := t.Imp.Run(try); e != nil {
		err = e
	} else if e := t.Post.Test(try); e != nil {
		err = e
		if t.Fini != nil {
			if s, e := try.Execute(rt.MakeStatements(t.Fini)); e != nil {
				panic(e)
			} else {
				err = errutil.New(err, strings.Join(s, ";"))
			}
		}
	}
	return
}
