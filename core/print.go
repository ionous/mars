package core

import (
	"github.com/ionous/mars/rt"
	"strconv"
)

type PrintNum struct {
	Num rt.NumberEval `mars:"print [num]"`
}

type PrintText struct {
	Text rt.TextEval `mars:"print [text]"`
}

type PrintObj struct {
	Obj rt.ObjEval `mars:"print [obj]"`
}

type PrintLine struct {
	Block rt.Statements
}

func (x PrintNum) Execute(run rt.Runtime) (err error) {
	if n, e := x.Num.GetNumber(run); e != nil {
		err = e
	} else if s := strconv.FormatFloat(n.Value, 'g', -1, 64); len(s) > 0 {
		err = run.Print(s)
	} else {
		err = run.Print("<num>")
	}
	return err
}

func (x PrintText) Execute(run rt.Runtime) (err error) {
	if s, e := x.Text.GetText(run); e != nil {
		err = e
	} else {
		err = run.Print(s)
	}
	return err
}

func (x PrintObj) Execute(run rt.Runtime) (err error) {
	if o, e := x.Obj.GetObject(run); e != nil {
		err = e
	} else {
		err = run.Print(o.GetOriginalName())
	}
	return err
}

func (p PrintLine) Execute(run rt.Runtime) (err error) {
	run.StartLine()
	err = p.Block.ExecuteList(run)
	run.EndLine()
	return
}
