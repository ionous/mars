package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
	"io"
)

type localRuntime struct {
	meta.Model
	rtm *Rtm
}

func (lr localRuntime) FindValue(name string) (meta.Generic, error) {
	return lr.rtm.scope.Top().FindValue(name)
}
func (lr localRuntime) ScopePath() rt.ScopePath {
	return lr.rtm.scope.Top().ScopePath()
}
func (lr localRuntime) Print(args ...interface{}) (err error) {
	if lr.rtm.lineWait {
		err = lr.rtm.output.Print(args...)
	} else {
		err = lr.rtm.output.Println(args...)
	}
	return
}
func (lr localRuntime) StartLine() {
	lr.rtm.lineWait = true
}
func (lr localRuntime) EndLine() {
	lr.rtm.output.Println("")
	lr.rtm.lineWait = false
}

func (lr localRuntime) RunAction(id ident.Id, scp rt.Scope, args ...meta.Generic) error {
	return lr.rtm.RunAction(id, scp, args...)
}
func (lr localRuntime) FindParent(rt.Object) (ret rt.Object, err error) {
	return
}
func (lr localRuntime) PushOutput(out io.Writer) {
	lr.rtm.output.PushOutput(out)
}

func (lr localRuntime) PopOutput() {
	lr.rtm.output.PopOutput()
}
