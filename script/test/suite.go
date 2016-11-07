package test

import (
	"github.com/ionous/mars/script/backend"
)

type Suite struct {
	Name   string
	Setup  backend.Script // an array of phrases
	Trials []Trial
}

func (s Suite) String() string {
	return s.Name
}

type Trial struct {
	Name      string
	Imp       Imp
	Pre, Post Conditions
}

func (suite Suite) Test(try Trytime) (err error) {
	for _, t := range suite.Trials {
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
	}
	return
}
