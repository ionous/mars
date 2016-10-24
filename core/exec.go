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

func (x Error) Execute(run rt.Runtime) (err error) {
	return x
}

// Error satifies the golang error interface
func (x Error) Error() string {
	return x.Reason
}

type StopNow struct {
}

// Error satifies the golang error interface
func (x StopNow) Error() string {
	return "stop"
}

func (x StopNow) Execute(run rt.Runtime) error {
	return x
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
