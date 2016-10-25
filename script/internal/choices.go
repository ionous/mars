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

func (sel Choices) BuildFragment(src Source, top Topic) (err error) {
	for _, choice := range sel.choices {
		fields := S.ChoiceFields{top.Subject, choice}
		if e := src.NewChoice(fields, UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}

func (sel Choices) And(choice string) Choices {
	sel.choices = append(sel.choices, choice)
	return sel
}
