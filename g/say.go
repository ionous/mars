package g

import (
	. "github.com/ionous/mars/core"
	rt "github.com/ionous/mars/rt"
)

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
