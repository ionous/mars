package frag

import (
	"github.com/ionous/mars/script"
	S "github.com/ionous/sashimi/source"
)

type Select struct {
	Choices []string
}

func (sel Select) Build(src script.Source, top Topic) (err error) {
	for _, choice := range sel.Choices {
		fields := S.ChoiceFields{top.Subject, choice}
		if e := src.NewChoice(fields, script.Unknown); e != nil {
			err = e
			break
		}
	}
	return err
}

func (sel Select) And(choice string) Select {
	sel.Choices = append(sel.Choices, choice)
	return sel
}
