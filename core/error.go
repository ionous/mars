package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// print the text on success, return it as an error on failure
func Try(message string, b rt.BoolEval) rt.Execute {
	return Choose{If: b,
		True:  rt.MakeStatements(PrintText{T(message)}),
		False: rt.MakeStatements(Error{message}),
	}
}

// Error satifies all runtime evaluations;
// in all cases returning an error string provided by "reason".
type Error struct{ Reason string }

func (x Error) Execute(run rt.Runtime) error {
	return errutil.New(x.Reason)
}
func (x Error) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetText(run rt.Runtime) (ret rt.Text, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetNumberStream(run rt.Runtime) (ret rt.NumberStream, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetTextStream(run rt.Runtime) (ret rt.TextStream, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	err = x.Execute(run)
	return
}
