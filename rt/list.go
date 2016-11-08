package rt

import (
	"github.com/ionous/sashimi/util/errutil"
)

// Numbers provides an array literal for floats.
type Numbers []float64

func (l Numbers) GetNumStream(Runtime) (NumberStream, error) {
	return &NumberIt{list: l}, nil
}

type NumberIt struct {
	list Numbers
	idx  int
}

func (it *NumberIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *NumberIt) GetNext() (ret Number, err error) {
	if !it.HasNext() {
		err = errutil.New("out of range")
	} else {
		ret = Number(it.list[it.idx])
		it.idx++
	}
	return
}

// Texts provides an array literal for strings.
type Texts []string

func (l Texts) GetTextStream(Runtime) (TextStream, error) {
	return &TextIt{list: l}, nil
}

type TextIt struct {
	list Texts
	idx  int
}

func (it *TextIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *TextIt) GetNext() (ret Text, err error) {
	if !it.HasNext() {
		err = errutil.New("out of range")
	} else {
		ret = Text(it.list[it.idx])
		it.idx++
	}
	return
}

// References provides an array literal for object ids.
type References []ObjEval

func (l References) GetObjStream(run Runtime) (ObjectStream, error) {
	return &RefIt{run: run, list: l}, nil
}

type RefIt struct {
	run  Runtime
	list References
	idx  int
}

func (it *RefIt) HasNext() bool {
	return it.idx < len(it.list)
}

func (it *RefIt) GetNext() (ret Object, err error) {
	if !it.HasNext() {
		err = errutil.New("out of range")
	} else {
		ref := it.list[it.idx]
		if obj, e := ref.GetObject(it.run); e != nil {
			err = e
		} else {
			ret = obj
			it.idx++
		}
	}
	return
}
