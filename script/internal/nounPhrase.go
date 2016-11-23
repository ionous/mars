package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

// NounPhrase builds "The" and "Our" statements.
// The statements are composed of "Fragment"s, all relating to a single subject.
// The presence of a Called fragements switches the subject.
// ( MARS its called oldstyle because i want to remove the called fragment )
type NounPhrase struct {
	Target    types.Subject `mars:"the [subject]"`
	Fragments []Fragment    `mars:"[fragment]"`
}

func (p NounPhrase) Generate(src *S.Statements) (err error) {
	if s, e := p.findSubject(); e != nil {
		err = e
	} else {
		topic := Topic{string(p.Target), s}
		for _, frag := range p.Fragments {
			if e := frag.GenFragment(src, topic); e != nil {
				err = errutil.Append(err, e)
			}
		}
	}
	return err
}

func (p NounPhrase) findSubject() (ret types.Subject, err error) {
	subject, found := p.Target, false // by default
	for _, f := range p.Fragments {
		if called, ok := f.(ScriptSubject); ok {
			if !found {
				subject = types.Subject(called.Subject)
				found = true
			} else {
				err = errutil.New("phrase has multiple subjects: was", sbuf.Q(subject), "now", sbuf.Q(called.Subject))
				break
			}
		}
	}
	if err == nil {
		ret = types.Subject(subject)
	}
	return
}
