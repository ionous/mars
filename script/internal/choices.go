package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type Choices struct {
	Is []string `mars: is [...]";choices"`
}

func (f Choices) GenFragment(src *S.Statements, top Topic) (err error) {
	for _, choice := range f.Is {
		fields := S.ChoiceFields{top.Subject, choice}
		if e := src.NewChoice(fields, S.UnknownLocation); e != nil {
			err = e
			break
		}
	}
	return err
}

func (f Choices) And(choice string) Choices {
	f.Is = append(f.Is, choice)
	return f
}
