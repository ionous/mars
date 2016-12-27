package core

import (
	"github.com/ionous/mars/rt"
	"strconv"
)

type PrintNum struct {
	Num rt.NumberEval
}

type PrintText struct {
	Text rt.TextEval
}

type PrintObj struct {
	Obj rt.ObjEval
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
		err = run.Println("<num>")
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

// Execute a little machine to add spaces between words, but not before punctuation.
func (p PrintLine) Execute(run rt.Runtime) (err error) {
	err = p.Block.ExecuteList(run)
	if err == nil {
		err = run.Println()
	}
	return
}
