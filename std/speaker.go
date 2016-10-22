package std

import (
	"github.com/ionous/mars/rt"
)

func Speaker(actor string) SpeechPhrase {
	return SpeechPhrase{actor: actor}
}

func (s SpeechPhrase) Says(lines string) rt.Execute {
	panic("not impleented")
	// return RunWithTex{
	// 	Id(s.actor),
	// 	"says",
	// 	rt.Text(lines)}
}

type SpeechPhrase struct {
	actor string
}
