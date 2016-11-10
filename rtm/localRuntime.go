package rtm

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/ident"
)

type localRuntime struct {
	meta.Model
	rtm *Rtm
}

func (lr localRuntime) FindValue(s string) (meta.Generic, error) {
	return lr.rtm.scope.Top().FindValue(s)
}
func (lr localRuntime) ScopePath() rt.ScopePath {
	return lr.rtm.scope.Top().ScopePath()
}
func (lr localRuntime) Print(args ...interface{}) error {
	return lr.rtm.output.Print(args...)
}
func (lr localRuntime) Println(args ...interface{}) error {
	return lr.rtm.output.Println(args...)
}
func (lr localRuntime) RunAction(id ident.Id, scp rt.Scope, args ...meta.Generic) error {
	return lr.rtm.RunAction(id, scp, args...)
}
func (lr localRuntime) LookupParent(i meta.Instance) (meta.Instance, meta.Property, bool) {
	return lr.rtm.parents.LookupParent(i)
}
