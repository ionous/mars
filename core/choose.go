package core

import "github.com/ionous/mars/rt"

type Choose struct {
	If          rt.BoolEval
	True, False rt.Execute
}

type ChooseNum struct {
	If          rt.BoolEval
	True, False rt.NumEval
}

type ChooseText struct {
	If          rt.BoolEval
	True, False rt.TextEval
}

type ChooseRef struct {
	If          rt.BoolEval
	True, False rt.RefEval
}

func (x Choose) GetBool(r rt.Runtime) (ret bool, err error) {
	if b, e := x.If.GetBool(r); e != nil {
		err = e
	} else {
		var next rt.Execute
		if b {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			err = next.Execute(r)
		}
		ret = b
	}
	return
}

func (x ChooseNum) GetNumber(r rt.Runtime) (ret rt.Number, err error) {
	if b, e := x.If.GetBool(r); e != nil {
		err = e
	} else {
		var next rt.NumEval
		if b {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetNumber(r)
		}
	}
	return
}

func (x ChooseText) GetText(r rt.Runtime) (ret rt.Text, err error) {
	if b, e := x.If.GetBool(r); e != nil {
		err = e
	} else {
		var next rt.TextEval
		if b {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetText(r)
		}
	}
	return
}

func (x ChooseRef) GetReference(r rt.Runtime) (ret rt.Reference, err error) {
	if b, e := x.If.GetBool(r); e != nil {
		err = e
	} else {
		var next rt.RefEval
		if b {
			next = x.True
		} else {
			next = x.False
		}
		if next != nil {
			ret, err = next.GetReference(r)
		}
	}
	return
}

// Execute evals, eats the returns
func (x Choose) Execute(r rt.Runtime) error {
	_, e := x.GetBool(r)
	return e
}
