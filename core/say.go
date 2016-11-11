package core

import (
	"bytes"
	"github.com/ionous/mars/rt"
)

// Say shortcut runs a bunch of statements and "collects" them via PrintLine
func Say(all ...interface{}) rt.Execute {
	txt := Format(all...)
	return PrintLine{txt}
}

func MakeText(all ...interface{}) rt.TextEval {
	txt := Format(all...)
	return Buffer{txt}
}

type Buffer struct {
	Buffer rt.Execute
}

func (m Buffer) GetText(run rt.Runtime) (ret rt.Text, err error) {
	var out bytes.Buffer
	defer run.PopOutput()
	run.PushOutput(&out)
	//
	if e := m.Buffer.Execute(run); e != nil {
		err = e
	} else {
		ret = rt.Text(out.String())
	}
	return
}

func Format(all ...interface{}) rt.Execute {
	sayWhat := ExecuteList{}
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
		case rt.ObjEval:
			sayWhat = append(sayWhat, PrintObject{val})
		default:
			panic("say what?")
		}
	}
	return sayWhat
}
