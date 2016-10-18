package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type Statements []rt.Execute

func (ss Statements) Execute(r rt.Runtime) (err error) {
	for _, s := range ss {
		if e := s.Execute(r); e != nil {
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

func (x Fails) Execute(r rt.Runtime) (err error) {
	if e := x.Other.Execute(r); e == nil {
		err = errutil.New(x.Message)
	} else {
		r.Println("failed okay with", e)
	}
	return
}

func (x Error) Execute(r rt.Runtime) (err error) {
	return errutil.New(x.Reason)
}
