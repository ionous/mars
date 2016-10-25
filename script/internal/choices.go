package internal

import (
	S "github.com/ionous/sashimi/source"
)

func SetChoices(choices ...string) Choices {
	return Choices{choices}
}

type Choices struct {
	choices []string
}

func (f Choices) BuildFragment(src Source, top Topic) (err error) {
	for _, choice := range f.choices {
		fields := S.ChoiceFields{top.Subject, choice}
		if e := src.NewChoice(fields, UnknownLocation); e != nil {
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
