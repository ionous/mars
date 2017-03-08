package core

import "github.com/ionous/mars/rt"

type Choose struct {
	If    rt.BoolEval   `mars:"if"`
	True  rt.Statements `mars:":"`
	False rt.Statements `mars:"otherwise:"`
}

type ChooseNum struct {
	If          rt.BoolEval
	True, False rt.NumberEval
}

type ChooseText struct {
	If          rt.BoolEval
	True, False rt.TextEval
}

type ChooseObj struct {
	If          rt.BoolEval
	True, False rt.ObjEval
}

func (x ChooseNum) GetNumber(run rt.Runtime) (ret rt.Number, err error) {
	if b, e := x.If.GetBool(run); e != nil {
		err = e
	} else {
		var next rt.NumberEval
		if b.Value {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetNumber(run)
		}
	}
	return
}

func (x ChooseText) GetText(run rt.Runtime) (ret rt.Text, err error) {
	if b, e := x.If.GetBool(run); e != nil {
		err = e
	} else {
		var next rt.TextEval
		if b.Value {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetText(run)
		}
	}
	return
}

func (x ChooseObj) GetObject(run rt.Runtime) (ret rt.Object, err error) {
	if b, e := x.If.GetBool(run); e != nil {
		err = e
	} else {
		var next rt.ObjEval
		if b.Value {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetObject(run)
		}
	}
	return
}

// Execute evals, eats the returns
func (x Choose) Execute(run rt.Runtime) (err error) {
	if b, e := x.If.GetBool(run); e != nil {
		err = e
	} else {
		var next rt.Statements
		if b.Value {
			next = x.True
		} else {
			next = x.False
		}
		err = next.ExecuteList(run)
	}
	return
}
