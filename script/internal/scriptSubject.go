package internal

import (
	. "github.com/ionous/mars/script/backend"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/util/lang"
	"strings"
)

func SetSubject(s string) ScriptSubject {
	return ScriptSubject{Subject: s}
}

type ScriptSubject struct {
	Subject string `mars:"called [subject]"` // name of the class or instance being declared
}

func (c ScriptSubject) GenFragment(src *S.Statements, top Topic) error {
	// FIX: this is only half measure --
	// really it needs to split into words, then compare the first article.
	name := strings.TrimSpace(top.Subject.String())
	article, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"article":   article,
		"long name": name,
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, S.UnknownLocation)
}

type ScriptSingular struct {
	Singular string `mars:"has singular name [subject]"`
}

func SetSingularName(s string) ScriptSingular {
	return ScriptSingular{s}
}

func (c ScriptSingular) GenFragment(src *S.Statements, top Topic) error {
	name := top.Subject.String()
	_, bare := lang.SliceArticle(name)
	opt := map[string]string{
		"singular name": c.Singular,
	}
	fields := S.AssertionFields{top.Target, bare, opt}
	return src.NewAssertion(fields, S.UnknownLocation)
}
