package scope

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/sashimi/meta"
)

type NumScope struct {
	val rt.NumEval
}

func NewNumScope(val rt.NumEval) rt.FindValue {
	return &NumScope{val}
}

func (s *NumScope) FindValue(name string) (ret meta.Generic, err error) {
	if name != "" {
		err = NotFound(s, "num is not an object")
	} else {
		ret = s.val
	}
	return
}

func (s *NumScope) ScopePath() []string {
	return []string{"num"}
}

type TextScope struct {
	val rt.TextEval
}

func NewTextScope(val rt.TextEval) rt.FindValue {
	return &TextScope{val}
}

func (s *TextScope) FindValue(name string) (ret meta.Generic, err error) {
	if name != "" {
		err = NotFound(s, "text is not an object")
	} else {
		ret = s.val
	}
	return
}

func (s *TextScope) ScopePath() []string {
	return []string{"text"}
}
