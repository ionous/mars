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

func (x PrintNum) Execute(r rt.Runtime) (err error) {
	if n, e := x.GetNumber(r); e != nil {
		err = e
	} else if s := n.String(); len(s) > 0 {
		err = r.Print(s)
	} else {
		err = r.Println("<num>")
	}
	return err
}

func (x PrintText) Execute(r rt.Runtime) (err error) {
	if s, e := x.GetText(r); e != nil {
		err = e
	} else {
		err = r.Print(s.String())
	}
	return err
}

// Execute a little machine to add spaces between words, but not before punctuation.
func (p PrintLine) Execute(r rt.Runtime) (err error) {
	// var fin bytes.Buffer
	// rt.PushOutput(&fin)
	err = p.Strings.Execute(r)
	// rt.PopOutput()
	if err == nil {
		err = r.Println()
	}
	return
}
