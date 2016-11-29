package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

func HasChoices(choices ...string) Choices {
	return Choices{choices}
}

type Choices struct {
	Choices types.NamedChoices `mars:"has"`
}

func (f Choices) GenFragment(src *S.Statements, top Topic) (err error) {
	for _, choice := range f.Choices {
		fields := S.ChoiceFields{top.Subject.String(), choice}
		if e := src.NewChoice(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}

func (f Choices) And(choice string) Choices {
	f.Choices = append(f.Choices, choice)
	return f
}
