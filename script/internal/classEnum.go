package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
	"strings"
)

type ClassEnum struct {
	Choices     types.NamedChoices
	UsualChoice types.NamedChoice
	//expects []S.PropertyExpectation
}

type EitherChoice types.NamedChoice

func (p EitherChoice) Or(secondChoice types.NamedChoice) ClassEnum {
	either, or := types.NamedChoice(p).String(), secondChoice.String()
	return ClassEnum{Choices: types.NamedChoices{either, or}}
}

func (f ClassEnum) Usually(choice types.NamedChoice) ClassEnum {
	f.UsualChoice = choice
	return f
}

func (f ClassEnum) GenFragment(src *S.Statements, top Topic) (err error) {
	name := f.Choices[0] //-property
	enum := S.EnumFields{top.Subject, name, f.Choices.Strings()}
	if len(f.UsualChoice) > 0 {
		for i, v := range f.Choices {
			if strings.EqualFold(v, f.UsualChoice.String()) {
				f.Choices[0], f.Choices[i] = f.Choices[i], f.Choices[0]
				break
			}
		}
		if f.Choices[0] != f.UsualChoice.String() {
			err = errutil.New("usually not found", f.UsualChoice)
		}
	}
	if err == nil {
		err = src.NewEnumeration(enum, S.UnknownLocation)
	}
	return
}
