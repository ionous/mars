package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

func HasChoices(choices ...string) Choices {
	return Choices{choices}
}

type Choices struct {
	choices []string
}

func (f Choices) GenFragment(src *S.Statements, top Topic) (err error) {
	for _, choice := range f.choices {
		fields := S.ChoiceFields{string(top.Subject), choice}
		if e := src.NewChoice(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}

func (f Choices) And(choice string) Choices {
	f.choices = append(f.choices, choice)
	return f
}
