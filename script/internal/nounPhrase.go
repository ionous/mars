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

func (sc NounPhrase) Build(src Source) (err error) {
	topic := Topic{sc.Target, sc.findSubject()}
	for _, frag := range sc.Fragments {
		if e := frag.BuildFragment(src, topic); e != nil {
			err = errutil.Append(err, e)
		}
	}
	return err
}

func (sc NounPhrase) findSubject() string {
	subject := sc.Target // by default,
	for _, f := range sc.Fragments {
		if called, ok := f.(ScriptSubject); ok {
			subject = called.Subject
			break
		}
	}
	return subject
}
