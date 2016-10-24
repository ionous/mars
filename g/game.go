package g

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// ScriptName searches for objects by name, as opposed to core.Id which uses direct lookup.
type ScriptName struct {
	Name string
}

// ScriptRef provides (some) backwards compatibility with the old game interface
type ScriptRef struct {
	rt.ObjEval
}

func The(s string) ScriptRef {
	return ScriptRef{ScriptName{s}}
}

func StopHere() rt.Execute {
	return StopNow{}
}

type Game struct {
	*ScriptName
	*ScriptRef
}

// GetObject searches through the scope for an object matching Name
func (h ScriptName) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if v, e := run.FindValue(h.Name); e != nil {
		err = errutil.New("ScriptRef.GetObject", e)
	} else if x, ok := v.(rt.ObjEval); !ok {
		err = errutil.New("ScriptRef.GetObject", h.Name, "is not an object")
	} else if r, e := x.GetObject(run); e != nil {
		err = errutil.New("ScriptRef.GetObject", e)
	} else {
		ret = r
	}
	return
}

// ex. g.Say(g.The("player").Text("greeting"))
func (h ScriptRef) Text(name PropertyName) rt.TextEval {
	return TextProperty{h, name}
}

func (h ScriptRef) Num(name PropertyName) rt.NumEval {
	return NumProperty{h, name}
}

func (h ScriptRef) Object(name PropertyName) ScriptRef {
	return ScriptRef{RefProperty{h, name}}
}

// 	// Get returns the named property.
// 	Get(string) IValue -> Num(), Text(), etc. might just return property, and have the caller use the appropriate field.

// 	ObjectList(string) []IObject
// 	Set(string, IObject)
// 	SetNum(string, float64)
// 	SetText(string, string)

// g.The("player").Go("test nothing"),
func (h ScriptRef) Go(run string, all ...interface{}) rt.Execute {
	parms := make([]meta.Generic, len(all)+1)
	parms[0] = h
	for i, a := range all {
		var ps meta.Generic
		switch val := a.(type) {
		case int:
			ps = I(val)
		case float64:
			ps = N(val)
		// note, rt.Number implements rt.NumEval. no need for a separate switch
		case rt.NumEval:
			ps = val
		case string:
			ps = T(val)
		// note, rt.Text  implements rt.TextEval. no need for a separate switch
		case rt.TextEval:
			ps = val
		// note, rt.Object implements rt.ObjEval, no need for a separate switch
		case rt.ObjEval:
			ps = val
		default:
			panic("go what?")
		}
		parms[i+1] = ps
	}
	return GoCall{
		Action:     MakeStringId(run),
		Parameters: parms,
	}
}
