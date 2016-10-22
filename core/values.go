package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

// FIX: i think this is all wrong:
// since generic stores evals, it should store list evals
// and so we shouldnt need these functions too
type ValueList struct {
	meta.Values
}

func (list ValueList) GetCount() int {
	return list.NumValue()
}

// FIX: see above
func (list ValueList) GetValueIdx(_ rt.Runtime, i int) (ret meta.Value, err error) {
	if n := list.NumValue(); i >= n {
		err = errutil.New("out of range", i, n)
	} else {
		ret = list.ValueNum(i)
	}
	return
}

type NumberValueList ValueList

func (list NumberValueList) GetNumberIdx(run rt.Runtime, i int) (ret rt.Number, err error) {
	if v, e := ValueList(list).GetValueIdx(run, i); e != nil {
		err = e
	} else {
		ret = N(v.GetNum())
	}
	return
}

type TextValueList ValueList

func (list TextValueList) GetTextIdx(run rt.Runtime, i int) (ret rt.Text, err error) {
	if v, e := ValueList(list).GetValueIdx(run, i); e != nil {
		err = e
	} else {
		ret = T(v.GetText())
	}
	return
}

type RefValueList ValueList

func (list RefValueList) GetReferenceIdx(run rt.Runtime, i int) (ret rt.Reference, err error) {
	if v, e := ValueList(list).GetValueIdx(run, i); e != nil {
		err = e
	} else {
		ret = rt.Reference(v.GetObject())
	}
	return
}
