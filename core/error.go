package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

// Error satifies all runtime evaluations;
// in all cases returning an error string provided by "reason".
type Error struct{ Reason rt.TextEval }

func (x Error) Execute(run rt.Runtime) (err error) {
	if s, e := x.Reason.GetText(run); e != nil {
		err = errutil.New("error processing error", e)
	} else {
		err = errutil.New(s)
	}
	return err
}
func (x Error) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetNum(run rt.Runtime) (ret rt.Number, err error) {
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
func (x Error) GetCount() int {
	return 0
}
func (x Error) GetNumberIdx(run rt.Runtime, _ int) (ret rt.Number, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetTextIdx(run rt.Runtime, _ int) (ret rt.Text, err error) {
	err = x.Execute(run)
	return
}
func (x Error) GetReferenceIdx(run rt.Runtime, _ int) (ret rt.Reference, err error) {
	err = x.Execute(run)
	return
}
