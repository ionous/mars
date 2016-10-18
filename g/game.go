package g

import (
	. "github.com/ionous/mars/core"
	rt "github.com/ionous/mars/rt"
)

type With struct {
	Ref rt.RefEval
}

func The(s string) With {
	return With{R(s)}
}

func (h With) Go(run string, all ...interface{}) rt.Execute {
	parms := rt.Parameters{}
	for _, a := range all {
		var ps rt.ParameterSource
		switch val := a.(type) {
		// in case of g.The()
		case With:
			ps = CallWithRef{val.Ref}
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
		// note, rt.Reference implements rt.RefEval
		case rt.RefEval:
			ps = CallWithRef{val}
		default:
			panic("go what?")
		}
		parms = append(parms, ps)
	}
	return GoCall{
		Action:     P(h.Ref, run),
		Parameters: parms,
	}
}

// Say shortcut runs a bunch of statements and "collects" them via PrintLine
func Say(all ...interface{}) rt.Execute {
	sayWhat := Statements{}
	for _, a := range all {
		switch val := a.(type) {
		case int:
			sayWhat = append(sayWhat, PrintNum{I(val)})
		case rt.NumEval:
			sayWhat = append(sayWhat, PrintNum{val})
		case string:
			sayWhat = append(sayWhat, PrintText{T(val)})
		case rt.TextEval:
			sayWhat = append(sayWhat, PrintText{val})
		case rt.Execute:
			// FIX: could buffer operations have a specialized interface implementation?
			sayWhat = append(sayWhat, val)
		default:
			panic("say what?")
		}
	}
	return PrintLine{sayWhat}
}
