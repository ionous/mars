package internal

import (
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/lang"
	"strings"
)

func SetSubject(s string) Fragment {
	return ScriptSubject{Subject: s}
}

type ScriptSubject struct {
	Subject  string // name of the class or instance being declared
	Singular string // optional singular version of that name
}

func (c ScriptSubject) WithSingularName(name string) Fragment {
	c.Singular = name
	return c
}

func (c ScriptSubject) BuildFragment(src Source, top Topic) error {
	// FIX: this is only half measure --
	// really it needs to split into words, then compare the first article.
	name := strings.TrimSpace(top.Subject)
	article, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"article":       article,
		"long name":     name,
		"singular name": c.Singular,
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, UnknownLocation)
}
