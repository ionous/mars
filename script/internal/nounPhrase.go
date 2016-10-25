package internal

import (
	"github.com/ionous/sashimi/util/errutil"
)

// NounPhrase builds "The" and "Our" statements.
// The statements are composed of "Fragment"s, all relating to a single subject.
// The presence of a Called fragements switches the subject.
// ( MARS its called oldstyle because i want to remove the called fragment )
type NounPhrase struct {
	Target    string
	Fragments Fragments
}

func (p NounPhrase) Build(src Source) (err error) {
	topic := Topic{p.Target, p.findSubject()}
	for _, frag := range p.Fragments {
		if e := frag.BuildFragment(src, topic); e != nil {
			err = errutil.Append(err, e)
		}
	}
	return err
}

func (p NounPhrase) findSubject() string {
	subject := p.Target // by default,
	for _, f := range p.Fragments {
		if called, ok := f.(ScriptSubject); ok {
			subject = called.Subject
			break
		}
	}
	return subject
}
