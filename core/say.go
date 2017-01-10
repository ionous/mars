package core

import (
	"bytes"
	"github.com/ionous/mars/rt"
)

// Say shortcut runs a bunch of statements and "collects" them via PrintLine
func Say(all ...interface{}) (ret rt.Execute) {
	if len(all) == 1 {
		ret = Print(all...)[0]
	} else {
		ret = PrintLine{Print(all...)}
	}
	return
}

func MakeText(all ...interface{}) rt.TextEval {
	txt := Print(all...)
	return Buffer{txt}
}

type Buffer struct {
	Buffer rt.Statements
}

func (m Buffer) GetText(run rt.Runtime) (ret rt.Text, err error) {
	var out bytes.Buffer
	defer run.PopOutput()
	run.PushOutput(&out)
	//
	if e := m.Buffer.ExecuteList(run); e != nil {
		err = e
	} else {
		ret = rt.Text{out.String()}
	}
	return
}

func Print(all ...interface{}) rt.Statements {
	sayWhat := []rt.Execute{}
	for _, a := range all {
		switch val := a.(type) {
		case int:
			sayWhat = append(sayWhat, PrintNum{I(val)})
		case rt.NumberEval:
			sayWhat = append(sayWhat, PrintNum{val})
		case string:
			sayWhat = append(sayWhat, PrintText{T(val)})
		case rt.TextEval:
			sayWhat = append(sayWhat, PrintText{val})
		case rt.Execute:
			// FIX: could buffer operations have a specialized interface implementation?
			sayWhat = append(sayWhat, val)
		case rt.ObjEval:
			sayWhat = append(sayWhat, PrintObj{val})
		case rt.NumListEval:
			l := ForEachNum{In: val, Go: rt.MakeStatements(PrintNum{GetNum{"@"}})}
			sayWhat = append(sayWhat, l)
		case rt.TextListEval:
			l := ForEachText{In: val, Go: rt.MakeStatements(PrintText{GetText{"@"}})}
			sayWhat = append(sayWhat, l)
		case rt.ObjListEval:
			l := ForEachObj{In: val, Go: rt.MakeStatements(PrintObj{GetObj{"@"}})}
			sayWhat = append(sayWhat, l)
		default:
			panic("say what?")
		}
	}
	return rt.MakeStatements(sayWhat...)
}
