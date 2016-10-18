package frag

import (
	"github.com/ionous/mars/script"
	"github.com/ionous/sashimi/util/errutil"
)

type Topic struct {
	Target, Subject string
}

type Fragment interface {
	Build(script.Source, Topic) error
}

type Fragments []Fragment

type TheOldStyle struct {
	Target    string
	Fragments Fragments
}

func (sc TheOldStyle) Build(src script.Source) (err error) {
	topic := Topic{sc.Target, sc.findSubject()}
	for _, frag := range sc.Fragments {
		if e := frag.Build(src, topic); e != nil {
			err = errutil.Append(err, e)
		}
	}
	return err
}

func (sc TheOldStyle) findSubject() (ret string) {
	for _, f := range sc.Fragments {
		if called, ok := f.(SetTopic); ok {
			ret = called.Subject
			break
		}
	}
	return ret
}
