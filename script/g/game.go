package g

import (
	"github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/std/compat"
	"github.com/ionous/sashimi/meta"
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

func TheObject() compat.ScriptRef {
	return compat.ScriptRef{core.GetObj{"@"}}
}

// Go shortcut runs a bunch of statements
func Go(all ...rt.Execute) rt.Execute {
	return core.ExecuteList{all}
}

func Call(act string, args ...meta.Generic) core.GoCall {
	return core.GoCall{core.MakeStringId(act), args}
}
