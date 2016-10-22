package g

import (
	. "github.com/ionous/mars/core"
	rt "github.com/ionous/mars/rt"
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

type Game struct {
	*ScriptName
	*ScriptRef
}

// GetObject searches through the scope for an object matching Name
func (h ScriptName) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	s, _ := run.GetScope()
	if v, e := s.FindValue(h.Name); e != nil {
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

func (h ScriptRef) Go(run string, all ...interface{}) rt.Execute {
	parms := rt.Parameters{}
	for _, a := range all {
		var ps rt.ParameterSource
		switch val := a.(type) {
		// note, rt.Number implements rt.NumEval. no need for a separate switch
		case rt.NumEval:
			ps = CallWithNum{val}
		case int:
			ps = CallWithNum{I(val)}
		case float64:
			ps = CallWithNum{N(val)}
			// note, rt.Number( implements rt.TextEval. no need for a separate switch)
		case rt.TextEval:
			ps = CallWithText{val}
		case string:
			ps = CallWithText{T(val)}
		// note, rt.Object implements rt.ObjEval
		case rt.ObjEval:
			ps = CallWithRef{val}
		default:
			panic("go what?")
		}
		parms = append(parms, ps)
	}
	panic("go call shouldnt push the object into scope, should it?")
	return GoCall{
		//Action:     P(h.Ref, run),
		Parameters: parms,
	}
}
