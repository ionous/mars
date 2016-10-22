package rt

import (
	"github.com/ionous/sashimi/util/errutil"
)

// Numbers provides an array literal for floats.
type Numbers []float64

func (l Numbers) GetCount() int {
	return len(l)
}

func (l Numbers) GetNumberIdx(run Runtime, i int) (ret Number, err error) {
	if cnt := len(l); i < cnt {
		ret = Number(l[i])
	} else {
		err = errutil.New("out of range")
	}
	return
}

// Texts provides an array literal for strings.
type Texts []string

func (l Texts) GetCount() int {
	return len(l)
}

// GetTextIdx implements TextListEval.
func (l Texts) GetTextIdx(run Runtime, i int) (ret Text, err error) {
	if cnt := len(l); i < cnt {
		ret = Text(l[i])
	} else {
		err = errutil.New("out of range")
	}
	return
}

// References provides an array literal for object ids.
type References []Reference

func (l References) GetCount() int {
	return len(l)
}

// GetReferenceIdx implements ObjListEval
func (l References) GetReferenceIdx(run Runtime, i int) (ret Reference, err error) {
	if cnt := len(l); i < cnt {
		ret = l[i]
	} else {
		err = errutil.New("out of range")
	}
	return
}
