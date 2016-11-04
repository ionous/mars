package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type ExecuteList []rt.Execute

func (x ExecuteList) Execute(run rt.Runtime) (err error) {
	for _, s := range x {
		if e := s.Execute(run); e != nil {
			err = e
			break
		}
	}
	return err
}

type StopNow struct{}

// Error satifies the golang error interface
func (x StopNow) Error() string {
	return "stop"
}

func (x StopNow) Execute(rt.Runtime) error {
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

type DoNothing struct{}

func (x DoNothing) Execute(rt.Runtime) error {
	return nil
}
