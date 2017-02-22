package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
)

// NounDirective builds "The" and "Our" statements.
// The statements are composed of "Fragment"s, all relating to a single subject.
// The presence of a Called fragements switches the subject.
// ( MARS its called oldstyle because i want to remove the called fragment )
type NounDirective struct {
	Target    string     `mars:"The [subject]"`
	Fragments []Fragment `mars:"[phrases]"`
}

func (p NounDirective) Generate(src *S.Statements) (err error) {
	if s, e := p.findSubject(); e != nil {
		err = e
	} else {
		topic := Topic{p.Target, s}
		for _, frag := range p.Fragments {
			if e := frag.GenFragment(src, topic); e != nil {
				err = errutil.Append(err, e)
			}
		}
	}
	return err
}

func (p NounDirective) findSubject() (ret string, err error) {
	subject, found := p.Target, false // by default
	for _, f := range p.Fragments {
		if called, ok := f.(ScriptSubject); ok {
			if !found {
				subject = called.Subject
				found = true
			} else {
				err = errutil.New("phrase has multiple subjects: was", sbuf.Q(subject), "now", sbuf.Q(called.Subject))
				break
			}
		}
	}
	if err == nil {
		ret = subject
	}
	return
}
