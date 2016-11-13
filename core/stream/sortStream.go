package stream

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
	"github.com/ionous/sashimi/util/errutil"
	"sort"
)

// Sort creates a stream of objects sorted by key.
type KeySort struct {
	Key string
	Src rt.ObjListEval
}

// GetObjStream implements ObjListEval
func (g KeySort) GetObjStream(run rt.Runtime) (ret rt.ObjectStream, err error) {
	if src, e := g.Src.GetObjStream(run); e != nil {
		err = errutil.New("KeySort", e)
	} else {
		os := &objectSorter{run: run, key: g.Key}
		for src.HasNext() {
			if obj, e := src.GetNext(); e != nil {
				err = errutil.New("KeySort GetNext", e)
				break
			} else {
				os.objects = append(os.objects, obj)
			}
		}
		if err == nil {
			sort.Sort(os)
			ret = &Objects{list: os.objects}
		}
	}
	return
}

// objectSorter joins a By function and a slice of Planets to be sorted.
type objectSorter struct {
	run     rt.Runtime
	objects []rt.Object
	key     string
	err     error
}

// Len implements sort.Interface.
func (s *objectSorter) Len() int {
	return len(s.objects)
}

// Swap implements sort.Interface.
func (s *objectSorter) Swap(i, j int) {
	s.objects[i], s.objects[j] = s.objects[j], s.objects[i]
}

// Less implements sort.Interface.
func (s *objectSorter) Less(i, j int) (ret bool) {
	oa, ob := &s.objects[i], &s.objects[j]
	if less, e := s.sort(oa, ob); e != nil {
		ret, s.err = i < j, e
	} else {
		ret = less
	}
	return
}

func (os *objectSorter) sort(oa *rt.Object, ob *rt.Object) (ret bool, err error) {
	// FIX: do FindProperty up front somehow?
	if pa, ok := oa.FindProperty(os.key); !ok {
		err = errutil.New("no such property", os.key, "in", oa)
	} else if pb, ok := ob.GetProperty(pa.GetId()); !ok {
		err = errutil.New("no such property", os.key, "in", ob)
	} else {
		run := os.run

		switch pa.GetType() {
		case meta.NumProperty:
			// FIX: maybe i should implement GetNumber and GetText, etc. directly.
			if va, ok := pa.GetGeneric().(rt.NumberEval); !ok {
				err = errutil.New("couldnt covert pa to number", pa)
			} else if vb, ok := pb.GetGeneric().(rt.NumberEval); !ok {
				err = errutil.New("couldnt covert pb to number", pb)
			} else if a, e := va.GetNumber(run); e != nil {
				err = errutil.New("couldnt get pa number", e)
			} else if b, e := vb.GetNumber(run); e != nil {
				err = errutil.New("couldnt get pb number", e)
			} else {
				ret = a < b
			}
		case meta.TextProperty:
			if va, ok := pa.GetGeneric().(rt.TextEval); !ok {
				err = errutil.New("couldnt covert pa to text", pa)
			} else if vb, ok := pb.GetGeneric().(rt.TextEval); !ok {
				err = errutil.New("couldnt covert pb to text", pb)
			} else if a, e := va.GetText(run); e != nil {
				err = errutil.New("couldnt get pa text", e)
			} else if b, e := vb.GetText(run); e != nil {
				err = errutil.New("couldnt get pb text", e)
			} else {
				ret = a < b
			}
		default:
			err = errutil.New("invalid property type", pa.GetName(), pa.GetType())
		}
	}
	return
}
