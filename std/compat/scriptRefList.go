package compat

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/util/errutil"
)

type ScriptRefList struct {
	List rt.ObjListEval
}

func (x ScriptRefList) GetObjStream(run rt.Runtime) (rt.ObjectStream, error) {
	return x.List.GetObjStream(run)
}

func (l ScriptRefList) Empty() rt.BoolEval {
	return ObjListIsEmpty{l}
}
func (l ScriptRefList) Contains(which rt.ObjEval) rt.BoolEval {
	return ObjListContains{l, which}
}

type ObjListIsEmpty struct {
	In rt.ObjListEval
}

func (op ObjListIsEmpty) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if s, e := op.In.GetObjStream(run); e != nil {
		err = errutil.New("ObjListContains, couldnt get stream", e)
	} else {
		ok := !s.HasNext()
		ret = rt.Bool{ok}
	}
	return
}

type ObjListContains struct {
	In    rt.ObjListEval
	Which rt.ObjEval
}

func (op ObjListContains) GetBool(run rt.Runtime) (ret rt.Bool, err error) {
	if s, e := op.In.GetObjStream(run); e != nil {
		err = errutil.New("ObjListContains, couldnt get stream", e)
	} else if obj, e := op.Which.GetObject(run); e != nil {
		err = errutil.New("ObjListContains, couldnt get object", e)
	} else {
		for s.HasNext() {
			if it, e := s.GetNext(); e != nil {
				err = errutil.New("ObjListContains, couldnt get element", e)
				break
			} else if obj.Equals(it) {
				ret = rt.Bool{true}
			}
		}
	}
	return
}
