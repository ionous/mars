package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type Statements []rt.Execute

func (ss Statements) Execute(run rt.Runtime) (err error) {
	for _, s := range ss {
		if e := s.Execute(run); e != nil {
			err = e
			break
		}
	}
	return err
}

type Error struct {
	Reason string
}

func StopHere() Error {
	return Error{}
}

// Fails expects the executed statement to return an error
type Fails struct {
	Other   rt.Execute
	Message string
}

func (x Fails) Execute(run rt.Runtime) (err error) {
	if e := x.Other.Execute(run); e == nil {
		err = errutil.New(x.Message)
	} else {
		run.Println("failed okay with", e)
	}
	return
}

func (x Error) Execute(run rt.Runtime) (err error) {
	return errutil.New(x.Reason)
}
