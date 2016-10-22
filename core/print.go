package core

import (
	"github.com/ionous/mars/rt"
)

type PrintNum struct {
	rt.NumEval
}

type PrintText struct {
	rt.TextEval
}

// PrintLine
type PrintLine struct {
	Strings Statements
}

func (x PrintNum) Execute(run rt.Runtime) (err error) {
	if n, e := x.GetNumber(run); e != nil {
		err = e
	} else if s := n.String(); len(s) > 0 {
		err = run.Print(s)
	} else {
		err = run.Println("<num>")
	}
	return err
}

func (x PrintText) Execute(run rt.Runtime) (err error) {
	if s, e := x.GetText(run); e != nil {
		err = e
	} else {
		err = run.Print(s.String())
	}
	return err
}

// Execute a little machine to add spaces between words, but not before punctuation.
func (p PrintLine) Execute(run rt.Runtime) (err error) {
	// var fin bytes.Buffer
	// rt.PushOutput(&fin)
	err = p.Strings.Execute(run)
	// rt.PopOutput()
	if err == nil {
		err = run.Println()
	}
	return
}
