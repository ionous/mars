package g

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/std/compat"
)

func Say(args ...interface{}) rt.Execute {
	return core.Say(args...)
}

func The(s string) compat.ScriptRef {
	return compat.ScriptRef{core.Named{s}}
}

func StopHere() rt.Execute {
	return core.StopNow{}
}

var Our = The
