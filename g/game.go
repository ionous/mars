package g

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

// ScriptRef provides (some) backwards compatibility with the old game interface
type ScriptRef struct {
	rt.ObjEval
}

func The(s string) ScriptRef {
	return ScriptRef{Named{s}}
}

func StopHere() rt.Execute {
	return StopNow{}
}

type Game struct {
	*Named
	*ScriptRef
}

// ex. g.Say(g.The("player").Text("greeting"))
func (h ScriptRef) Text(name PropertyName) rt.TextEval {
	return PropertyText{h, name}
}

func (h ScriptRef) Num(name PropertyName) rt.NumEval {
	return PropertyNum{h, name}
}

func (h ScriptRef) Object(name PropertyName) ScriptRef {
	return ScriptRef{PropertyRef{h, name}}
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
