package internal

import (
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

type ClassEnum struct {
	Choices     []string
	UsualChoice string
	//expects []S.PropertyExpectation
}

type EitherChoice string

func (p EitherChoice) Or(secondChoice string) ClassEnum {
	return ClassEnum{Choices: []string{string(p), secondChoice}}
}

func (f ClassEnum) Usually(choice string) ClassEnum {
	f.UsualChoice = choice
	return f
}

func (f ClassEnum) GenFragment(src *S.Statements, top Topic) (err error) {
	name := f.Choices[0] //-property
	enum := S.EnumFields{top.Subject, name, f.Choices}
	if len(f.UsualChoice) > 0 {
		for i, v := range f.Choices {
			if strings.EqualFold(v, f.UsualChoice) {
				f.Choices[0], f.Choices[i] = f.Choices[i], f.Choices[0]
				break
			}
		}
		if f.Choices[0] != f.UsualChoice {
			err = errutil.New("usually not found %s", f.UsualChoice)
		}
	}
	if err == nil {
		err = src.NewEnumeration(enum, S.UnknownLocation)
	}
	return
}
