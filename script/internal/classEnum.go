package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

type ClassEnum struct {
	Choices     []string `mars:";choice"`
	UsualChoice string   `mars:";choice"`
	//expects []S.PropertyExpectation
}

type EitherChoice string

func (p EitherChoice) Or(secondChoice string) ClassEnum {
	either, or := string(p), secondChoice
	return ClassEnum{Choices: []string{either, or}}
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
			err = errutil.New("usually not found", f.UsualChoice)
		}
	}
	if err == nil {
		err = src.NewEnumeration(enum, S.UnknownLocation)
	}
	return
}
