package core

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
)

type ValueList struct {
	meta.Values
}

func (list ValueList) GetCount() int {
	return list.NumValue()
}

func (list ValueList) GetValueIdx(_ rt.Runtime, i int) (ret meta.Value, err error) {
	if n := list.NumValue(); i >= n {
		err = errutil.New("out of range")
	} else {
		ret = list.ValueNum(i)
	}
	return
}

type NumberValueList ValueList

func (list NumberValueList) GetNumberIdx(r rt.Runtime, i int) (ret rt.Number, err error) {
	if v, e := ValueList(list).GetValueIdx(r, i); e != nil {
		err = e
	} else {
		ret = N(v.GetNum())
	}
	return
}

type TextValueList ValueList

func (list TextValueList) GetTextIdx(r rt.Runtime, i int) (ret rt.Text, err error) {
	if v, e := ValueList(list).GetValueIdx(r, i); e != nil {
		err = e
	} else {
		ret = T(v.GetText())
	}
	return
}

type RefValueList ValueList

func (list RefValueList) GetReferenceIdx(r rt.Runtime, i int) (ret rt.Reference, err error) {
	if v, e := ValueList(list).GetValueIdx(r, i); e != nil {
		err = e
	} else {
		ret = rt.Reference(v.GetObject())
	}
	return
}
