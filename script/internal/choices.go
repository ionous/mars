package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
)

type Choice struct {
	Is string `mars:"is [choice];choice"`
}

func (f Choice) GenFragment(src *S.Statements, top Topic) (err error) {
	fields := S.ChoiceFields{top.Subject, f.Is}
	return src.NewChoice(fields, S.UnknownLocation)
}
